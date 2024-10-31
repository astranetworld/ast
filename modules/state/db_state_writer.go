// Copyright 2023 The N42 Authors
// This file is part of the N42 library.
//
// The N42 library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The N42 library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the N42 library. If not, see <http://www.gnu.org/licenses/>.

package state

import (
	"bytes"
	"fmt"
	"github.com/n42blockchain/N42/common/account"
	"github.com/n42blockchain/N42/common/crypto"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/modules"
	"math"

	"github.com/n42blockchain/N42/modules/changeset"

	"github.com/RoaringBitmap/roaring/roaring64"

	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/n42blockchain/N42/modules/ethdb/bitmapdb"
)

func originalAccountData(original *account.StateAccount, omitHashes bool) []byte {
	var originalData []byte
	if !original.Initialised {
		originalData = []byte{}
	} else if omitHashes {
		testAcc := original.SelfCopy()
		copy(testAcc.CodeHash[:], emptyCodeHash)
		testAcc.Root = crypto.Keccak256Hash(nil)
		originalDataLen := testAcc.EncodingLengthForStorage()
		originalData = make([]byte, originalDataLen)
		testAcc.EncodeForStorage(originalData)
	} else {
		originalDataLen := original.EncodingLengthForStorage()
		originalData = make([]byte, originalDataLen)
		original.EncodeForStorage(originalData)
	}
	return originalData
}

func writeIndex(blocknum uint64, changes *changeset.ChangeSet, bucket string, changeDb kv.RwTx) error {
	buf := bytes.NewBuffer(nil)
	for _, change := range changes.Changes {
		k := modules.CompositeKeyWithoutIncarnation(change.Key)

		index, err := bitmapdb.Get64(changeDb, bucket, k, math.MaxUint32, math.MaxUint32)
		if err != nil {
			return fmt.Errorf("find chunk failed: %w", err)
		}
		index.Add(blocknum)
		if err = bitmapdb.WalkChunkWithKeys64(k, index, bitmapdb.ChunkLimit, func(chunkKey []byte, chunk *roaring64.Bitmap) error {
			buf.Reset()
			if _, err = chunk.WriteTo(buf); err != nil {
				return err
			}
			return changeDb.Put(bucket, chunkKey, types.CopyBytes(buf.Bytes()))
		}); err != nil {
			return err
		}
	}

	return nil
}
