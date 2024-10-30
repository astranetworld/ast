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

package accounts

import (
	"bytes"
	"github.com/N42world/ast/common/crypto"
	"github.com/N42world/ast/common/math"
	"github.com/ledgerwatch/secp256k1"
	"testing"

	"github.com/N42world/ast/common/hexutil"
)

func TestTextHash(t *testing.T) {
	hash := TextHash([]byte("Hello Joe"))
	want := hexutil.MustDecode("0xa080337ae51c4e064c189e113edd0ba391df9206e2f49db658bb32cf2911730b")
	if !bytes.Equal(hash, want) {
		t.Fatalf("wrong hash: %x", hash)
	}
}

func TestSign(t *testing.T) {
	private, err := crypto.HexToECDSA("DEBF9EAE7820E23201EEE9D51413B6D2CDF06C320D7152C2D3BC1FB6C42DA23D")
	if nil != err {
		t.Error(err)
	}
	seckey := math.PaddedBigBytes(private.D, private.Params().BitSize/8)
	defer zeroBytes(seckey)

	msg, _ := hexutil.Decode("0x08712134afd46d42a45a5ed0e9311933138a06041b88062e590a613c2673c29f")
	signature, err := secp256k1.Sign(msg, seckey)

	t.Logf("%v", err)
	t.Logf("%s", hexutil.Encode(signature))
}

func zeroBytes(bytes []byte) {
	for i := range bytes {
		bytes[i] = 0
	}
}
