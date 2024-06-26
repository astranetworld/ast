// Copyright 2022 The astranet Authors
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

package conf

import (
	"github.com/astranetworld/ast/common/types"
	"github.com/holiman/uint256"

	"github.com/astranetworld/ast/params"
)

type Genesis struct {
	Config     *params.ChainConfig `json:"config" yaml:"config"`
	Nonce      uint64              `json:"nonce"`
	Timestamp  uint64              `json:"timestamp"`
	ExtraData  []byte              `json:"extraData"`
	GasLimit   uint64              `json:"gasLimit"   gencodec:"required"`
	Difficulty *uint256.Int        `json:"difficulty" gencodec:"required"`
	Mixhash    types.Hash          `json:"mixHash"`
	Coinbase   types.Address       `json:"coinbase"`

	//Engine *ConsensusConfig `json:"engine" yaml:"engine"`
	Miners []string     `json:"miners" yaml:"miners"`
	Alloc  GenesisAlloc `json:"alloc" yaml:"alloc"  gencodec:"required"`

	// These fields are used for consensus tests. Please don't use them
	// in actual genesis blocks.
	Number     uint64       `json:"number"`
	GasUsed    uint64       `json:"gasUsed"`
	ParentHash types.Hash   `json:"parentHash"`
	BaseFee    *uint256.Int `json:"baseFeePerGas"`
}

// GenesisAlloc specifies the initial state that is part of the genesis block.
type GenesisAlloc map[types.Address]GenesisAccount

type GenesisAccount struct {
	//Address string                    `json:"address" toml:"address"`
	Balance string                    `json:"balance"`
	Code    []byte                    `json:"code,omitempty"`
	Storage map[types.Hash]types.Hash `json:"storage,omitempty"`
	Nonce   uint64                    `json:"nonce,omitempty"`
}
