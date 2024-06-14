// Copyright 2023 The astranet Authors
// This file is part of the astranet library.
//
// The astranet library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The astranet library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the astranet library. If not, see <http://www.gnu.org/licenses/>.

package rpchelper

import (
	"fmt"
	"github.com/astranetworld/ast/modules/rawdb"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
)

func GetLatestBlockNumber(tx kv.Tx) (*uint256.Int, error) {
	current := rawdb.ReadCurrentBlock(tx)
	if current == nil {
		return nil, fmt.Errorf("cannot get current block")
	}
	return current.Number64(), nil
}

func GetFinalizedBlockNumber(tx kv.Tx) (*uint256.Int, error) {
	//todo
	return GetLatestBlockNumber(tx)
}

func GetSafeBlockNumber(tx kv.Tx) (*uint256.Int, error) {
	//todo
	return GetLatestBlockNumber(tx)
}
