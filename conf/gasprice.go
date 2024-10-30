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

package conf

import (
	"github.com/N42world/ast/params"
	"math/big"
)

var (
	DefaultMaxPrice    = big.NewInt(500 * params.GWei)
	DefaultIgnorePrice = big.NewInt(2 * params.Wei)
)

type GpoConfig struct {
	Blocks           int
	Percentile       int
	MaxHeaderHistory int
	MaxBlockHistory  int
	Default          *big.Int `toml:",omitempty"`
	MaxPrice         *big.Int `toml:",omitempty"`
	IgnorePrice      *big.Int `toml:",omitempty"`
}

// FullNodeGPO contains default gasprice oracle settings for full node.
var FullNodeGPO = GpoConfig{
	Blocks:           20,
	Percentile:       60,
	MaxHeaderHistory: 1024,
	MaxBlockHistory:  1024,
	MaxPrice:         DefaultMaxPrice,
	IgnorePrice:      DefaultIgnorePrice,
}

// LightClientGPO contains default gasprice oracle settings for light client.
var LightClientGPO = GpoConfig{
	Blocks:           2,
	Percentile:       60,
	MaxHeaderHistory: 300,
	MaxBlockHistory:  5,
	MaxPrice:         DefaultMaxPrice,
	IgnorePrice:      DefaultIgnorePrice,
}
