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

package main

import (
	"context"
	"fmt"
	common2 "github.com/astranetworld/ast/common"
	block2 "github.com/astranetworld/ast/common/block"
	"github.com/astranetworld/ast/common/transaction"
	"github.com/astranetworld/ast/common/types"
	"github.com/astranetworld/ast/internal"
	"github.com/astranetworld/ast/internal/consensus/apos"
	"github.com/astranetworld/ast/internal/consensus/misc"
	"github.com/astranetworld/ast/internal/vm"
	"github.com/astranetworld/ast/modules/ethdb/olddb"
	"github.com/astranetworld/ast/modules/state"
	"github.com/astranetworld/ast/params"
	"unsafe"
)

func verify(ctx context.Context, msg *state.EntireCode) types.Hash {
	codeMap := make(map[types.Hash][]byte)
	for _, pair := range msg.Codes {
		codeMap[pair.Hash] = pair.Code
	}

	readCodeF := func(hash types.Hash) ([]byte, error) {
		if code, ok := codeMap[hash]; ok {
			return code, nil
		}
		return nil, nil
	}

	hashMap := make(map[uint64]types.Hash)
	for _, h := range msg.Headers {
		hashMap[h.Number.Uint64()] = h.Hash()
	}
	getNumberHash := func(n uint64) types.Hash {
		if hash, ok := hashMap[n]; ok {
			return hash
		}
		return types.Hash{}
	}

	var txs transaction.Transactions
	for _, tByte := range msg.Entire.Transactions {
		tmp := &transaction.Transaction{}
		if err := tmp.Unmarshal(tByte); nil != err {
			panic(err)
		}
		txs = append(txs, tmp)
	}

	body := &block2.Body{
		Txs: txs,
	}

	block := block2.NewBlockFromStorage(msg.Entire.Header.Hash(), msg.Entire.Header, body)
	batch := olddb.NewHashBatch(nil, ctx.Done(), "")
	defer batch.Rollback()
	old := make(map[string][]byte, len(msg.Entire.Snap.Items))
	for _, v := range msg.Entire.Snap.Items {
		old[*(*string)(unsafe.Pointer(&v.Key))] = v.Value
	}
	stateReader := olddb.NewStateReader(old, nil, batch, block.Number64().Uint64())
	stateReader.SetReadCodeF(readCodeF)
	ibs := state.New(stateReader)
	ibs.SetSnapshot(msg.Entire.Snap)
	ibs.SetHeight(block.Number64().Uint64())
	ibs.SetGetOneFun(batch.GetOne)

	root, err := checkBlock(getNumberHash, block, ibs, msg.CoinBase, msg.Rewards)
	if nil != err {
		panic(err)
	}
	return root
}

func checkBlock(getHashF func(n uint64) types.Hash, block *block2.Block, ibs *state.IntraBlockState, coinbase types.Address, rewards []*block2.Reward) (types.Hash, error) {
	header := block.Header().(*block2.Header)
	chainConfig := params.MainnetChainConfig
	if chainConfig.DAOForkSupport && chainConfig.DAOForkBlock != nil && chainConfig.DAOForkBlock.Cmp(block.Number64().ToBig()) == 0 {
		misc.ApplyDAOHardFork(ibs)
	}
	noop := state.NewNoopWriter()

	usedGas := new(uint64)
	gp := new(common2.GasPool)
	gp.AddGas(block.GasLimit())
	cfg := vm.Config{}
	//cfg := vm.Config{Debug: true, Tracer: logger.NewMarkdownLogger(nil, os.Stdout)}

	engine := apos.NewFaker()
	for i, tx := range block.Transactions() {
		ibs.Prepare(tx.Hash(), block.Hash(), i)
		_, _, err := internal.ApplyTransaction(chainConfig, getHashF, engine, &coinbase, gp, ibs, noop, header, tx, usedGas, cfg)
		if err != nil {

			return types.Hash{}, err
		}
	}

	if !cfg.StatelessExec && *usedGas != header.GasUsed {
		return types.Hash{}, fmt.Errorf("gas used by execution: %d, in header: %d", *usedGas, header.GasUsed)
	}

	if len(rewards) > 0 {
		for _, reward := range rewards {
			if reward.Amount != nil && !reward.Amount.IsZero() {
				if !ibs.Exist(reward.Address) {
					ibs.CreateAccount(reward.Address, false)
				}
				ibs.AddBalance(reward.Address, reward.Amount)
			}
		}
		ibs.SoftFinalise()
	}

	return ibs.IntermediateRoot(), nil
}
