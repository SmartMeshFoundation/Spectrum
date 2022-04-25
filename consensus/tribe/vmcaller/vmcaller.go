package vmcaller

import (
	"github.com/MeshBoxTech/mesh-chain/common"
	"github.com/MeshBoxTech/mesh-chain/core"
	"github.com/MeshBoxTech/mesh-chain/core/state"
	"github.com/MeshBoxTech/mesh-chain/core/types"
	"github.com/MeshBoxTech/mesh-chain/core/vm"
	"github.com/MeshBoxTech/mesh-chain/log"
	"github.com/MeshBoxTech/mesh-chain/params"
	"math/big"
)

//// ExecuteMsg executes transaction sent to system contracts.
func ExecuteMsg(msg core.Message, state *state.StateDB, header *types.Header, chainContext core.ChainContext, chainConfig *params.ChainConfig) (ret []byte, err error) {
	evmContext := core.NewEVMContext(msg, header, chainContext, nil)
	vmenv := vm.NewEVM(evmContext, state, chainConfig, vm.Config{})

	ret, _, err = vmenv.Call(vm.AccountRef(msg.From()), *msg.To(), msg.Data(), msg.Gas().Uint64(), msg.Value())
	// Finalise the statedb so any changes can take effect,
	// and especially if the `from` account is empty, it can be finally deleted.
	state.Finalise(true)
	if err != nil {
		log.Error("ExecuteMsg failed", "err", err, "ret", string(ret))
	}
	return ret, err
}

// NewLegacyMessage builds a message for consensus and system governance actions, it will not consumes any fee.
func NewLegacyMessage(from common.Address, to *common.Address, nonce uint64, amount *big.Int, gasLimit *big.Int, gasPrice *big.Int, data []byte, checkNonce bool) types.Message {
	return types.NewMessage(from, to, nonce, amount, gasLimit, gasPrice, data, checkNonce)
}
