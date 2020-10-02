// Copyright (c) 2020 IoTeX Foundation
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package filedao

import (
	"context"
	"encoding/hex"
	"os"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/iotexproject/go-pkgs/crypto"
	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-core/blockchain/block"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/pkg/compress"
)

func TestChecksumNamespaceAndKeys(t *testing.T) {
	r := require.New(t)

	a := []hash.Hash256{
		// filedao
		hash.BytesToHash256([]byte(blockHashHeightMappingNS)),
		hash.BytesToHash256([]byte(systemLogNS)),
		hash.BytesToHash256(topHeightKey),
		hash.BytesToHash256(topHashKey),
		hash.BytesToHash256(hashPrefix),
		// filedao_legacy
		hash.BytesToHash256([]byte(blockNS)),
		hash.BytesToHash256([]byte(blockHeaderNS)),
		hash.BytesToHash256([]byte(blockBodyNS)),
		hash.BytesToHash256([]byte(blockFooterNS)),
		hash.BytesToHash256([]byte(receiptsNS)),
		hash.BytesToHash256(heightPrefix),
		hash.BytesToHash256(heightToFileBucket),
		// filedao_v2
		hash.BytesToHash256([]byte(FileV2)),
		hash.BytesToHash256([]byte{16}),
		hash.BytesToHash256([]byte(compress.Gzip)),
		hash.BytesToHash256([]byte(compress.Snappy)),
		hash.BytesToHash256([]byte(hashDataNS)),
		hash.BytesToHash256([]byte(blockDataNS)),
		hash.BytesToHash256([]byte(headerDataNs)),
		hash.BytesToHash256(fileHeaderKey),
	}

	checksum := crypto.NewMerkleTree(a)
	r.NotNil(checksum)
	h := checksum.HashTree()
	r.Equal("18747e1ac5364ce3f398e03092f159121b55166449657f65ba1f9243e8830391", hex.EncodeToString(h[:]))
}

func TestReadFileHeader(t *testing.T) {
	r := require.New(t)

	cfg := config.Default.DB
	cfg.DbPath = "./filedao_v2.db"

	// test non-existing file
	h, err := readFileHeader(cfg.DbPath, FileLegacyMaster)
	r.Equal(ErrFileNotExist, err)
	h, err = readFileHeader(cfg.DbPath, FileAll)
	r.Equal(ErrFileNotExist, err)

	// empty legacy file is invalid
	legacy, err := newFileDAOLegacy(cfg)
	r.NoError(err)
	ctx := context.Background()
	r.NoError(legacy.Start(ctx))
	r.NoError(legacy.Stop(ctx))
	h, err = readFileHeader(cfg.DbPath, FileLegacyMaster)
	r.Equal(ErrFileInvalid, err)
	h, err = readFileHeader(cfg.DbPath, FileAll)
	r.Equal(ErrFileInvalid, err)

	// commit 1 block to make it a valid legacy file
	r.NoError(legacy.Start(ctx))
	builder := block.NewTestingBuilder()
	blk := createTestingBlock(builder, 1, hash.ZeroHash256)
	r.NoError(legacy.PutBlock(ctx, blk))
	height, err := legacy.Height()
	r.NoError(err)
	r.EqualValues(1, height)
	r.NoError(legacy.Stop(ctx))

	type testCheckFile struct {
		checkType, version string
		err                error
	}

	// test valid legacy master file
	test1 := []testCheckFile{
		{FileLegacyMaster, FileLegacyMaster, nil},
		{FileLegacyAuxiliary, FileLegacyMaster, nil},
		{FileV2, "", ErrFileInvalid},
		{FileAll, FileLegacyMaster, nil},
	}
	for _, v := range test1 {
		h, err = readFileHeader(cfg.DbPath, v.checkType)
		r.Equal(v.err, err)
		if err == nil {
			r.Equal(v.version, h.Version)
		}
	}
	os.RemoveAll(cfg.DbPath)

	// test valid v2 master file
	r.NoError(createNewV2File(1, cfg))
	defer os.RemoveAll(cfg.DbPath)

	test2 := []testCheckFile{
		{FileLegacyMaster, "", ErrFileInvalid},
		{FileLegacyAuxiliary, "", ErrFileInvalid},
		{FileV2, FileV2, nil},
		{FileAll, FileV2, nil},
	}
	for _, v := range test2 {
		h, err = readFileHeader(cfg.DbPath, v.checkType)
		r.Equal(v.err, err)
		if err == nil {
			r.Equal(v.version, h.Version)
		}
	}

	r.Panics(func() { readFileHeader(cfg.DbPath, "") })
}

func TestNewFileDAOSplitV2(t *testing.T) {
	r := require.New(t)

	cfg := config.Default.DB
	cfg.DbPath = "./filedao_v2.db"
	defer os.RemoveAll(cfg.DbPath)

	// test non-existing file
	h, err := checkMasterChainDBFile(cfg.DbPath)
	r.Equal(ErrFileNotExist, err)

	// test empty db file, this will create new v2 file
	fd, err := NewFileDAO(cfg)
	r.NoError(err)
	r.NotNil(fd)
	h, err = readFileHeader(cfg.DbPath, FileAll)
	r.NoError(err)
	r.Equal(FileV2, h.Version)
	ctx := context.Background()
	r.NoError(fd.Start(ctx))
	r.NoError(testCommitBlocks(t, fd, 1, 10, hash.ZeroHash256))
	testVerifyChainDB(t, fd, 1, 10)

	// use test FileDAO that will split at height 20 and 40
	testFd := newTestSplitFile(fd.(*fileDAO), []uint64{20, 40})
	r.NoError(testCommitBlocks(t, testFd, 11, 45, hash.ZeroHash256))
	testVerifyChainDB(t, testFd, 1, 45)
	r.NoError(fd.Stop(ctx))
	top, files := checkAuxFiles(cfg.DbPath, FileV2)
	r.EqualValues(2, top)
	r.Equal(2, len(files))
	file1 := kthAuxFileName("./filedao_v2.db", 1)
	file2 := kthAuxFileName("./filedao_v2.db", 2)
	r.Equal(files[0], file1)
	r.Equal(files[1], file2)
	os.RemoveAll(file1)
	os.RemoveAll(file2)
}

func TestNewFileDAOSplitLegacy(t *testing.T) {
	r := require.New(t)

	cfg := config.Default.DB
	cfg.DbPath = "./filedao_v2.db"
	defer os.RemoveAll(cfg.DbPath)

	cfg.SplitDBHeight = 5
	cfg.SplitDBSizeMB = 20
	fd, err := newFileDAOLegacy(cfg)
	r.NoError(err)
	ctx := context.Background()
	r.NoError(fd.Start(ctx))
	r.NoError(testCommitBlocks(t, fd, 1, 10, hash.ZeroHash256))
	testVerifyChainDB(t, fd, 1, 10)
	r.NoError(fd.Stop(ctx))
	// block 1~5 in default file.db, block 6~10 in file-000000001.db
	file1 := kthAuxFileName("./filedao_v2.db", 1)
	defer os.RemoveAll(file1)

	// now use test FileDAO that will split at height 20, 30 and 50
	fd, err = NewFileDAO(cfg)
	r.NoError(err)
	r.EqualValues(1, fd.(*fileDAO).topIndex)
	r.NoError(fd.Start(ctx))
	testFd := newTestSplitFile(fd.(*fileDAO), []uint64{20, 30, 50})
	r.NoError(testCommitBlocks(t, testFd, 11, 28, hash.ZeroHash256))
	// skip block 29, commit block 30 which is a split height
	r.Equal(ErrInvalidTipHeight, testCommitBlocks(t, testFd, 30, 55, hash.ZeroHash256))
	r.NoError(testCommitBlocks(t, testFd, 29, 55, hash.ZeroHash256))
	testVerifyChainDB(t, testFd, 1, 55)
	r.NoError(fd.Stop(ctx))

	// now we should have:
	// block 1~5 in legacy file.db
	// block 6~20 in legacy file-000000001.db
	// block 21~30 in v2 file-000000002.db
	// block 31~50 in v2 file-000000003.db
	// block 51~55 in v2 file-000000004.db
	file2 := kthAuxFileName("./filedao_v2.db", 2)
	file3 := kthAuxFileName("./filedao_v2.db", 3)
	file4 := kthAuxFileName("./filedao_v2.db", 4)
	defer os.RemoveAll(file2)
	defer os.RemoveAll(file3)
	defer os.RemoveAll(file4)
	h, err := readFileHeader(cfg.DbPath, FileAll)
	r.NoError(err)
	r.Equal(FileLegacyMaster, h.Version)
	h, err = readFileHeader(file1, FileLegacyAuxiliary)
	r.NoError(err)
	r.Equal(FileLegacyAuxiliary, h.Version)
	h, err = readFileHeader(file2, FileV2)
	r.NoError(err)
	r.Equal(FileV2, h.Version)
	h, err = readFileHeader(file3, FileV2)
	r.NoError(err)
	r.Equal(FileV2, h.Version)
	h, err = readFileHeader(file4, FileV2)
	r.NoError(err)
	r.Equal(FileV2, h.Version)
	top, files := checkAuxFiles(cfg.DbPath, FileLegacyAuxiliary)
	r.EqualValues(1, top)
	r.Equal(1, len(files))
	r.Equal(files[0], file1)
	top, files = checkAuxFiles(cfg.DbPath, FileV2)
	r.EqualValues(4, top)
	r.Equal(3, len(files))
	r.Equal(files[0], file2)
	r.Equal(files[1], file3)
	r.Equal(files[2], file4)

	// open 4 db files and verify again
	fd, err = NewFileDAO(cfg)
	r.NoError(err)
	r.EqualValues(4, fd.(*fileDAO).topIndex)
	r.NoError(fd.Start(ctx))
	defer fd.Stop(ctx)
	testVerifyChainDB(t, fd, 1, 55)
}

func TestCheckFiles(t *testing.T) {
	r := require.New(t)

	auxTests := []struct {
		file, base string
		index      uint64
		ok         bool
	}{
		{"/tmp/chain-00000003.db", "/tmp/chain", 0, false},
		{"/tmp/chain-00000003", "/tmp/chain.db", 0, false},
		{"/tmp/chain-00000003.dat", "/tmp/chain.db", 0, false},
		{"/tmp/chair-00000003.dat", "/tmp/chain.db", 0, false},
		{"/tmp/chain=00000003.db", "/tmp/chain.db", 0, false},
		{"/tmp/chain-0000003.db", "/tmp/chain.db", 0, false},
		{"/tmp/chain--0000003.db", "/tmp/chain.db", 0, false},
		{"/tmp/chain-00000003.db", "/tmp/chain.db", 3, true},
	}

	for _, v := range auxTests {
		index, ok := isAuxFile(v.file, v.base)
		r.Equal(v.index, index)
		r.Equal(v.ok, ok)
	}

	cfg := config.Default.DB
	cfg.DbPath = "./filedao_v2.db"
	_, files := checkAuxFiles(cfg.DbPath, FileLegacyAuxiliary)
	r.Nil(files)
	_, files = checkAuxFiles(cfg.DbPath, FileV2)
	r.Nil(files)

	// create 3 v2 files
	for i := 1; i <= 3; i++ {
		cfg.DbPath = kthAuxFileName("./filedao_v2.db", uint64(i))
		r.NoError(createNewV2File(1, cfg))
	}
	defer func() {
		for i := 1; i <= 3; i++ {
			os.RemoveAll(kthAuxFileName("./filedao_v2.db", uint64(i)))
		}
	}()
	cfg.DbPath = "./filedao_v2.db"
	top, files := checkAuxFiles(cfg.DbPath, FileV2)
	r.EqualValues(3, top)
	r.Equal(3, len(files))
	for i := 1; i <= 3; i++ {
		r.Equal(files[i-1], kthAuxFileName("./filedao_v2.db", uint64(i)))
	}
}
