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

package jsonrpc

func init() {

}

//func TestHttpClient(t *testing.T) {
//	client, _ := Dial("http://127.0.0.1:8545")
//	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
//	defer cancel()
//	m := make(map[string]string, 0)
//	if err := client.CallContext(ctx, &m, "rpc_modules"); err != nil {
//		fmt.Println("can't get rpc modules:", err)
//		return
//	}
//	if len(m) == 0 {
//		t.Fail()
//	}
//}

//func TestIPC(t *testing.T) {
//	client, _ := Dial("")
//	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
//	defer cancel()
//	m := make(map[string]string, 0)
//	if err := client.CallContext(ctx, &m, "rpc_modules"); err != nil {
//		fmt.Println("can't get rpc modules:", err)
//		return
//	}
//	if len(m) == 0 {
//		t.Fail()
//	}
//}
