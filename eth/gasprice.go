// Copyright 2015 The go-expanse Authors
// This file is part of the go-expanse library.
//
// The go-expanse library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-expanse library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-expanse library. If not, see <http://www.gnu.org/licenses/>.

package exp

import (
	"math/big"
	"math/rand"
	"sync"

	"github.com/expanse-project/go-expanse/core"
	"github.com/expanse-project/go-expanse/core/types"
	"github.com/expanse-project/go-expanse/event"
	"github.com/expanse-project/go-expanse/logger"
	"github.com/expanse-project/go-expanse/logger/glog"
)

const gpoProcessPastBlocks = 100

type blockPriceInfo struct {
	baseGasPrice *big.Int
}

type GasPriceOracle struct {
	exp                           *Expanse
	chain                         *core.ChainManager
	events                        event.Subscription
	blocks                        map[uint64]*blockPriceInfo
	firstProcessed, lastProcessed uint64
	lastBaseMutex                 sync.Mutex
	lastBase, minBase             *big.Int
}

func NewGasPriceOracle(exp *Expanse) (self *GasPriceOracle) {
	self = &GasPriceOracle{}
	self.blocks = make(map[uint64]*blockPriceInfo)
	self.exp = exp
	self.chain = exp.chainManager
	self.events = exp.EventMux().Subscribe(
		core.ChainEvent{},
		core.ChainSplitEvent{},
	)

	minbase := new(big.Int).Mul(self.eth.GpoMinGasPrice, big.NewInt(100))
	minbase = minbase.Div(minbase, big.NewInt(int64(self.eth.GpobaseCorrectionFactor)))
	self.minBase = minbase

	self.processPastBlocks()
	go self.listenLoop()
	return
}

func (self *GasPriceOracle) processPastBlocks() {
	last := int64(-1)
	cblock := self.chain.CurrentBlock()
	if cblock != nil {
		last = int64(cblock.NumberU64())
	}
	first := int64(0)
	if last > gpoProcessPastBlocks {
		first = last - gpoProcessPastBlocks
	}
	self.firstProcessed = uint64(first)
	for i := first; i <= last; i++ {
		block := self.chain.GetBlockByNumber(uint64(i))
		if block != nil {
			self.processBlock(block)
		}
	}

}

func (self *GasPriceOracle) listenLoop() {
	for {
		ev, isopen := <-self.events.Chan()
		if !isopen {
			break
		}
		switch ev := ev.(type) {
		case core.ChainEvent:
			self.processBlock(ev.Block)
		case core.ChainSplitEvent:
			self.processBlock(ev.Block)
		}
	}
	self.events.Unsubscribe()
}

func (self *GasPriceOracle) processBlock(block *types.Block) {
	i := block.NumberU64()
	if i > self.lastProcessed {
		self.lastProcessed = i
	}

	lastBase := self.exp.GpoMinGasPrice
	bpl := self.blocks[i-1]
	if bpl != nil {
		lastBase = bpl.baseGasPrice
	}
	if lastBase == nil {
		return
	}

	var corr int
	lp := self.lowestPrice(block)
	if lp == nil {
		return
	}

	if lastBase.Cmp(lp) < 0 {
		corr = self.exp.GpobaseStepUp
	} else {
		corr = -self.exp.GpobaseStepDown
	}

	crand := int64(corr * (900 + rand.Intn(201)))
	newBase := new(big.Int).Mul(lastBase, big.NewInt(1000000+crand))
	newBase.Div(newBase, big.NewInt(1000000))

	if newBase.Cmp(self.minBase) < 0 {
		newBase = self.minBase
	}

	bpi := self.blocks[i]
	if bpi == nil {
		bpi = &blockPriceInfo{}
		self.blocks[i] = bpi
	}
	bpi.baseGasPrice = newBase
	self.lastBaseMutex.Lock()
	self.lastBase = newBase
	self.lastBaseMutex.Unlock()

	glog.V(logger.Detail).Infof("Processed block #%v, base price is %v\n", block.NumberU64(), newBase.Int64())
}

// returns the lowers possible price with which a tx was or could have been included
func (self *GasPriceOracle) lowestPrice(block *types.Block) *big.Int {
	gasUsed := big.NewInt(0)

	receipts := self.exp.BlockProcessor().GetBlockReceipts(block.Hash())
	if len(receipts) > 0 {
		if cgu := receipts[len(receipts)-1].CumulativeGasUsed; cgu != nil {
			gasUsed = receipts[len(receipts)-1].CumulativeGasUsed
		}
	}

	if new(big.Int).Mul(gasUsed, big.NewInt(100)).Cmp(new(big.Int).Mul(block.GasLimit(),
		big.NewInt(int64(self.exp.GpoFullBlockRatio)))) < 0 {
		// block is not full, could have posted a tx with MinGasPrice
		return big.NewInt(0)
	}

	txs := block.Transactions()
	if len(txs) == 0 {
		return big.NewInt(0)
	}
	// block is full, find smallest gasPrice
	minPrice := txs[0].GasPrice()
	for i := 1; i < len(txs); i++ {
		price := txs[i].GasPrice()
		if price.Cmp(minPrice) < 0 {
			minPrice = price
		}
	}
	return minPrice
}

func (self *GasPriceOracle) SuggestPrice() *big.Int {
	self.lastBaseMutex.Lock()
	base := self.lastBase
	self.lastBaseMutex.Unlock()

	if base == nil {
		base = self.exp.GpoMinGasPrice
	}
	if base == nil {
		return big.NewInt(10000000000000) // apparently MinGasPrice is not initialized during some tests
	}

	baseCorr := new(big.Int).Mul(base, big.NewInt(int64(self.exp.GpobaseCorrectionFactor)))
	baseCorr.Div(baseCorr, big.NewInt(100))

	if baseCorr.Cmp(self.exp.GpoMinGasPrice) < 0 {
		return self.exp.GpoMinGasPrice
	}

	if baseCorr.Cmp(self.exp.GpoMaxGasPrice) > 0 {
		return self.exp.GpoMaxGasPrice
	}

	return baseCorr
}
