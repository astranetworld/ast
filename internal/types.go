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

package internal

import (
	"github.com/astranetworld/ast/common/block"
	"github.com/astranetworld/ast/common/types"
	"github.com/astranetworld/ast/modules/state"
	"github.com/holiman/uint256"
)

// Validator is an interface which defines the standard for block validation. It
// is only responsible for validating block contents, as the header validation is
// done by the specific consensus engines.
//type Validator interface {
//	// ValidateBody validates the given block's content.
//	ValidateBody(block *block.Block) error
//
//	// ValidateState validates the given statedb and optionally the receipts and
//	// gas used.
//	ValidateState(block *block.Block, state *state.StateDB, receipts block.Receipts, usedGas uint64) error
//}
//
//// Prefetcher is an interface for pre-caching transaction signatures and state.
//type Prefetcher interface {
//	// Prefetch processes the state changes according to the Ethereum rules by running
//	// the transaction messages using the statedb, but any changes are discarded. The
//	// only goal is to pre-cache transaction signatures and state trie nodes.
//	Prefetch(block *block.Block, statedb *state.StateDB, cfg vm.Config, interrupt *uint32)
//}

// Processor is an interface for processing blocks using a given initial state.
type Processor interface {
	// Process processes the state changes according to the Ethereum rules by running
	// the transaction messages using the statedb and applying any rewards to both
	// the processor (coinbase) and any included uncles.
	Process(b *block.Block, ibs *state.IntraBlockState, stateReader state.StateReader, stateWriter state.WriterWithChangeSets, blockHashFunc func(n uint64) types.Hash) (block.Receipts, map[types.Address]*uint256.Int, []*block.Log, uint64, error)
}
