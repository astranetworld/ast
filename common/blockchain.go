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

package common

import (
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/n42blockchain/N42/common/block"
	"github.com/n42blockchain/N42/common/types"
	"github.com/n42blockchain/N42/internal/consensus"
	"github.com/n42blockchain/N42/modules/state"
	"github.com/n42blockchain/N42/params"
)

type IHeaderChain interface {
	GetHeaderByNumber(number *uint256.Int) block.IHeader
	GetHeaderByHash(h types.Hash) (block.IHeader, error)
	InsertHeader(headers []block.IHeader) (int, error)
	GetBlockByHash(h types.Hash) (block.IBlock, error)
	GetBlockByNumber(number *uint256.Int) (block.IBlock, error)
}

type IBlockChain interface {
	IHeaderChain
	Config() *params.ChainConfig
	CurrentBlock() block.IBlock
	Blocks() []block.IBlock
	Start() error
	GenesisBlock() block.IBlock
	NewBlockHandler(payload []byte, peer peer.ID) error
	InsertChain(blocks []block.IBlock) (int, error)
	InsertBlock(blocks []block.IBlock, isSync bool) (int, error)
	SetEngine(engine consensus.Engine)
	GetBlocksFromHash(hash types.Hash, n int) (blocks []block.IBlock)
	SealedBlock(b block.IBlock) error
	Engine() consensus.Engine
	GetReceipts(blockHash types.Hash) (block.Receipts, error)
	GetLogs(blockHash types.Hash) ([][]*block.Log, error)
	SetHead(head uint64) error
	AddFutureBlock(block block.IBlock) error

	GetHeader(types.Hash, *uint256.Int) block.IHeader
	// alias for GetBlocksFromHash?
	GetBlock(hash types.Hash, number uint64) block.IBlock
	StateAt(tx kv.Tx, blockNr uint64) *state.IntraBlockState

	GetTd(hash types.Hash, number *uint256.Int) *uint256.Int
	HasBlock(hash types.Hash, number uint64) bool

	DB() kv.RwDB
	Quit() <-chan struct{}

	Close() error

	WriteBlockWithState(block block.IBlock, receipts []*block.Receipt, ibs *state.IntraBlockState, nopay map[types.Address]*uint256.Int) error

	GetDepositInfo(address types.Address) (*uint256.Int, *uint256.Int)
	GetAccountRewardUnpaid(account types.Address) (*uint256.Int, error)
}

type IMiner interface {
	Start()
	PendingBlockAndReceipts() (block.IBlock, block.Receipts)
}
