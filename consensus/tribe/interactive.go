package tribe

import (
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/consensus"
	"github.com/MeshBoxTech/mesh-chain/core/types"
)

type chainContext struct {
	chainReader consensus.ChainReader
	engine      consensus.Engine
}

func newChainContext(chainReader consensus.ChainReader, engine consensus.Engine) *chainContext {
	return &chainContext{
		chainReader: chainReader,
		engine:      engine,
	}
}

// Engine retrieves the chain's consensus engine.
func (cc *chainContext) Engine() consensus.Engine {
	return cc.engine
}

// GetHeader returns the hash corresponding to their hash.
func (cc *chainContext) GetHeader(hash common.Hash, number uint64) *types.Header {
	return cc.chainReader.GetHeader(hash, number)
}
