package filedao

import (
	"context"
	"encoding/hex"
	"math/big"
	"strings"
	"testing"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/require"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"

	"github.com/iotexproject/iotex-core/blockchain/block"
	"github.com/iotexproject/iotex-core/config"
	"github.com/iotexproject/iotex-core/db"
	"github.com/iotexproject/iotex-core/pkg/compress"
	"github.com/iotexproject/iotex-core/testutil"
)

const (
	blockStoreBatchSize = 16
)

func TestNewFileDAOv2(t *testing.T) {
	r := require.New(t)

	testPath, err := testutil.PathOfTempFile("test-newfd")
	r.NoError(err)
	defer func() {
		testutil.CleanupPath(t, testPath)
	}()
	cfg := config.Default.DB
	r.Equal(compress.Snappy, cfg.Compressor)
	r.Equal(16, cfg.BlockStoreBatchSize)
	cfg.DbPath = testPath

	fd, err := newFileDAOv2(2, cfg)
	r.NoError(err)
	r.NotNil(fd)
	ctx := context.Background()
	r.NoError(fd.Start(ctx))
	defer fd.Stop(ctx)

	// new file does not use legacy's namespaces
	for _, v := range []string{
		blockNS,
		blockHeaderNS,
		blockBodyNS,
		blockFooterNS,
		receiptsNS,
	} {
		_, err = fd.kvStore.Get(v, []byte{})
		r.Error(err)
		r.True(strings.HasPrefix(err.Error(), "bucket = "+hex.EncodeToString([]byte(v))+" doesn't exist"))
	}

	// test counting index add empty transaction log
	ser := (&block.BlkTransactionLog{}).Serialize()
	for _, test := range []struct {
		compress string
		height   uint64
	}{
		{"", 0},
		{compress.Gzip, 1},
		{compress.Snappy, 2},
	} {
		data := ser
		if test.compress != "" {
			var err error
			data, err = compress.Compress(ser, test.compress)
			r.NoError(err)
		}
		r.NoError(addOneEntryToBatch(fd.hashStore, data, fd.batch))
		r.NoError(fd.kvStore.WriteBatch(fd.batch))
		v, err := fd.hashStore.Get(test.height)
		r.NoError(err)
		r.Equal(data, v)
		if test.compress != "" {
			v, err = compress.Decompress(v, test.compress)
		}
		r.NoError(err)
		r.Equal(ser, v)
	}
}

func TestNewFdInterface(t *testing.T) {
	testFdInterface := func(cfg config.DB, start uint64, t *testing.T) {
		r := require.New(t)

		testutil.CleanupPath(t, cfg.DbPath)
		fd, err := newFileDAOv2(start, cfg)
		r.NoError(err)

		ctx := context.Background()
		r.NoError(fd.Start(ctx))
		defer fd.Stop(ctx)

		height, err := fd.Bottom()
		r.NoError(err)
		r.Equal(start, height)
		height, err = fd.Height()
		r.NoError(err)
		r.Equal(start-1, height)

		// cannot commit height != tip+1
		builder := block.NewTestingBuilder()
		h := hash.ZeroHash256
		blk := createTestingBlock(builder, start-1, h)
		r.Equal(ErrInvalidTipHeight, fd.PutBlock(ctx, blk))
		blk = createTestingBlock(builder, start+1, h)
		r.Equal(ErrInvalidTipHeight, fd.PutBlock(ctx, blk))

		// commit blockStoreBatchSize blocks
		for i := uint64(0); i < fd.header.BlockStoreSize; i++ {
			blk = createTestingBlock(builder, start+i, h)
			r.NoError(fd.PutBlock(ctx, blk))
			h = blk.HashBlock()
			height, err = fd.Height()
			r.NoError(err)
			r.Equal(start+i, height)
			if i < fd.header.BlockStoreSize-1 {
				r.EqualValues(0, fd.lowestBlockOfStoreTip())
				r.Equal(start-1, fd.highestBlockOfStoreTip())
			} else {
				r.Equal(start, fd.lowestBlockOfStoreTip())
				r.Equal(start+fd.header.BlockStoreSize-1, fd.highestBlockOfStoreTip())
				r.Equal(start+fd.header.BlockStoreSize-1, height)
			}
		}

		// commit 3 more blocks
		for i := uint64(1); i <= 3; i++ {
			blk = createTestingBlock(builder, height+i, h)
			r.NoError(fd.PutBlock(ctx, blk))
			h = blk.HashBlock()
			r.Equal(start, fd.lowestBlockOfStoreTip())
			r.Equal(start+fd.header.BlockStoreSize-1, fd.highestBlockOfStoreTip())
		}
		height, err = fd.Height()
		r.NoError(err)
		r.Equal(start+fd.header.BlockStoreSize+2, height)

		// verify API for all blocks
		r.True(fd.ContainsTransactionLog())
		for i := height; i >= start; i-- {
			height, err = fd.Bottom()
			r.NoError(err)
			r.Equal(start, height)
			r.True(fd.ContainsHeight(i))
			height, err = fd.Height()
			r.NoError(err)
			r.Equal(i, height)
			h, err = fd.GetBlockHash(i)
			r.NoError(err)
			height, err = fd.GetBlockHeight(h)
			r.NoError(err)
			r.Equal(height, i)
			blk, err = fd.GetBlockByHeight(i)
			r.NoError(err)
			r.Equal(h, blk.HashBlock())
			receipt, err := fd.GetReceipts(i)
			r.NoError(err)
			r.EqualValues(1, receipt[0].Status)
			r.Equal(height, receipt[0].BlockHeight)
			r.Equal(blk.Header.PrevHash(), receipt[0].ActionHash)
			log, err := fd.TransactionLogs(i)
			r.NoError(err)
			l := log.Logs[0]
			r.Equal(receipt[0].ActionHash[:], l.ActionHash)
			r.EqualValues(1, l.NumTransactions)
			tx := l.Transactions[0]
			r.Equal(big.NewInt(100).String(), tx.Amount)
			r.Equal(hex.EncodeToString(l.ActionHash[:]), tx.Sender)
			r.Equal(hex.EncodeToString(l.ActionHash[:]), tx.Recipient)
			r.Equal(iotextypes.TransactionLogType_NATIVE_TRANSFER, tx.Type)

			// test DeleteTipBlock()
			r.NoError(fd.DeleteTipBlock())
			r.False(fd.ContainsHeight(i))
			_, err = fd.GetBlockHash(i)
			r.Equal(db.ErrNotExist, err)
			_, err = fd.GetBlockHeight(h)
			r.Equal(db.ErrNotExist, errors.Cause(err))
			_, err = fd.GetBlock(h)
			r.Equal(db.ErrNotExist, errors.Cause(err))
			_, err = fd.GetBlockByHeight(i)
			r.Equal(db.ErrNotExist, errors.Cause(err))
			_, err = fd.GetReceipts(i)
			r.Equal(db.ErrNotExist, errors.Cause(err))
			log, err = fd.TransactionLogs(i)
			r.Equal(ErrNotSupported, err)
		}

		// after deleting all blocks
		height, err = fd.Height()
		r.NoError(err)
		r.Equal(start-1, height)
		h, err = fd.GetBlockHash(height)
		r.NoError(err)
		r.Equal(hash.ZeroHash256, h)
		r.EqualValues(0, fd.lowestBlockOfStoreTip())
		r.Equal(start-1, fd.highestBlockOfStoreTip())
	}

	r := require.New(t)
	testPath, err := testutil.PathOfTempFile("test-interface")
	r.NoError(err)
	defer func() {
		testutil.CleanupPath(t, testPath)
	}()

	cfg := config.Default.DB
	cfg.DbPath = testPath
	_, err = newFileDAOv2(0, cfg)
	r.Equal(ErrNotSupported, err)

	for _, compress := range []string{"", compress.Snappy} {
		for _, start := range []uint64{1, 5, blockStoreBatchSize + 1, 4 * blockStoreBatchSize} {
			cfg.Compressor = compress
			t.Run("test fileDAOv2 interface", func(t *testing.T) {
				testFdInterface(cfg, start, t)
			})
		}
	}
}

func TestNewFdStart(t *testing.T) {
	testFdStart := func(cfg config.DB, start uint64, t *testing.T) {
		r := require.New(t)

		for _, num := range []uint64{3, blockStoreBatchSize - 1, blockStoreBatchSize, 2*blockStoreBatchSize - 1} {
			testutil.CleanupPath(t, cfg.DbPath)
			fd, err := newFileDAOv2(start, cfg)
			r.NoError(err)
			ctx := context.Background()
			r.NoError(fd.Start(ctx))
			defer fd.Stop(ctx)

			testCommitBlocks(t, fd, start, start+num-1, hash.ZeroHash256)
			height, err := fd.Height()
			r.NoError(err)
			r.Equal(start+num-1, height)
			r.NoError(fd.Stop(ctx))

			// start from existing file
			fd = openFileDAOv2(cfg)
			r.NoError(fd.Start(ctx))
			height, err = fd.Bottom()
			r.NoError(err)
			r.Equal(start, height)
			height, err = fd.Height()
			r.NoError(err)
			r.Equal(start+num-1, height)

			// verify API for all blocks
			for i := start; i < start+num; i++ {
				r.True(fd.ContainsHeight(i))
				h, err := fd.GetBlockHash(i)
				r.NoError(err)
				height, err = fd.GetBlockHeight(h)
				r.NoError(err)
				r.Equal(height, i)
				blk, err := fd.GetBlockByHeight(i)
				r.NoError(err)
				r.Equal(h, blk.HashBlock())
				receipt, err := fd.GetReceipts(i)
				r.NoError(err)
				r.EqualValues(1, receipt[0].Status)
				r.Equal(height, receipt[0].BlockHeight)
				r.Equal(blk.Header.PrevHash(), receipt[0].ActionHash)
				log, err := fd.TransactionLogs(i)
				r.NoError(err)
				r.NotNil(log)
				l := log.Logs[0]
				r.Equal(receipt[0].ActionHash[:], l.ActionHash)
				r.EqualValues(1, l.NumTransactions)
				tx := l.Transactions[0]
				r.Equal(big.NewInt(100).String(), tx.Amount)
				r.Equal(hex.EncodeToString(l.ActionHash[:]), tx.Sender)
				r.Equal(hex.EncodeToString(l.ActionHash[:]), tx.Recipient)
				r.Equal(iotextypes.TransactionLogType_NATIVE_TRANSFER, tx.Type)
			}
		}
	}

	r := require.New(t)
	testPath, err := testutil.PathOfTempFile("test-start")
	r.NoError(err)
	defer func() {
		testutil.CleanupPath(t, testPath)
	}()

	cfg := config.Default.DB
	cfg.DbPath = testPath
	for _, compress := range []string{"", compress.Gzip} {
		for _, start := range []uint64{1, 5, blockStoreBatchSize + 1, 4 * blockStoreBatchSize} {
			cfg.Compressor = compress
			t.Run("test fileDAOv2 start", func(t *testing.T) {
				testFdStart(cfg, start, t)
			})
		}
	}
}
