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

package transaction

import (
	"github.com/holiman/uint256"
	"github.com/n42blockchain/N42/common/types"
)

type SakuragiTx struct {
	Nonce    uint64      // nonce of sender account
	GasPrice uint256.Int // wei per gas
	Gas      uint64      // gas limit
	To       *types.Address
	From     *types.Address
	Value    uint256.Int // wei amount
	Data     []byte      // contract invocation input data
	Sign     []byte      // signature values
}
