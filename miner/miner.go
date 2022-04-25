// Copyright 2014 The mesh-chain Authors
// This file is part of the mesh-chain library.
//
// The mesh-chain library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The mesh-chain library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the mesh-chain library. If not, see <http://www.gnu.org/licenses/>.

// Package miner implements Ethereum block creation and mining.
package miner

import (
	"fmt"
	"sync"
	"sync/atomic"

	"github.com/MeshBoxTech/mesh-chain/accounts"
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/consensus"
	"github.com/MeshBoxTech/mesh-chain/consensus/tribe"
	"github.com/MeshBoxTech/mesh-chain/core"
	"github.com/MeshBoxTech/mesh-chain/core/state"
	"github.com/MeshBoxTech/mesh-chain/core/types"
	"github.com/MeshBoxTech/mesh-chain/eth/downloader"
	"github.com/MeshBoxTech/mesh-chain/ethdb"
	"github.com/MeshBoxTech/mesh-chain/event"
	"github.com/MeshBoxTech/mesh-chain/log"
	"github.com/MeshBoxTech/mesh-chain/params"
)

// Backend wraps all methods required for mining.
type Backend interface {
	AccountManager() *accounts.Manager
	BlockChain() *core.BlockChain
	TxPool() *core.TxPool
	ChainDb() ethdb.Database
}

// Miner creates blocks and searches for proof-of-work values.
type Miner struct {
	stop *sync.Map
	mux  *event.TypeMux

	worker *worker

	coinbase common.Address
	mining   int32
	eth      Backend
	engine   consensus.Engine

	canStart    int32 // can start indicates whether we can start the mining operation
	shouldStart int32 // should start indicates whether we should start after sync
}

func New(eth Backend, config *params.ChainConfig, mux *event.TypeMux, engine consensus.Engine) *Miner {
	miner := &Miner{
		eth:      eth,
		mux:      mux,
		engine:   engine,
		worker:   newWorker(config, engine, common.Address{}, eth, mux),
		canStart: 1,
		stop:     new(sync.Map),
	}
	miner.Register(NewCpuAgent(eth.BlockChain(), engine))
	go miner.update()
	if tribe, ok := miner.engine.(*tribe.Tribe); ok {
		close(params.TribeReadyForAcceptTxs)
		go func() {
			log.Info("miner wait miner address")
			rtn := make(chan common.Address)
			tribe.GetMinerAddressByChan(rtn)
			tma := <-rtn
			log.Info("miner get miner address")
			go miner.Start(tma)
			log.Info("👷 Tribe and miner is started .")
		}()
	}
	return miner
}

// update keeps track of the downloader events. Please be aware that this is a one shot type of update loop.
// It's entered once and as soon as `Done` or `Failed` has been broadcasted the events are unregistered and
// the loop is exited. This to prevent a major security vuln where external parties can DOS you with blocks
// and halt your mining operation for as long as the DOS continues.
func (self *Miner) update() {
	events := self.mux.Subscribe(downloader.StartEvent{}, downloader.DoneEvent{}, downloader.FailedEvent{})
out:
	for ev := range events.Chan() {
		switch ev.Data.(type) {
		case downloader.StartEvent:
			atomic.StoreInt32(&self.canStart, 0)
			if self.Mining() {
				self.Stop()
				atomic.StoreInt32(&self.shouldStart, 1)
				log.Info("Mining aborted due to sync")
			}
		case downloader.DoneEvent, downloader.FailedEvent:
			shouldStart := atomic.LoadInt32(&self.shouldStart) == 1

			atomic.StoreInt32(&self.canStart, 1)
			atomic.StoreInt32(&self.shouldStart, 0)
			if shouldStart {
				log.Info("miner start after sync complete")
				go self.Start(self.coinbase)
			}
			// unsubscribe. we're only interested in this event once
			events.Unsubscribe()
			// stop immediately and ignore all further pending events
			break out
		}
	}
}
// async run
var xx int32 = 0

func (self *Miner) Start(coinbase common.Address) {
	stop := make(chan struct{})
	self.stop.Store(stop, struct{}{})
	atomic.AddInt32(&xx, 1)
	atomic.StoreInt32(&self.shouldStart, 1)

	self.worker.setEtherbase(coinbase)
	self.coinbase = coinbase
	if atomic.LoadInt32(&self.canStart) == 0 {
		log.Info("Network syncing, will start miner afterwards")
		return
	}
	atomic.StoreInt32(&self.mining, 1)
	log.Info("Starting mining operation")

	//if tribe, ok := self.engine.(*tribe.Tribe); ok && self.eth.BlockChain().CurrentBlock().NumberU64() > 3 {
	//	i := 0
	//	for {
	//		log.Info("<<MinerStart>> loop_start", "i", i, "num", self.eth.BlockChain().CurrentBlock().Number())
	//		m := tribe.Status.GetMinerAddress()
	//		s, err := self.worker.chain.State()
	//		if err != nil {
	//			log.Error("miner start fail", err)
	//		}
	//		cn := self.eth.BlockChain().CurrentBlock().Number()
	//		// SIP100 skip this verfiy
	//		if params.IsSIP100Block(cn) {
	//			break
	//		}
	//		if s.GetBalance(m).Cmp(params.ChiefBaseBalance) >= 0 {
	//			break
	//		}
	//		if atomic.LoadInt32(&self.mining) == 0 {
	//			log.Error("miner>>>>>>>>", err)
	//			return
	//		}
	//		select {
	//		case <-stop:
	//			return
	//		default:
	//			log.Info("default>>>>>>>>")
	//			<-time.After(time.Second * 7)
	//		}
	//		i++
	//	}
	//}
	// may be pending at 'tribe.WaitingNomination' in 'worker.start' so change to async
	go func() {
		s := make(chan int)
		self.worker.start(s)
		select {
		case <-stop:
			return
		case <-s:
			self.worker.commitNewWork()
		}
	}()

}

func (self *Miner) dostop() {
	self.stop.Range(func(k, v interface{}) bool {
		defer func() { recover() }()
		defer close(k.(chan struct{}))
		self.stop.Delete(k)
		return true
	})
}

func (self *Miner) Stop() {
	self.dostop()
	self.worker.stop()
	atomic.StoreInt32(&self.mining, 0)
	atomic.StoreInt32(&self.shouldStart, 0)
}

func (self *Miner) Register(agent Agent) {
	if self.Mining() {
		agent.Start()
	}
	self.worker.register(agent)
}

func (self *Miner) Unregister(agent Agent) {
	self.worker.unregister(agent)
}

func (self *Miner) Mining() bool {
	return atomic.LoadInt32(&self.mining) > 0
}

func (self *Miner) HashRate() (tot int64) {
	if pow, ok := self.engine.(consensus.PoW); ok {
		tot += int64(pow.Hashrate())
	}
	// do we care this might race? is it worth we're rewriting some
	// aspects of the worker/locking up agents so we can get an accurate
	// hashrate?
	for agent := range self.worker.agents {
		if _, ok := agent.(*CpuAgent); !ok {
			tot += agent.GetHashRate()
		}
	}
	return
}

func (self *Miner) SetExtra(extra []byte) error {
	if uint64(len(extra)) > params.MaximumExtraDataSize {
		return fmt.Errorf("Extra exceeds max length. %d > %v", len(extra), params.MaximumExtraDataSize)
	}
	self.worker.setExtra(extra)
	return nil
}

// Pending returns the currently pending block and associated state.
func (self *Miner) Pending() (*types.Block, *state.StateDB) {
	return self.worker.pending()
}

// PendingBlock returns the currently pending block.
//
// Note, to access both the pending block and the pending state
// simultaneously, please use Pending(), as the pending state can
// change between multiple method calls
func (self *Miner) PendingBlock() *types.Block {
	return self.worker.pendingBlock()
}

func (self *Miner) SetEtherbase(addr common.Address) {
	self.coinbase = addr
	self.worker.setEtherbase(addr)
}
