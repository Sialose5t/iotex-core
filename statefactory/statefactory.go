// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package statefactory

import (
	"bytes"
	"container/heap"
	"encoding/gob"
	"fmt"
	"math/big"

	"github.com/pkg/errors"

	trx "github.com/iotexproject/iotex-core/blockchain/trx"
	"github.com/iotexproject/iotex-core/common"
	"github.com/iotexproject/iotex-core/iotxaddress"
	"github.com/iotexproject/iotex-core/trie"
)

//const (
//	delegateSize  = 101
//	candidateSize = 400
//	bufferSize    = 10000
//)

const (
	// Level 1 is for delegate pool
	delegatePool = 1
	// Level 2 is for candidate pool
	candidatePool = delegatePool + 1
	// Level 3 is for buffer
	buffer = candidatePool + 1
)

const (
	delegateSize  = 2
	candidateSize = 3
	bufferSize    = 4
)

var (
	stateFactoryKVNameSpace = "StateFactory"

	// ErrInvalidAddr is the error that the address format is invalid, cannot be decoded
	ErrInvalidAddr = errors.New("address format is invalid")

	// ErrNotEnoughBalance is the error that the balance is not enough
	ErrNotEnoughBalance = errors.New("not enough balance")

	// ErrAccountNotExist is the error that the account does not exist
	ErrAccountNotExist = errors.New("the account does not exist")

	// ErrFailedToMarshalState is the error that the state marshaling is failed
	ErrFailedToMarshalState = errors.New("failed to marshal state")

	// ErrFailedToUnmarshalState is the error that the state un-marshaling is failed
	ErrFailedToUnmarshalState = errors.New("failed to unmarshal state")
)

type (
	// State is the canonical representation of an account.
	State struct {
		Nonce        uint64
		Balance      *big.Int
		Address      string
		IsCandidate  bool
		VotingWeight *big.Int
		Votee        common.PKHash
		Voters       map[common.PKHash]*big.Int
	}

	// StateFactory defines an interface for managing states
	StateFactory interface {
		CreateState(string, uint64) (*State, error)
		Balance(string) (*big.Int, error)
		CommitStateChanges([]*trx.Transfer, []*trx.Vote) error
		SetNonce(string, uint64) error
		Nonce(string) (uint64, error)
		RootHash() common.Hash32B
		Delegates() []*Candidate
		Candidates() []*Candidate
	}

	// stateFactory implements StateFactory interface, tracks changes in a map and batch-commits to trie/db
	stateFactory struct {
		trie             trie.Trie
		delegateHeap     CandidateMinPQ
		candidateMinHeap CandidateMinPQ
		candidateMaxHeap CandidateMaxPQ
		minBuffer        CandidateMinPQ
		maxBuffer        CandidateMaxPQ
	}
)

func stateToBytes(s *State) ([]byte, error) {
	var ss bytes.Buffer
	e := gob.NewEncoder(&ss)
	if err := e.Encode(s); err != nil {
		return nil, ErrFailedToMarshalState
	}
	return ss.Bytes(), nil
}

func bytesToState(ss []byte) (*State, error) {
	var state State
	e := gob.NewDecoder(bytes.NewBuffer(ss))
	if err := e.Decode(&state); err != nil {
		return nil, ErrFailedToUnmarshalState
	}
	return &state, nil
}

//======================================
// functions for State
//======================================

// AddBalance adds balance for state
func (st *State) AddBalance(amount *big.Int) error {
	st.Balance.Add(st.Balance, amount)
	return nil
}

// SubBalance subtracts balance for state
func (st *State) SubBalance(amount *big.Int) error {
	// make sure there's enough fund to spend
	if amount.Cmp(st.Balance) == 1 {
		return ErrNotEnoughBalance
	}
	st.Balance.Sub(st.Balance, amount)
	return nil
}

//======================================
// functions for StateFactory
//======================================

// NewStateFactory creates a new state factory
func NewStateFactory(tr trie.Trie) StateFactory {
	return &stateFactory{
		trie:             tr,
		delegateHeap:     make(CandidateMinPQ, 0),
		candidateMinHeap: make(CandidateMinPQ, 0),
		candidateMaxHeap: make(CandidateMaxPQ, 0),
		minBuffer:        make(CandidateMinPQ, 0),
		maxBuffer:        make(CandidateMaxPQ, 0)}
}

// NewStateFactoryTrieDB creates a new stateFactory from Trie
func NewStateFactoryTrieDB(dbPath string) (StateFactory, error) {
	if len(dbPath) == 0 {
		return nil, nil
	}
	tr, err := trie.NewTrie(dbPath)
	if err != nil {
		return nil, err
	}
	return &stateFactory{
		trie:             tr,
		delegateHeap:     make(CandidateMinPQ, 0),
		candidateMinHeap: make(CandidateMinPQ, 0),
		candidateMaxHeap: make(CandidateMaxPQ, 0),
		minBuffer:        make(CandidateMinPQ, 0),
		maxBuffer:        make(CandidateMaxPQ, 0)}, nil
}

// CreateState adds a new State with initial balance to the factory
func (sf *stateFactory) CreateState(addr string, init uint64) (*State, error) {
	pubKeyHash := iotxaddress.GetPubkeyHash(addr)
	if pubKeyHash == nil {
		return nil, ErrInvalidAddr
	}
	balance := big.NewInt(0)
	balance.SetUint64(init)
	s := State{Address: addr, Balance: balance}
	mstate, err := stateToBytes(&s)
	if err != nil {
		return nil, err
	}
	if err := sf.trie.Upsert(pubKeyHash, mstate); err != nil {
		return nil, err
	}
	return &s, nil
}

// Balance returns balance
func (sf *stateFactory) Balance(addr string) (*big.Int, error) {
	state, err := sf.getState(addr)
	if err != nil {
		return nil, err
	}
	return state.Balance, nil
}

// Nonce returns the nonce if the account exists
func (sf *stateFactory) Nonce(addr string) (uint64, error) {
	state, err := sf.getState(addr)
	if err != nil {
		return 0, err
	}
	return state.Nonce, nil
}

// SetNonce returns the nonce if the account exists
func (sf *stateFactory) SetNonce(addr string, value uint64) error {
	state, err := sf.getState(addr)
	if err != nil {
		return err
	}
	state.Nonce = value

	mstate, err := stateToBytes(state)
	if err != nil {
		return err
	}
	if err := sf.trie.Upsert(iotxaddress.GetPubkeyHash(addr), mstate); err != nil {
		return err
	}
	return nil
}

// RootHash returns the hash of the root node of the trie
func (sf *stateFactory) RootHash() common.Hash32B {
	return sf.trie.RootHash()
}

// CommitStateChanges updates a State from the given actions
func (sf *stateFactory) CommitStateChanges(tsf []*trx.Transfer, vote []*trx.Vote) error {
	pending := make(map[common.PKHash]*State)
	for _, tx := range tsf {
		var pubKeyHash common.PKHash
		var err error
		// check sender
		pkhash := iotxaddress.GetPubkeyHash(tx.Sender)
		if pkhash == nil {
			return ErrInvalidAddr
		}
		copy(pubKeyHash[:], pkhash)
		sender, exist := pending[pubKeyHash]
		if !exist {
			if sender, err = sf.getState(tx.Sender); err != nil {
				return err
			}
			pending[pubKeyHash] = sender
		}
		if tx.Amount.Cmp(sender.Balance) == 1 {
			return ErrNotEnoughBalance
		}
		// update sender balance
		if err := sender.SubBalance(tx.Amount); err != nil {
			return err
		}
		// update sender nonce
		if tx.Nonce > sender.Nonce {
			sender.Nonce = tx.Nonce
		}
		// check recipient
		if pkhash = iotxaddress.GetPubkeyHash(tx.Recipient); pkhash == nil {
			return ErrInvalidAddr
		}
		copy(pubKeyHash[:], pkhash)
		recipient, exist := pending[pubKeyHash]
		if !exist {
			recipient, err = sf.getState(tx.Recipient)
			switch {
			case err == ErrAccountNotExist:
				if _, e := sf.CreateState(tx.Recipient, 0); e != nil {
					return e
				}
			case err != nil:
				return err
			}
			pending[pubKeyHash] = recipient
		}
		// update recipient balance
		if err := recipient.AddBalance(tx.Amount); err != nil {
			return err
		}
	}
	// construct <k, v> list of pending state
	transferK := [][]byte{}
	transferV := [][]byte{}
	for pkhash, state := range pending {
		ss, err := stateToBytes(state)
		if err != nil {
			return err
		}
		addr := make([]byte, len(pkhash))
		copy(addr, pkhash[:])
		transferK = append(transferK, addr)
		transferV = append(transferV, ss)
	}
	// commit the state changes to Trie in a batch
	return sf.trie.Commit(transferK, transferV)
}

// Delegates returns array of block producers
func (sf *stateFactory) Delegates() []*Candidate {
	return sf.delegateHeap.CandidateList()
}

// Candidates returns array of candidates
func (sf *stateFactory) Candidates() []*Candidate {
	return sf.candidateMinHeap.CandidateList()
}

// Buffer returns array of candidates in buffer
func (sf *stateFactory) Buffer() []*Candidate {
	return sf.minBuffer.CandidateList()
}

//======================================
// private functions
//=====================================
// getState pulls an existing State
func (sf *stateFactory) getState(addr string) (*State, error) {
	pubKeyHash := iotxaddress.GetPubkeyHash(addr)
	if pubKeyHash == nil {
		return nil, ErrInvalidAddr
	}
	mstate, err := sf.trie.Get(pubKeyHash)
	if errors.Cause(err) == trie.ErrNotExist {
		return nil, ErrAccountNotExist
	}
	if err != nil {
		return nil, err
	}
	state, err := bytesToState(mstate)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func (sf *stateFactory) updateVotes(candidate *Candidate, votes *big.Int) {
	candidate.Votes = votes
	c, level := sf.inPool(candidate.Address)
	if level == delegatePool { // if candidate is already in delegate pool
		sf.delegateHeap.update(c, candidate.Votes)
	} else if level == candidatePool { // if candidate is already in candidate pool
		sf.candidateMinHeap.update(c, candidate.Votes)
		sf.candidateMaxHeap.update(c, candidate.Votes)
	} else if level == buffer { // if candidate is already in buffer
		sf.minBuffer.update(c, candidate.Votes)
		sf.maxBuffer.update(c, candidate.Votes)
	} else { // candidate is not in any of three pools
		transitCandidate := candidate
		if len(sf.delegateHeap) == 0 || len(sf.delegateHeap) < delegateSize || candidate.Votes.Cmp(sf.delegateHeap.Top().(*Candidate).Votes) >= 0 {
			// Push candidate into delegate pool
			heap.Push(&sf.delegateHeap, transitCandidate)
			transitCandidate = nil
			if len(sf.delegateHeap) > delegateSize {
				transitCandidate = heap.Pop(&sf.delegateHeap).(*Candidate)
			}
		}
		if transitCandidate != nil &&
			(len(sf.candidateMinHeap) == 0 || len(sf.candidateMinHeap) < candidateSize || candidate.Votes.Cmp(sf.candidateMinHeap.Top().(*Candidate).Votes) >= 0) {
			// Push candidate into candidate pool
			heap.Push(&sf.candidateMinHeap, transitCandidate)
			heap.Push(&sf.candidateMaxHeap, transitCandidate)
			transitCandidate = nil
			if len(sf.candidateMinHeap) > candidateSize {
				transitCandidate = heap.Pop(&sf.candidateMinHeap).(*Candidate)
				heap.Remove(&sf.candidateMaxHeap, transitCandidate.maxIndex)
			}
		}
		if transitCandidate != nil &&
			(len(sf.minBuffer) == 0 || len(sf.minBuffer) < bufferSize || candidate.Votes.Cmp(sf.minBuffer.Top().(*Candidate).Votes) >= 0) {
			// Push candidate into buffer
			heap.Push(&sf.minBuffer, transitCandidate)
			heap.Push(&sf.maxBuffer, transitCandidate)
			transitCandidate = nil
			if len(sf.minBuffer) > bufferSize {
				transitCandidate = heap.Pop(&sf.minBuffer).(*Candidate)
				heap.Remove(&sf.maxBuffer, transitCandidate.maxIndex)
			}
		}
	}
	sf.balance()
}

func (sf *stateFactory) balance() {
	if len(sf.candidateMinHeap) > 0 && len(sf.maxBuffer) > 0 && sf.candidateMinHeap.Top().(*Candidate).Votes.Cmp(sf.maxBuffer.Top().(*Candidate).Votes) < 0 {
		cFromCandidatePool := heap.Pop(&sf.candidateMinHeap).(*Candidate)
		heap.Remove(&sf.candidateMaxHeap, cFromCandidatePool.maxIndex)
		cFromBuffer := heap.Pop(&sf.maxBuffer).(*Candidate)
		heap.Remove(&sf.minBuffer, cFromBuffer.minIndex)
		heap.Push(&sf.candidateMinHeap, cFromBuffer)
		heap.Push(&sf.candidateMaxHeap, cFromBuffer)
		heap.Push(&sf.minBuffer, cFromCandidatePool)
		heap.Push(&sf.maxBuffer, cFromCandidatePool)
	}
	if len(sf.delegateHeap) > 0 && len(sf.candidateMaxHeap) > 0 && sf.delegateHeap.Top().(*Candidate).Votes.Cmp(sf.candidateMaxHeap.Top().(*Candidate).Votes) < 0 {
		cFromDelegatePool := heap.Pop(&sf.delegateHeap).(*Candidate)
		cFromCandidatePool := heap.Pop(&sf.candidateMaxHeap).(*Candidate)
		heap.Remove(&sf.candidateMinHeap, cFromCandidatePool.minIndex)
		heap.Push(&sf.delegateHeap, cFromCandidatePool)
		heap.Push(&sf.candidateMinHeap, cFromDelegatePool)
		heap.Push(&sf.candidateMaxHeap, cFromDelegatePool)
	}
	if len(sf.candidateMinHeap) > 0 && len(sf.maxBuffer) > 0 && sf.candidateMinHeap.Top().(*Candidate).Votes.Cmp(sf.maxBuffer.Top().(*Candidate).Votes) < 0 {
		cFromCandidatePool := heap.Pop(&sf.candidateMinHeap).(*Candidate)
		heap.Remove(&sf.candidateMaxHeap, cFromCandidatePool.maxIndex)
		cFromBuffer := heap.Pop(&sf.maxBuffer).(*Candidate)
		heap.Remove(&sf.minBuffer, cFromBuffer.minIndex)
		heap.Push(&sf.candidateMinHeap, cFromBuffer)
		heap.Push(&sf.candidateMaxHeap, cFromBuffer)
		heap.Push(&sf.minBuffer, cFromCandidatePool)
		heap.Push(&sf.maxBuffer, cFromCandidatePool)
	}

	if len(sf.candidateMinHeap) != len(sf.candidateMaxHeap) || len(sf.minBuffer) != len(sf.maxBuffer) {
		fmt.Println("***************heap not sync***************")
	}
}

func (sf *stateFactory) inPool(address string) (*Candidate, int) {
	if c := sf.delegateHeap.exist(address); c != nil {
		return c, delegatePool // The candidate exists in the Delegate pool
	}
	if c := sf.candidateMinHeap.exist(address); c != nil {
		return c, candidatePool // The candidate exists in the Candidate pool
	}
	if c := sf.minBuffer.exist(address); c != nil {
		return c, buffer // The candidate exists in the buffer
	}
	return nil, 0
}
