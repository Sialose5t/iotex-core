// Copyright (c) 2018 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided ‘as is’ and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package blockchain

import (
	"github.com/pkg/errors"

	"github.com/iotexproject/iotex-core/common"
	"github.com/iotexproject/iotex-core/common/service"
	"github.com/iotexproject/iotex-core/common/utils"
	"github.com/iotexproject/iotex-core/db"
)

const (
	blockNS                  = "blocks"
	blockHashHeightMappingNS = "hash<->height"
)

var (
	hashPrefix   = []byte("hash.")
	heightPrefix = []byte("height.")
	topHeightKey = []byte("top-height")
)

var (
	// ErrNotExist indicates certain item does not exist in Blockchain database
	ErrNotExist = errors.New("not exist in DB")
	// ErrAlreadyExist indicates certain item already exists in Blockchain database
	ErrAlreadyExist = errors.New("already exist in DB")
)

type blockDAO struct {
	service.CompositeService
	kvstore db.KVStore
}

// newBlockDAO instantiates a block DAO
func newBlockDAO(kvstore db.KVStore) *blockDAO {
	blockDAO := &blockDAO{kvstore: kvstore}
	blockDAO.AddService(kvstore)
	return blockDAO
}

// Start starts block DAO and initiates the top height if it doesn't exist
func (dao *blockDAO) Start() error {
	err := dao.CompositeService.Start()
	if err != nil {
		return errors.Wrap(err, "failed to start child services")
	}

	// set init height value
	err = dao.kvstore.PutIfNotExists(blockNS, topHeightKey, make([]byte, 8))
	if err != nil {
		return errors.Wrap(err, "failed to write initial value for top height")
	}
	return nil
}

// getBlockHash returns the block hash by height
func (dao *blockDAO) getBlockHash(height uint64) (common.Hash32B, error) {
	key := append(heightPrefix, utils.Uint64ToBytes(height)...)
	value, err := dao.kvstore.Get(blockHashHeightMappingNS, key)
	var hash common.Hash32B
	if err != nil {
		return hash, errors.Wrap(err, "failed to get block hash")
	}
	copy(hash[:], value)
	return hash, nil
}

// getBlockHeight returns the block height by hash
func (dao *blockDAO) getBlockHeight(hash common.Hash32B) (uint64, error) {
	key := append(hashPrefix, hash[:]...)
	value, err := dao.kvstore.Get(blockHashHeightMappingNS, key)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get block height")
	}
	if value == nil || len(value) == 0 {
		return 0, errors.Wrapf(ErrNotExist, "height not found for the block with hash = %x", hash)
	}
	return common.MachineEndian.Uint64(value), nil
}

// getBlock returns a block
func (dao *blockDAO) getBlock(hash common.Hash32B) (*Block, error) {
	value, err := dao.kvstore.Get(blockNS, hash[:])
	if err != nil {
		return nil, errors.Wrap(err, "failed to get block")
	}
	if len(value) == 0 {
		return nil, nil
	}
	blk := Block{}
	if err = blk.Deserialize(value); err != nil {
		return nil, errors.Wrap(err, "failed to deserialize block")
	}
	return &blk, nil
}

// getBlockchainHeight returns the blockchain height
func (dao *blockDAO) getBlockchainHeight() (uint64, error) {
	value, err := dao.kvstore.Get(blockNS, topHeightKey)
	if err != nil {
		return 0, errors.Wrap(err, "failed to get top height")
	}
	if value == nil || len(value) == 0 {
		return 0, errors.Wrap(ErrNotExist, "blockchain height is missing")
	}
	return common.MachineEndian.Uint64(value), nil
}

// putBlock puts a block
func (dao *blockDAO) putBlock(blk *Block) error {
	hash := blk.HashBlock()
	existingbBlk, err := dao.getBlock(hash)
	if err != nil {
		return err
	}
	if existingbBlk != nil {
		return errors.Wrapf(ErrAlreadyExist, "Block with hash = %x", hash)
	}

	height := utils.Uint64ToBytes(blk.Height())
	serialized, err := blk.Serialize()
	if err != nil {
		return errors.Wrap(err, "failed to serialize block")
	}
	if err = dao.kvstore.Put(blockNS, hash[:], serialized); err != nil {
		return errors.Wrap(err, "failed to put block")
	}
	hashKey := append(hashPrefix, hash[:]...)
	if err = dao.kvstore.Put(blockHashHeightMappingNS, hashKey, height); err != nil {
		return errors.Wrap(err, "failed to put hash -> height mapping")
	}
	heightKey := append(heightPrefix, height...)
	if err = dao.kvstore.Put(blockHashHeightMappingNS, heightKey, hash[:]); err != nil {
		return errors.Wrap(err, "failed to put height -> hash mapping")
	}
	value, err := dao.kvstore.Get(blockNS, topHeightKey)
	topHeight := common.MachineEndian.Uint64(value)
	if blk.Height() > topHeight {
		dao.kvstore.Put(blockNS, topHeightKey, height)
		if err != nil {
			return errors.Wrap(err, "failed to get top height")
		}
	}
	return nil
}
