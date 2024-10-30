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

package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/N42world/ast/common/account"
	common "github.com/N42world/ast/common/types"
	"github.com/N42world/ast/internal/node"
	"github.com/N42world/ast/log"
	"github.com/N42world/ast/modules"
	"github.com/N42world/ast/params"
	"github.com/N42world/ast/turbo/backup"
	"github.com/holiman/uint256"
	"github.com/ledgerwatch/erigon-lib/kv"
	"math/big"
	"os"
	"time"
)

var (
	exportCommand = &cli.Command{
		Name:        "export",
		Usage:       "Export N42 data",
		ArgsUsage:   "",
		Description: ``,
		Subcommands: []*cli.Command{
			{
				Name:      "txs",
				Usage:     "Export All N42 Transactions",
				ArgsUsage: "",
				Action:    exportTransactions,
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: ``,
			},
			{
				Name:      "balance",
				Usage:     "Export All N42 account balance",
				ArgsUsage: "",
				Action:    exportBalance,
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: ``,
			},
			{
				Name:      "dbState",
				Usage:     "Export All MDBX Buckets disk space",
				ArgsUsage: "",
				Action:    exportDBState,
				Flags: []cli.Flag{
					DataDirFlag,
				},
				Description: ``,
			},
			{
				Name:      "dbCopy",
				Usage:     "copy data from '--chaindata' to '--chaindata.to'",
				ArgsUsage: "",
				Action:    dbCopy,
				Flags: []cli.Flag{
					FromDataDirFlag,
					ToDataDirFlag,
				},
				Description: ``,
			},
		},
	}
)

func exportTransactions(ctx *cli.Context) error {

	stack, err := node.NewNode(ctx, &DefaultConfig)
	if err != nil {
		return err
	}

	blockChain := stack.BlockChain()
	defer stack.Close()

	currentBlock := blockChain.CurrentBlock()

	for i := uint64(0); i < currentBlock.Number64().Uint64(); i++ {

		block, err := blockChain.GetBlockByNumber(uint256.NewInt(i + 1))
		if err != nil {
			panic("cannot get block")
		}
		for _, transaction := range block.Transactions() {
			if transaction.To() == nil {
				continue
			}
			fmt.Printf("%d,%s,%s,%d,%s,%.2f\n",
				block.Number64().Uint64(),
				time.Unix(int64(block.Time()), 0).Format(time.RFC3339),
				transaction.From().Hex(),
				transaction.Nonce(),
				transaction.To().Hex(),
				new(big.Float).Quo(new(big.Float).SetInt(transaction.Value().ToBig()), new(big.Float).SetInt(big.NewInt(params.AMT))),
			)
		}
	}

	return nil
}

func exportBalance(ctx *cli.Context) error {

	stack, err := node.NewNode(ctx, &DefaultConfig)
	if err != nil {
		return err
	}
	db := stack.Database()
	defer stack.Close()

	roTX, err := db.BeginRo(ctx.Context)
	if err != nil {
		return err
	}
	defer roTX.Rollback()
	//kv.ReadAhead(ctx.Context, roTX.(kv.RoDB), atomic.NewBool(false), name, nil, 1<<32-1) // MaxUint32
	//
	srcC, err := roTX.Cursor("Account")
	if err != nil {
		return err
	}

	for k, v, err := srcC.First(); k != nil; k, v, err = srcC.Next() {
		if err != nil {
			return err
		}
		var acc account.StateAccount
		if err = acc.DecodeForStorage(v); err != nil {
			return err
		}

		fmt.Printf("%x, %.2f\n",
			k,
			new(big.Float).Quo(new(big.Float).SetInt(acc.Balance.ToBig()), new(big.Float).SetInt(big.NewInt(params.AMT))),
		)
	}

	return nil
}

func exportDBState(ctx *cli.Context) error {

	stack, err := node.NewNode(ctx, &DefaultConfig)
	if err != nil {
		return err
	}
	db := stack.Database()
	defer stack.Close()

	var tsize uint64

	roTX, err := db.BeginRo(ctx.Context)
	if err != nil {
		return err
	}
	defer roTX.Rollback()

	migrator, ok := roTX.(kv.BucketMigrator)
	if !ok {
		return fmt.Errorf("cannot open db as BucketMigrator")
	}
	Buckets, err := migrator.ListBuckets()
	for _, Bucket := range Buckets {
		size, _ := roTX.BucketSize(Bucket)
		tsize += size
		Cursor, _ := roTX.Cursor(Bucket)
		count, _ := Cursor.Count()
		Cursor.Close()
		if count != 0 {
			fmt.Printf("%30v count %10d size: %s \r\n", Bucket, count, common.StorageSize(size))
		}
	}
	fmt.Printf("total %s \n", common.StorageSize(tsize))
	return nil
}

func dbCopy(ctx *cli.Context) error {

	modules.AstInit()
	kv.ChaindataTablesCfg = modules.AstTableCfg

	fromChaindata := ctx.String(FromDataDirFlag.Name)
	toChaindata := ctx.String(ToDataDirFlag.Name)

	if f, err := os.Stat(fromChaindata); err != nil || !f.IsDir() {
		log.Errorf("fromChaindata do not exists or is not a dir, err: %s", err)
		return err
	}
	if f, err := os.Stat(toChaindata); err != nil || !f.IsDir() {
		log.Errorf("toChaindata do not exists or is not a dir, err: %s", err)
		return err
	}

	from, to := backup.OpenPair(fromChaindata, toChaindata, kv.ChainDB, 0)
	err := backup.Kv2kv(ctx.Context, from, to, nil, backup.ReadAheadThreads)
	if err != nil && !errors.Is(err, context.Canceled) {
		if !errors.Is(err, context.Canceled) {
			log.Error(err.Error())
		}
		return nil
	}

	return nil
}
