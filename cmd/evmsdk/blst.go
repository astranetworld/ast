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

package evmsdk

import (
	"encoding/hex"

	"github.com/astranetworld/ast/common/crypto/bls"
)

func BlsSign(privKey, msg string) (interface{}, error) {
	msgBytes, err := hex.DecodeString(msg)
	if err != nil {
		return nil, err
	}
	privKeyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	arr := [32]byte{}
	copy(arr[:], privKeyBytes[:])
	sk, err := bls.SecretKeyFromRandom32Byte(arr)
	if err != nil {
		return nil, err
	}
	resBytes := sk.Sign(msgBytes).Marshal()
	return hex.EncodeToString(resBytes), nil
}

func BlsPublicKey(privKey string) (interface{}, error) {
	privKeyBytes, err := hex.DecodeString(privKey)
	if err != nil {
		return nil, err
	}
	arr := [32]byte{}
	copy(arr[:], privKeyBytes[:])
	sk, err := bls.SecretKeyFromRandom32Byte(arr)
	if err != nil {
		return nil, err
	}
	resBytes := sk.PublicKey().Marshal()
	return hex.EncodeToString(resBytes), nil
}
