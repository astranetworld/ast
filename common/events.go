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

package common

import (
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/n42blockchain/N42/common/block"
	"github.com/n42blockchain/N42/common/transaction"
	"github.com/n42blockchain/N42/modules/state"
)

// NewLocalTxsEvent local txs
type NewLocalTxsEvent struct{ Txs []*transaction.Transaction }

// NewTxsEvent txs
type NewTxsEvent struct{ Txs []*transaction.Transaction }

// NewLogsEvent new logs
type NewLogsEvent struct{ Logs []*block.Log }

// RemovedLogsEvent is posted when a reorg happens // todo blockchain v2
type RemovedLogsEvent struct{ Logs []*block.Log }

// NewPendingLogsEvent is posted when a reorg happens // todo miner v2
type NewPendingLogsEvent struct{ Logs []*block.Log }

// PeerJoinEvent Peer join
type PeerJoinEvent struct{ Peer peer.ID }

// PeerDropEvent Peer drop
type PeerDropEvent struct{ Peer peer.ID }

// DownloaderStartEvent start download
type DownloaderStartEvent struct{}

// DownloaderFinishEvent finish download
type DownloaderFinishEvent struct{}

type ChainHighestBlock struct {
	Block    block.Block
	Inserted bool
}
type MinedEntireEvent struct {
	Entire state.EntireCode
}
