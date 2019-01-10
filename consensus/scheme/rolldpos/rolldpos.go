// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package rolldpos

import (
	"context"
	"time"

	"github.com/facebookgo/clock"
	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/zjshen14/go-fsm"
	"go.uber.org/zap"

	"github.com/iotexproject/iotex-core/action"
	"github.com/iotexproject/iotex-core/actpool"
	"github.com/iotexproject/iotex-core/blockchain"
	"github.com/iotexproject/iotex-core/blockchain/block"
	"github.com/iotexproject/iotex-core/blocksync"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/consensus/scheme"
	"github.com/iotexproject/iotex-core/crypto"
	"github.com/iotexproject/iotex-core/endorsement"
	"github.com/iotexproject/iotex-core/explorer/idl/explorer"
	"github.com/iotexproject/iotex-core/iotxaddress"
	"github.com/iotexproject/iotex-core/pkg/hash"
	"github.com/iotexproject/iotex-core/pkg/log"
	"github.com/iotexproject/iotex-core/proto"
	"github.com/iotexproject/iotex-core/state"
)

var (
	timeSlotMtc = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "iotex_consensus_time_slot",
			Help: "Consensus time slot",
		},
		[]string{},
	)
)

const sigSize = 5 // number of uint32s in BLS sig

func init() {
	prometheus.MustRegister(timeSlotMtc)
}

var (
	// ErrNewRollDPoS indicates the error of constructing RollDPoS
	ErrNewRollDPoS = errors.New("error when constructing RollDPoS")
	// ErrZeroDelegate indicates seeing 0 delegates in the network
	ErrZeroDelegate = errors.New("zero delegates in the network")
)

type rollDPoSCtx struct {
	cfg              config.RollDPoS
	addr             *iotxaddress.Address
	chain            blockchain.Blockchain
	actPool          actpool.ActPool
	broadcastHandler scheme.Broadcast
	epoch            epochCtx
	round            roundCtx
	clock            clock.Clock
	rootChainAPI     explorer.Explorer
	// candidatesByHeightFunc is only used for testing purpose
	candidatesByHeightFunc func(uint64) ([]*state.Candidate, error)
	sync                   blocksync.BlockSync
}

var (
	// ErrNotEnoughCandidates indicates there are not enough candidates from the candidate pool
	ErrNotEnoughCandidates = errors.New("Candidate pool does not have enough candidates")
)

func (ctx *rollDPoSCtx) OnConsensusReached() error {
	pendingBlock := ctx.round.block.(*blockWrapper)
	// If the pending block is a secret block, record the secret share generated by producer
	if ctx.shouldHandleDKG() {
		for _, secretProposal := range pendingBlock.SecretProposals {
			if secretProposal.DstAddr() == ctx.addr.RawAddress {
				ctx.epoch.committedSecrets[secretProposal.SrcAddr()] = secretProposal.Secret()
				break
			}
		}
	}
	// Commit and broadcast the pending block
	if err := ctx.chain.CommitBlock(pendingBlock.Block); err != nil {
		log.L().Error("Error when committing a block.",
			zap.Error(err),
			zap.Uint64("blockHeight", pendingBlock.Height()))
	}
	// Remove transfers in this block from ActPool and reset ActPool state
	ctx.actPool.Reset()
	// Broadcast the committed block to the network
	if blkProto := pendingBlock.ConvertToBlockPb(); blkProto != nil {
		if err := ctx.Broadcast(blkProto); err != nil {
			log.L().Error("Error when broadcasting blkProto.",
				zap.Error(err),
				zap.Uint64("blockHeight", pendingBlock.Height()))
		}
		// putblock to parent chain if the current node is proposer and current chain is a sub chain
		if ctx.round.proposer == ctx.addr.RawAddress && ctx.chain.ChainAddress() != "" {
			putBlockToParentChain(ctx.rootChainAPI, ctx.chain.ChainAddress(), ctx.addr, pendingBlock.Block)
		}
	} else {
		log.L().Error("Error whenconverting a block into a proto msg.",
			zap.Uint64("blockHeight", pendingBlock.Height()))
	}
	return nil
}

type blockWrapper struct {
	*block.Block

	round uint32
}

func (bw *blockWrapper) Hash() []byte {
	hash := bw.HashBlock()

	return hash[:]
}

func (bw *blockWrapper) Proposer() string {
	return bw.ProducerAddress()
}

func (bw *blockWrapper) Round() uint32 {
	return bw.round
}

func (ctx *rollDPoSCtx) MintBlock() (Block, error) {
	if blk := ctx.round.block; blk != nil {
		return blk, nil
	}
	blk, err := ctx.mintBlock()
	if err != nil {
		return nil, err
	}

	return &blockWrapper{
		blk,
		ctx.round.number,
	}, nil
}

func (ctx *rollDPoSCtx) validateProposeBlock(blk Block, expectedProposer string) bool {
	blkHash := blk.Hash()
	errorLog := log.L().With(zap.Uint64("expectedHeight", ctx.round.height),
		zap.String("expectedProposer", expectedProposer),
		log.Hex("hash", blkHash))
	if blk.Height() != ctx.round.height {
		errorLog.Error("Error when validating the block height.",
			zap.Uint64("blockHeight", blk.Height()))
		return false
	}
	producer := blk.Proposer()

	if producer == "" || producer != expectedProposer {
		errorLog.Error("Error when validating the block proposer.", zap.String("proposer", producer))
		return false
	}
	block := blk.(*blockWrapper)
	if !block.VerifySignature() {
		errorLog.Error("Error when validating the block signature.")
		return false
	}
	// TODO: in long term, block in process and after process should be represented differently
	if producer == ctx.addr.RawAddress && block.WorkingSet != nil {
		// If the block is self proposed and working set is not nil (meaning not obtained from wire), skip validation
		return true
	}
	containCoinbase := true
	if ctx.cfg.EnableDKG {
		if ctx.shouldHandleDKG() {
			containCoinbase = false
		} else if err := verifyDKGSignature(block.Block, ctx.epoch.seed); err != nil {
			// Verify dkg signature failed
			errorLog.Error("Failed to verify the DKG signature.", zap.Error(err))
			return false
		}

	}
	if err := ctx.chain.ValidateBlock(block.Block, containCoinbase); err != nil {
		errorLog.Error("error when validating the proposed block", zap.Error(err))
		return false
	}

	return true
}

func verifyDKGSignature(blk *block.Block, seedByte []byte) error {
	return crypto.BLS.Verify(blk.DKGPubkey(), seedByte, blk.DKGSignature())
}

func (ctx *rollDPoSCtx) Broadcast(msg proto.Message) error {
	var t iproto.ConsensusPb_ConsensusMessageType
	switch msg.(type) {
	case *iproto.BlockPb:
		return ctx.broadcastHandler(msg)
	case *iproto.ProposePb:
		t = iproto.ConsensusPb_PROPOSAL
	case *iproto.EndorsePb:
		t = iproto.ConsensusPb_ENDORSEMENT
	default:
		return errors.New("Invalid message type")
	}
	data, err := proto.Marshal(msg)
	if err != nil {
		return err
	}
	consensusMsg := &iproto.ConsensusPb{
		Height:    ctx.round.height,
		Round:     ctx.round.number,
		Type:      t,
		Data:      data,
		Timestamp: uint64(ctx.clock.Now().Unix()),
	}

	return ctx.broadcastHandler(consensusMsg)
}

// rollingDelegates will only allows the delegates chosen for given epoch to enter the epoch
func (ctx *rollDPoSCtx) rollingDelegates(epochNum uint64) ([]string, error) {
	numDlgs := ctx.cfg.NumDelegates
	height := uint64(numDlgs) * uint64(ctx.cfg.NumSubEpochs) * (epochNum - 1)
	var candidates []*state.Candidate
	var err error
	if ctx.candidatesByHeightFunc != nil {
		// Test only
		candidates, err = ctx.candidatesByHeightFunc(height)
	} else {
		candidates, err = ctx.chain.CandidatesByHeight(height)
	}
	if err != nil {
		return []string{}, errors.Wrap(err, "error when getting delegates from the candidate pool")
	}
	if len(candidates) < int(numDlgs) {
		return []string{}, errors.Wrapf(ErrNotEnoughCandidates, "only %d delegates from the candidate pool", len(candidates))
	}

	var candidatesAddress []string
	for _, candidate := range candidates {
		candidatesAddress = append(candidatesAddress, candidate.Address)
	}
	crypto.SortCandidates(candidatesAddress, epochNum, ctx.epoch.seed)

	return candidatesAddress[:numDlgs], nil
}

// calcEpochNum calculates the epoch ordinal number and the epoch start height offset, which is based on the height of
// the next block to be produced
func (ctx *rollDPoSCtx) calcEpochNumAndHeight() (uint64, uint64, error) {
	height := ctx.chain.TipHeight()
	numDlgs := ctx.cfg.NumDelegates
	numSubEpochs := ctx.getNumSubEpochs()
	epochNum := height/(uint64(numDlgs)*uint64(numSubEpochs)) + 1
	epochHeight := uint64(numDlgs)*uint64(numSubEpochs)*(epochNum-1) + 1
	return epochNum, epochHeight, nil
}

// calcSubEpochNum calculates the sub-epoch ordinal number
func (ctx *rollDPoSCtx) calcSubEpochNum() (uint64, error) {
	height := ctx.chain.TipHeight() + 1
	if height < ctx.epoch.height {
		return 0, errors.New("Tip height cannot be less than epoch height")
	}
	numDlgs := ctx.cfg.NumDelegates
	subEpochNum := (height - ctx.epoch.height) / uint64(numDlgs)
	return subEpochNum, nil
}

// shouldHandleDKG indicates whether a node is in DKG stage
func (ctx *rollDPoSCtx) shouldHandleDKG() bool {
	if !ctx.cfg.EnableDKG {
		return false
	}
	return ctx.epoch.subEpochNum == 0
}

// generateDKGSecrets generates DKG secrets and witness
func (ctx *rollDPoSCtx) generateDKGSecrets() ([][]uint32, [][]byte, error) {
	idList := make([][]uint8, 0)
	for _, addr := range ctx.epoch.delegates {
		dkgID := iotxaddress.CreateID(addr)
		idList = append(idList, dkgID)
		if addr == ctx.addr.RawAddress {
			ctx.epoch.dkgAddress = iotxaddress.DKGAddress{ID: dkgID}
		}
	}
	_, secrets, witness, err := crypto.DKG.Init(crypto.DKG.SkGeneration(), idList)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate DKG Secrets and Witness")
	}
	return secrets, witness, nil
}

// TODO: numDlgs should also be configurable in BLS. For test purpose, let's make it 21.
// generateDKGKeyPair generates DKG key pair
func (ctx *rollDPoSCtx) generateDKGKeyPair() ([]byte, []uint32, error) {
	numDlgs := ctx.cfg.NumDelegates
	if numDlgs != 21 {
		return nil, nil, errors.New("Number of delegates must be 21 for test purpose")
	}
	shares := make([][]uint32, numDlgs)
	shareStatusMatrix := make([][21]bool, numDlgs)
	for i := range shares {
		shares[i] = make([]uint32, sigSize)
	}
	for i, delegate := range ctx.epoch.delegates {
		if secret, ok := ctx.epoch.committedSecrets[delegate]; ok {
			shares[i] = secret
			for j := 0; j < int(numDlgs); j++ {
				shareStatusMatrix[j][i] = true
			}
		}
	}
	_, dkgPubKey, dkgPriKey, err := crypto.DKG.KeyPairGeneration(shares, shareStatusMatrix)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate DKG key pair")
	}
	return dkgPubKey, dkgPriKey, nil
}

// getNumSubEpochs returns max(configured number, 1)
func (ctx *rollDPoSCtx) getNumSubEpochs() uint {
	num := uint(1)
	if ctx.cfg.NumSubEpochs > 0 {
		num = ctx.cfg.NumSubEpochs
	}
	if ctx.cfg.EnableDKG {
		num++
	}
	return num
}

// rotatedProposer will rotate among the delegates to choose the proposer. It is pseudo order based on the position
// in the delegate list and the block height
func (ctx *rollDPoSCtx) rotatedProposer() (string, uint64, uint32, error) {
	// Next block height
	height := ctx.chain.TipHeight() + 1
	round, proposer, err := ctx.calcProposer(height, ctx.epoch.delegates)

	return proposer, height, round, err
}

// calcProposer calculates the proposer for the block at a given height
func (ctx *rollDPoSCtx) calcProposer(height uint64, delegates []string) (uint32, string, error) {
	numDelegates := len(delegates)
	if numDelegates == 0 {
		return 0, "", ErrZeroDelegate
	}
	timeSlotIndex := uint32(0)
	if ctx.cfg.ProposerInterval > 0 {
		duration, err := ctx.calcDurationSinceLastBlock()
		if err != nil || duration < 0 {
			if !ctx.cfg.TimeBasedRotation {
				return 0, delegates[(height)%uint64(numDelegates)], nil
			}
			return 0, "", errors.Wrap(err, "error when computing the duration since last block time")
		}
		if duration > ctx.cfg.ProposerInterval {
			timeSlotIndex = uint32(duration/ctx.cfg.ProposerInterval) - 1
		}
	}
	if !ctx.cfg.TimeBasedRotation {
		return timeSlotIndex, delegates[(height)%uint64(numDelegates)], nil
	}
	// TODO: should downgrade to debug level in the future
	log.L().Info("Calculate time slot offset.", zap.Uint32("slot", timeSlotIndex))
	timeSlotMtc.WithLabelValues().Set(float64(timeSlotIndex))
	return timeSlotIndex, delegates[(height+uint64(timeSlotIndex))%uint64(numDelegates)], nil
}

// mintBlock mints a new block to propose
func (ctx *rollDPoSCtx) mintBlock() (*block.Block, error) {
	if ctx.shouldHandleDKG() {
		return ctx.mintSecretBlock()
	}
	return ctx.mintCommonBlock()
}

// mintSecretBlock collects DKG secret proposals and witness and creates a block to propose
func (ctx *rollDPoSCtx) mintSecretBlock() (*block.Block, error) {
	secrets := ctx.epoch.secrets
	witness := ctx.epoch.witness
	if len(secrets) != len(ctx.epoch.delegates) {
		return nil, errors.New("Number of secrets does not match number of delegates")
	}
	confirmedNonce, err := ctx.chain.Nonce(ctx.addr.RawAddress)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get the confirmed nonce of secret block producer")
	}
	nonce := confirmedNonce + 1
	secretProposals := make([]*action.SecretProposal, 0)
	for i, delegate := range ctx.epoch.delegates {
		secretProposal, err := action.NewSecretProposal(nonce, ctx.addr.RawAddress, delegate, secrets[i])
		if err != nil {
			return nil, errors.Wrap(err, "failed to create the secret proposal")
		}
		secretProposals = append(secretProposals, secretProposal)
		nonce++
	}
	secretWitness, err := action.NewSecretWitness(nonce, ctx.addr.RawAddress, witness)
	if err != nil {
		return nil, errors.Wrap(err, "failed to create the secret witness")
	}
	blk, err := ctx.chain.MintNewSecretBlock(secretProposals, secretWitness, ctx.addr)
	if err != nil {
		return nil, err
	}
	log.L().Info("minted a new secret block",
		zap.Uint64("height", blk.Height()),
		zap.Int("secretProposals", len(blk.SecretProposals)))
	return blk, nil
}

// mintCommonBlock picks the actions and creates a common block to propose
func (ctx *rollDPoSCtx) mintCommonBlock() (*block.Block, error) {
	actions := ctx.actPool.PickActs()
	log.L().Debug("Pick actions from the action pool.", zap.Int("action", len(actions)))
	blk, err := ctx.chain.MintNewBlock(actions, ctx.addr, &ctx.epoch.dkgAddress,
		ctx.epoch.seed, "")
	if err != nil {
		return nil, err
	}
	log.L().Info("Minted a new block.",
		zap.Uint64("height", blk.Height()),
		zap.Int("actions", len(blk.Actions)))
	return blk, nil
}

// calcDurationSinceLastBlock returns the duration since last block time
func (ctx *rollDPoSCtx) calcDurationSinceLastBlock() (time.Duration, error) {
	height := ctx.chain.TipHeight()
	blk, err := ctx.chain.GetBlockByHeight(height)
	if err != nil {
		return 0, errors.Wrapf(err, "error when getting the block at height: %d", height)
	}
	return ctx.clock.Now().Sub(time.Unix(blk.Header.Timestamp(), 0)), nil
}

// calcQuorum calculates if more than 2/3 vote yes or no including self's vote
func (ctx *rollDPoSCtx) calcQuorum(decisions map[string]bool) (bool, bool) {
	yes := 0
	no := 0
	for _, decision := range decisions {
		if decision {
			yes++
		} else {
			no++
		}
	}
	numDelegates := len(ctx.epoch.delegates)
	return yes >= numDelegates*2/3+1, no >= numDelegates*1/3
}

// isEpochFinished checks the epoch is finished or not
func (ctx *rollDPoSCtx) isEpochFinished() (bool, error) {
	height := ctx.chain.TipHeight()
	// if the height of the last committed block is already the last one should be minted from this epochStart, go back
	// to epochStart start
	if height >= ctx.epoch.height+uint64(uint(len(ctx.epoch.delegates))*ctx.epoch.numSubEpochs)-1 {
		return true, nil
	}
	return false, nil
}

// isDKGFinished checks the DKG sub-epoch is finished or not
func (ctx *rollDPoSCtx) isDKGFinished() bool {
	height := ctx.chain.TipHeight()
	return height >= ctx.epoch.height+uint64(len(ctx.epoch.delegates))-1
}

// updateSeed returns the seed for the next epoch
func (ctx *rollDPoSCtx) updateSeed() ([]byte, error) {
	epochNum, epochHeight, err := ctx.calcEpochNumAndHeight()
	if err != nil {
		return hash.Hash256b(ctx.epoch.seed), errors.Wrap(err, "Failed to do decode seed")
	}
	if epochNum <= 1 {
		return crypto.CryptoSeed, nil
	}
	selectedID := make([][]uint8, 0)
	selectedSig := make([][]byte, 0)
	selectedPK := make([][]byte, 0)
	for h := uint64(ctx.cfg.NumDelegates)*uint64(ctx.cfg.NumSubEpochs)*(epochNum-2) + 1; h < epochHeight && len(selectedID) <= crypto.Degree; h++ {
		blk, err := ctx.chain.GetBlockByHeight(h)
		if err != nil {
			continue
		}
		if len(blk.DKGID()) > 0 && len(blk.DKGPubkey()) > 0 && len(blk.DKGSignature()) > 0 {
			selectedID = append(selectedID, blk.DKGID())
			selectedSig = append(selectedSig, blk.DKGSignature())
			selectedPK = append(selectedPK, blk.DKGPubkey())
		}
	}

	if len(selectedID) <= crypto.Degree {
		return hash.Hash256b(ctx.epoch.seed), errors.New("DKG signature/pubic key is not enough to aggregate")
	}

	aggregateSig, err := crypto.BLS.SignAggregate(selectedID, selectedSig)
	if err != nil {
		return hash.Hash256b(ctx.epoch.seed), errors.Wrap(err, "Failed to generate aggregate signature to update Seed")
	}
	if err = crypto.BLS.VerifyAggregate(selectedID, selectedPK, ctx.epoch.seed, aggregateSig); err != nil {
		return hash.Hash256b(ctx.epoch.seed), errors.Wrap(err, "Failed to verify aggregate signature to update Seed")
	}
	return aggregateSig, nil
}

// epochCtx keeps the context data for the current epoch
type epochCtx struct {
	// num is the ordinal number of an epoch
	num uint64
	// height means offset for current epochStart (i.e., the height of the first block generated in this epochStart)
	height uint64
	// numSubEpochs defines number of sub-epochs/rotations will happen in an epochStart
	numSubEpochs uint
	// subEpochNum is the ordinal number of sub-epoch within the current epoch
	subEpochNum uint64
	// secrets are the dkg secrets sent from current node to other delegates
	secrets [][]uint32
	// witness is the dkg secret witness sent from current node to other delegates
	witness [][]byte
	// committedSecrets are the secret shares within the secret blocks committed by current node
	committedSecrets map[string][]uint32
	delegates        []string
	dkgAddress       iotxaddress.DKGAddress
	seed             []byte
}

// roundCtx keeps the context data for the current round and block.
type roundCtx struct {
	height          uint64
	number          uint32
	proofOfLock     *endorsement.Set
	timestamp       time.Time
	block           Block
	endorsementSets map[string]*endorsement.Set
	proposer        string
}

// RollDPoS is Roll-DPoS consensus main entrance
type RollDPoS struct {
	cfsm *cFSM
	ctx  *rollDPoSCtx
}

// Start starts RollDPoS consensus
func (r *RollDPoS) Start(ctx context.Context) error {
	if err := r.cfsm.Start(ctx); err != nil {
		return errors.Wrap(err, "error when starting the consensus FSM")
	}
	r.cfsm.produce(r.cfsm.newCEvt(eRollDelegates), r.ctx.cfg.Delay)
	return nil
}

// Stop stops RollDPoS consensus
func (r *RollDPoS) Stop(ctx context.Context) error {
	return errors.Wrap(r.cfsm.Stop(ctx), "error when stopping the consensus FSM")
}

// HandleConsensusMsg handles incoming consensus message
func (r *RollDPoS) HandleConsensusMsg(msg *iproto.ConsensusPb) error {
	data := msg.Data
	tipHeight := r.ctx.chain.TipHeight()
	switch msg.Type {
	case iproto.ConsensusPb_PROPOSAL:
		pPb := &iproto.ProposePb{}
		if err := proto.Unmarshal(data, pPb); err != nil {
			return err
		}
		pbEvt, err := r.cfsm.newProposeBlkEvtFromProposePb(pPb)
		if err != nil {
			return errors.Wrap(err, "error when casting a proto msg to proposeBlkEvt")
		}
		if pbEvt.height() <= tipHeight {
			return errors.New("old block proposal")
		}
		r.cfsm.produce(pbEvt, 0)
	case iproto.ConsensusPb_ENDORSEMENT:
		ePb := &iproto.EndorsePb{}
		if err := proto.Unmarshal(data, ePb); err != nil {
			return err
		}
		eEvt, err := r.cfsm.newEndorseEvtWithEndorsePb(ePb)
		if err != nil {
			return errors.Wrap(err, "error when casting a proto msg to endorse")
		}
		if eEvt.height() <= tipHeight && eEvt.endorse.Endorser() != r.ctx.addr.RawAddress {
			log.L().Debug("ignore old endorsement message",
				zap.Uint64("eventHeight", eEvt.height()),
				zap.Uint64("chainHeight", tipHeight))
			return nil
		}
		r.cfsm.produce(eEvt, 0)
	default:
		return errors.Errorf("Invalid consensus message type %s", msg.Type)
	}
	return nil
}

// Metrics returns RollDPoS consensus metrics
func (r *RollDPoS) Metrics() (scheme.ConsensusMetrics, error) {
	var metrics scheme.ConsensusMetrics
	// Compute the epoch ordinal number
	epochNum, _, err := r.ctx.calcEpochNumAndHeight()
	if err != nil {
		return metrics, errors.Wrap(err, "error when calculating the epoch ordinal number")
	}
	// Compute delegates
	delegates, err := r.ctx.rollingDelegates(epochNum)
	if err != nil {
		return metrics, errors.Wrap(err, "error when getting the rolling delegates")
	}
	// Compute the height
	height := r.ctx.chain.TipHeight()
	// Compute block producer
	_, producer, err := r.ctx.calcProposer(height+1, delegates)
	if err != nil {
		return metrics, errors.Wrap(err, "error when calculating the block producer")
	}
	// Get all candidates
	candidates, err := r.ctx.chain.CandidatesByHeight(height)
	if err != nil {
		return metrics, errors.Wrap(err, "error when getting all candidates")
	}
	candidateAddresses := make([]string, len(candidates))
	for i, c := range candidates {
		candidateAddresses[i] = c.Address
	}

	crypto.SortCandidates(candidateAddresses, epochNum, r.ctx.epoch.seed)

	return scheme.ConsensusMetrics{
		LatestEpoch:         epochNum,
		LatestHeight:        height,
		LatestDelegates:     delegates,
		LatestBlockProducer: producer,
		Candidates:          candidateAddresses,
	}, nil
}

// NumPendingEvts returns the number of pending events
func (r *RollDPoS) NumPendingEvts() int {
	return len(r.cfsm.evtq)
}

// CurrentState returns the current state
func (r *RollDPoS) CurrentState() fsm.State {
	return r.cfsm.fsm.CurrentState()
}

// Builder is the builder for RollDPoS
type Builder struct {
	cfg config.RollDPoS
	// TODO: we should use keystore in the future
	addr                   *iotxaddress.Address
	chain                  blockchain.Blockchain
	actPool                actpool.ActPool
	broadcastHandler       scheme.Broadcast
	clock                  clock.Clock
	rootChainAPI           explorer.Explorer
	candidatesByHeightFunc func(uint64) ([]*state.Candidate, error)
}

// NewRollDPoSBuilder instantiates a Builder instance
func NewRollDPoSBuilder() *Builder {
	return &Builder{}
}

// SetConfig sets RollDPoS config
func (b *Builder) SetConfig(cfg config.RollDPoS) *Builder {
	b.cfg = cfg
	return b
}

// SetAddr sets the address and key pair for signature
func (b *Builder) SetAddr(addr *iotxaddress.Address) *Builder {
	b.addr = addr
	return b
}

// SetBlockchain sets the blockchain APIs
func (b *Builder) SetBlockchain(chain blockchain.Blockchain) *Builder {
	b.chain = chain
	return b
}

// SetActPool sets the action pool APIs
func (b *Builder) SetActPool(actPool actpool.ActPool) *Builder {
	b.actPool = actPool
	return b
}

// SetBroadcast sets the broadcast callback
func (b *Builder) SetBroadcast(broadcastHandler scheme.Broadcast) *Builder {
	b.broadcastHandler = broadcastHandler
	return b
}

// SetClock sets the clock
func (b *Builder) SetClock(clock clock.Clock) *Builder {
	b.clock = clock
	return b
}

// SetRootChainAPI sets root chain API
func (b *Builder) SetRootChainAPI(api explorer.Explorer) *Builder {
	b.rootChainAPI = api
	return b
}

// SetCandidatesByHeightFunc sets candidatesByHeightFunc, which is only used by tests
func (b *Builder) SetCandidatesByHeightFunc(
	candidatesByHeightFunc func(uint64) ([]*state.Candidate, error),
) *Builder {
	b.candidatesByHeightFunc = candidatesByHeightFunc
	return b
}

// Build builds a RollDPoS consensus module
func (b *Builder) Build() (*RollDPoS, error) {
	if b.chain == nil {
		return nil, errors.Wrap(ErrNewRollDPoS, "blockchain APIs is nil")
	}
	if b.actPool == nil {
		return nil, errors.Wrap(ErrNewRollDPoS, "action pool APIs is nil")
	}
	if b.broadcastHandler == nil {
		return nil, errors.Wrap(ErrNewRollDPoS, "broadcast callback is nil")
	}
	if b.clock == nil {
		b.clock = clock.New()
	}
	ctx := rollDPoSCtx{
		cfg:                    b.cfg,
		addr:                   b.addr,
		chain:                  b.chain,
		actPool:                b.actPool,
		broadcastHandler:       b.broadcastHandler,
		clock:                  b.clock,
		rootChainAPI:           b.rootChainAPI,
		candidatesByHeightFunc: b.candidatesByHeightFunc,
	}
	cfsm, err := newConsensusFSM(&ctx)
	if err != nil {
		return nil, errors.Wrap(err, "error when constructing the consensus FSM")
	}
	return &RollDPoS{
		cfsm: cfsm,
		ctx:  &ctx,
	}, nil
}
