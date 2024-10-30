// Copyright 2022 The N42 Authors
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

package api

import (
	"github.com/N42world/ast/common/hexutil"
	"github.com/N42world/ast/common/types"
	"github.com/N42world/ast/internal/avm/common"
)

type DumpAccount struct {
	Balance   string                `json:"balance"`
	Nonce     uint64                `json:"nonce"`
	Root      hexutil.Bytes         `json:"root"`
	CodeHash  hexutil.Bytes         `json:"codeHash"`
	Code      hexutil.Bytes         `json:"code,omitempty"`
	Storage   map[types.Hash]string `json:"storage,omitempty"`
	Address   *common.Address       `json:"address,omitempty"` // Address only present in iterative (line-by-line) mode
	SecureKey hexutil.Bytes         `json:"key,omitempty"`     // If we don't have address, we can output the key

}
