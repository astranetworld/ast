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

package main

import (
	"github.com/n42blockchain/N42/params"
	"math/big"
	"time"

	"github.com/n42blockchain/N42/conf"
)

var DefaultConfig = conf.Config{
	NodeCfg: conf.NodeConfig{
		NodePrivate: "",
		HTTP:        true,
		HTTPHost:    "127.0.0.1",
		HTTPPort:    "8545",
		IPCPath:     "ast.ipc",
		Miner:       false,
	},
	NetworkCfg: conf.NetWorkConfig{
		Bootstrapped: true,
	},
	LoggerCfg: conf.LoggerConfig{
		LogFile:    "./logger.log",
		Level:      "debug",
		MaxSize:    10,
		MaxBackups: 10,
		MaxAge:     30,
		Compress:   true,
	},
	PprofCfg: conf.PprofConfig{
		MaxCpu:     0,
		Port:       6060,
		TraceMutex: true,
		TraceBlock: true,
		Pprof:      false,
	},
	DatabaseCfg: conf.DatabaseConfig{
		DBType:     "lmdb",
		DBPath:     "chaindata",
		DBName:     "ast",
		SubDB:      []string{"chain"},
		Debug:      false,
		IsMem:      false,
		MaxDB:      100,
		MaxReaders: 1000,
	},
	MetricsCfg: conf.MetricsConfig{
		Port: 6060,
		HTTP: "127.0.0.1",
	},

	P2PCfg: &conf.P2PConfig{P2PLimit: &conf.P2PLimit{}},

	//GenesisCfg: ReadGenesis("allocs/mainnet.json"),
	GPO: conf.FullNodeGPO,
	Miner: conf.MinerConfig{
		GasCeil:  30000000,
		GasPrice: big.NewInt(params.GWei),
		Recommit: 4 * time.Second,
	},
}
