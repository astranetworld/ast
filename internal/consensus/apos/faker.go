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

package apos

import (
	"github.com/astranetworld/ast/common/block"
	"github.com/astranetworld/ast/common/transaction"
	"github.com/astranetworld/ast/common/types"
	"github.com/astranetworld/ast/internal/consensus"
	"github.com/astranetworld/ast/modules/rpc/jsonrpc"
	"github.com/astranetworld/ast/modules/state"
	"github.com/astranetworld/ast/params"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
)

type Faker struct{}

func (f Faker) Author(header block.IHeader) (types.Address, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyHeader(chain consensus.ChainHeaderReader, header block.IHeader, seal bool) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyHeaders(chain consensus.ChainHeaderReader, headers []block.IHeader, seals []bool) (chan<- struct{}, <-chan error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) VerifyUncles(chain consensus.ChainReader, block block.IBlock) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Prepare(chain consensus.ChainHeaderReader, header block.IHeader) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Finalize(chain consensus.ChainHeaderReader, header block.IHeader, state *state.IntraBlockState, txs []*transaction.Transaction, uncles []block.IHeader) ([]*block.Reward, map[types.Address]*uint256.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) FinalizeAndAssemble(chain consensus.ChainHeaderReader, header block.IHeader, state *state.IntraBlockState, txs []*transaction.Transaction, uncles []block.IHeader, receipts []*block.Receipt) (block.IBlock, []*block.Reward, map[types.Address]*uint256.Int, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Rewards(tx kv.RwTx, header block.IHeader, state *state.IntraBlockState, setRewards bool) ([]*block.Reward, error) {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Seal(chain consensus.ChainHeaderReader, block block.IBlock, results chan<- block.IBlock, stop <-chan struct{}) error {
	//TODO implement me
	panic("implement me")
}

func (f Faker) SealHash(header block.IHeader) types.Hash {
	//TODO implement me
	panic("implement me")
}

func (f Faker) CalcDifficulty(chain consensus.ChainHeaderReader, time uint64, parent block.IHeader) *uint256.Int {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Type() params.ConsensusType {
	return params.Faker
}

func (f Faker) APIs(chain consensus.ChainReader) []jsonrpc.API {
	//TODO implement me
	panic("implement me")
}

func (f Faker) Close() error {
	//TODO implement me
	panic("implement me")
}

func NewFaker() consensus.Engine {
	return &Faker{}
}

func (f Faker) IsServiceTransaction(sender types.Address, syscall consensus.SystemCall) bool {
	return false
}
