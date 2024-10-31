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

import (
	"context"
	"github.com/n42blockchain/N42/log"
	"net"
)

func (s *Server) ServeListener(l net.Listener) error {
	for {
		conn, err := l.Accept()
		if IsTemporaryError(err) {
			log.Warn("RPC accept error", "err", err)
			continue
		} else if err != nil {
			return err
		}
		log.Debug("Accepted RPC connection", "conn", conn.RemoteAddr())
		go s.ServeCodec(NewCodec(conn), 0)
	}
}

func DialIPC(ctx context.Context, endpoint string) (*Client, error) {
	return newClient(ctx, func(ctx context.Context) (ServerCodec, error) {
		conn, err := newIPCConnection(ctx, endpoint)
		if err != nil {
			return nil, err
		}
		return NewCodec(conn), err
	})
}
