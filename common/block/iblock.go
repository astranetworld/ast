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

package block

import (
	"github.com/N42world/ast/common/transaction"
	"github.com/N42world/ast/common/types"
	"github.com/holiman/uint256"
	"google.golang.org/protobuf/proto"
)

type IHeader interface {
	Number64() *uint256.Int
	BaseFee64() *uint256.Int
	Hash() types.Hash
	ToProtoMessage() proto.Message
	FromProtoMessage(message proto.Message) error
	Marshal() ([]byte, error)
	Unmarshal(data []byte) error
	StateRoot() types.Hash
}

type IBody interface {
	Verifier() []*Verify
	Reward() []*Reward
	Transactions() []*transaction.Transaction
	ToProtoMessage() proto.Message
	FromProtoMessage(message proto.Message) error
}

type IBlock interface {
	IHeader
	Header() IHeader
	Body() IBody
	Transaction(hash types.Hash) *transaction.Transaction
	Transactions() []*transaction.Transaction
	Number64() *uint256.Int
	Difficulty() *uint256.Int
	Time() uint64
	GasLimit() uint64
	GasUsed() uint64
	Nonce() uint64
	Coinbase() types.Address
	ParentHash() types.Hash
	TxHash() types.Hash
	WithSeal(header IHeader) *Block
	//ToProtoMessage() proto.Message
	//FromProtoMessage(message proto.Message) error
}

type Blocks []IBlock
