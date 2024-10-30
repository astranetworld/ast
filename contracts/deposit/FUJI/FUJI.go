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

package fujideposit

import (
	"bytes"
	"embed"
	"github.com/N42world/ast/accounts/abi"
	"github.com/N42world/ast/common/crypto"
	"github.com/N42world/ast/common/hexutil"
	"github.com/N42world/ast/common/types"
	"github.com/N42world/ast/log"
	"github.com/holiman/uint256"
	"github.com/pkg/errors"
	"math/big"
)

//go:embed abi.json
var abiJson embed.FS
var contractAbi abi.ABI
var depositSignature = crypto.Keccak256Hash([]byte("DepositEvent(bytes,uint256,bytes)"))
var withdrawnSignature = crypto.Keccak256Hash([]byte("WithdrawnEvent(uint256)"))

func init() {
	var (
		depositAbiCode []byte
		err            error
	)
	if depositAbiCode, err = abiJson.ReadFile("abi.json"); err != nil {
		panic("Could not open abi.json")
	}

	if contractAbi, err = abi.JSON(bytes.NewReader(depositAbiCode)); err != nil {
		panic("unable to parse AMT deposit contract abi")
	}
}

type Contract struct {
}

func (c Contract) IsDepositAction(sigdata [4]byte) bool {
	var (
		method *abi.Method
		err    error
	)
	if method, err = contractAbi.MethodById(sigdata[:]); err != nil {
		return false
	}

	if !bytes.Equal(method.ID, contractAbi.Methods["deposit"].ID) {
		return false
	}
	return true
}

func (Contract) WithdrawnSignature() types.Hash {
	return withdrawnSignature
}

func (Contract) DepositSignature() types.Hash {
	return depositSignature
}

func (Contract) UnpackDepositLogData(data []byte) (publicKey []byte, signature []byte, depositAmount *uint256.Int, err error) {
	var (
		unpackedLogs []interface{}
		overflow     bool
	)
	//
	if unpackedLogs, err = contractAbi.Unpack("DepositEvent", data); err != nil {
		err = errors.Wrap(err, "unable to unpack logs")
		return
	}
	//
	if depositAmount, overflow = uint256.FromBig(unpackedLogs[1].(*big.Int)); overflow {
		err = errors.New("unable to unpack amount")
		return
	}
	publicKey, signature = unpackedLogs[0].([]byte), unpackedLogs[2].([]byte)
	log.Debug("unpacked DepositEvent Logs", "publicKey", hexutil.Encode(unpackedLogs[0].([]byte)), "signature", hexutil.Encode(unpackedLogs[2].([]byte)), "message", hexutil.Encode(depositAmount.Bytes()))
	return
}
