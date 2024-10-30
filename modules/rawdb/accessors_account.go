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

package rawdb

import (
	"github.com/N42world/ast/common/account"
	"github.com/N42world/ast/common/types"
	"github.com/N42world/ast/modules"
	"github.com/ledgerwatch/erigon-lib/kv"
)

func GetAccount(db kv.Tx, addr types.Address, acc *account.StateAccount) (bool, error) {
	enc, err := db.GetOne(modules.Account, addr[:])
	if err != nil {
		return false, err
	}
	if len(enc) == 0 {
		return false, nil
	}
	if err = acc.DecodeForStorage(enc); err != nil {
		return false, err
	}
	return true, nil
}
