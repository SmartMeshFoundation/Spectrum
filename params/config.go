// Copyright 2016 The Spectrum Authors
// This file is part of the Spectrum library.
//
// The Spectrum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The Spectrum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the Spectrum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"fmt"
	"math/big"

	"github.com/SmartMeshFoundation/Spectrum/common"
)

var (
	MainnetGenesisHash = common.HexToHash("0xa8ab5ecb308dd061e6baa8c8a04a62b3e35fae7aae1393921f8d52deac2c5712") // enforce below configs on
	TestnetGenesisHash = common.HexToHash("0x28242478c4f01d9208a79d962b7f6383af488565f36aea7038a34020c07db697") // Testnet genesis hash to enforce below configs on
	DevnetGenesisHash  = common.HexToHash("0xe11f21ee330cfb50c3f31d9b792df2fb5e196739d562e642416974f339aa4304")
)

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(20180430),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    nil,
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		// add by liangc : change default consensus for dev
		// if skip this vsn please set 0 or nil to block and set common.Address{} to address
		// 0.0.5 : ready for release
		Chief005Block:   big.NewInt(2),
		Chief005Address: common.HexToAddress("0xf0d2ad4e0d25cfea98e9640329993bbc52396abd"),
		// 0.0.6 : ready for release
		Chief006Block:   big.NewInt(595888),
		Chief006Address: common.HexToAddress("0xba7f507d5aab3e931312512c234fbeb85cbd9dce"),

		Meshbox001Block:   big.NewInt(1870333),
		Meshbox001Address: common.HexToAddress("0xf0ced0b1ce8738eeac06fdca51e0ff398328634b"),
		// Note that : meshbox002 must is superset of meshbox001

		ChiefBaseAddress: common.HexToAddress("0x"),

		Anmap001Block:   big.NewInt(0),
		Anmap001Address: common.HexToAddress("0x"),
		MinMinerBalance: new(big.Int).Mul(big.NewInt(100000), big.NewInt(Ether)),

		SIP001Block: big.NewInt(0), // new rules for chief.tx of gaspool
		SIP002Block: big.NewInt(588888),
		SIP003Block: big.NewInt(808888),
		SIP004Block: big.NewInt(0),
		//SIP004Block: big.NewInt(2888888), // maybe begin at 2019-8-18
		Tribe: &TribeConfig{},
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(3),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: true,
		EIP150Block:    nil,
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(62678),
		//ByzantiumBlock: big.NewInt(1700000),
		//Clique: &CliqueConfig{ Period: 15, Epoch:  30000}, //Ethash: new(EthashConfig),
		// if skip this vsn please set 0 or nil to block and set common.Address{} to address
		// 0.0.2
		Chief002Block:   big.NewInt(2),
		Chief002Address: common.HexToAddress("0x9ec55c1dafd4a487e41da33e344aef86da41ab82"),
		// 0.0.3
		Chief003Block:   big.NewInt(113772), //hard fork testnet
		Chief003Address: common.HexToAddress("0xac28e532b3fac82554fc7b0b8b62549deeeb33a9"),
		// 0.0.4 : fix bug redeploy
		Chief004Block:   big.NewInt(120305), //hard fork testnet : fix bugs and debug chief
		Chief004Address: common.HexToAddress("0xe242e2bcf5b0da6518320210fab0a27458bc0674"),
		// 0.0.5 : ready for release
		Chief005Block:   big.NewInt(176244),
		Chief005Address: common.HexToAddress("0xe90da8175922925dfb40e6505b508f1042e807aa"),
		// 0.0.6 : ready for release
		Chief006Block:   big.NewInt(595710),
		Chief006Address: common.HexToAddress("0x53cb83888e6d28cf7ec168308c65172001f441aa"),

		Meshbox001Block:   big.NewInt(1976666),
		Meshbox001Address: common.HexToAddress("0x0f15e1e44322b2946215705d2ed60cba899f0b38"),
		// Note that : meshbox002 must is superset of meshbox001

		Anmap001Block:   big.NewInt(2200050),
		Anmap001Address: common.HexToAddress("0xffab698a4cead35a6f52569e328f9127e66413bb"),

		MinMinerBalance: new(big.Int).Mul(big.NewInt(10), big.NewInt(Ether)),

		// new rules for chief.tx of gaspool
		SIP001Block: big.NewInt(176222),
		SIP002Block: big.NewInt(525233),
		SIP003Block: big.NewInt(917013),
		SIP004Block: big.NewInt(2212557),

		// base block link to Chief100Block
		ChiefBaseAddress: common.HexToAddress("0x4651817f6595e9908de3dcf9d8b585796e209c32"),

		// PocBlock must less than Chief100Block
		PocAddress: common.HexToAddress("0xfb6da09a22c0b2bd09a2b24274ceac642c73be14"),

		Chief100Block:   big.NewInt(2591595),
		Chief100Address: common.HexToAddress("0x3f0faed157d6d6dfe191fb71500b63382f0354dd"),

		Tribe: &TribeConfig{},
	}

	// DevnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	DevnetChainConfig = &ChainConfig{
		ChainId:        big.NewInt(4),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    nil,
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),

		// if skip this vsn please set 0 or nil to block and set common.Address{} to address
		// 0.0.6 : ready for release
		// before block 2 gas used : 2631756
		// after gas used : 2892524
		Chief007Block:   big.NewInt(3),
		Chief007Address: common.HexToAddress("0x57d2bcd8d702999daf240793919675c90b12a17a"),

		// ############
		// # DEBUG >>>>
		// ############

		// Note that : meshbox002 must is superset of meshbox001
		Meshbox002Block:   big.NewInt(6),
		Meshbox002Address: common.HexToAddress("0x7880adce4504fd39645aabb3efb53824d9b0c21b"),

		Anmap001Block:   big.NewInt(6),
		Anmap001Address: common.HexToAddress("0x143084accd6472ad502b59c3197f9ed5f797b966"),

		// base block link to Chief100Block
		//ChiefBaseAddress: common.HexToAddress("0xad61f1201f592fbf13d2645f9c59d8d5f82a1837"), //liang
		ChiefBaseAddress: common.HexToAddress("0xe551a89706753b64eaa74198b1d0d19cbc97489a"),

		// PocBlock must less than Chief100Block
		//PocBlock: big.NewInt(35),
		//PocAddress: common.HexToAddress("0x901c0636c4fc83f353bca2db85e2ace886a9416d"), //liang
		PocAddress: common.HexToAddress("0xe62e6655dd0737794335a1b8f1f51204919d2ced"),

		Chief100Block: big.NewInt(50),
		//Chief100Address: common.HexToAddress("0x"),
		//Chief100Address: common.HexToAddress("0x6d05f6aa4e19e20cd781fa3def97bbfd0b980534"), // liang
		Chief100Address: common.HexToAddress("0x3f81857b5b2637c954b8a5a7c467d22814c4ec53"),

		/*
			Meshbox002Block:   big.NewInt(6),
			Meshbox002Address: common.HexToAddress("0xc4a2d182fe92f0eadffbddea9a0977d5b95b31a5"),

			Anmap001Block:   big.NewInt(6),
			Anmap001Address: common.HexToAddress("0x57d2bcd8d702999daf240793919675c90b12a17a"),

			// base block link to Chief100Block
			ChiefBaseAddress: common.HexToAddress("0x7880adce4504fd39645aabb3efb53824d9b0c21b"),

			// PocBlock must less than Chief100Block
			PocBlock:   big.NewInt(20),
			PocAddress: common.HexToAddress("0xad61f1201f592fbf13d2645f9c59d8d5f82a1837"),

			Chief100Block:   big.NewInt(22),
			Chief100Address: common.HexToAddress("0x0f91c3f2e10a0b53d6b3b4d6c7b41ab77c7d0674"),

		*/
		// ############
		// # DEBUG <<<<
		// ############

		MinMinerBalance: new(big.Int).Mul(big.NewInt(1), big.NewInt(Ether)),

		// new rules for chief.tx of gaspool
		SIP001Block: big.NewInt(0),
		SIP002Block: big.NewInt(1),
		SIP003Block: big.NewInt(3),
		SIP004Block: big.NewInt(3),

		Tribe: &TribeConfig{Period: 3},
	}

	// AllEthashProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Ethash consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllEthashProtocolChanges = &ChainConfig{
		ChainId:        big.NewInt(1337),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		Clique:         nil,
		Tribe:          nil,
	}

	// AllCliqueProtocolChanges contains every protocol change (EIPs) introduced
	// and accepted by the Ethereum core developers into the Clique consensus.
	//
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllCliqueProtocolChanges = &ChainConfig{
		ChainId:        big.NewInt(1337),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         nil,
		Clique:         &CliqueConfig{Period: 0, Epoch: 30000},
		Tribe:          nil,
	}
	TestChainConfig = &ChainConfig{
		ChainId:        big.NewInt(1),
		HomesteadBlock: big.NewInt(0),
		DAOForkBlock:   nil,
		DAOForkSupport: false,
		EIP150Block:    big.NewInt(0),
		EIP150Hash:     common.Hash{},
		EIP155Block:    big.NewInt(0),
		EIP158Block:    big.NewInt(0),
		ByzantiumBlock: big.NewInt(0),
		Ethash:         new(EthashConfig),
		Clique:         nil,
		Tribe:          nil,
	}
	TestRules = TestChainConfig.Rules(new(big.Int))
)

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainId *big.Int `json:"chainId"` // Chain id identifies the current chain and is used for replay protection

	HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)

	DAOForkBlock   *big.Int `json:"daoForkBlock,omitempty"`   // TheDAO hard-fork switch block (nil = no fork)
	DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int    `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block

	ByzantiumBlock *big.Int `json:"byzantiumBlock,omitempty"` // Byzantium switch block (nil = no fork, 0 = already on byzantium)

	// Various consensus engines
	Ethash *EthashConfig `json:"ethash,omitempty"`
	Clique *CliqueConfig `json:"clique,omitempty"`
	Tribe  *TribeConfig  `json:"tribe,omitempty"` // add by liangc

	// >>> add by liangc : set chief start number >>>
	// chief.sol vsn 0.0.2
	Chief002Block   *big.Int       `json:"chief002Block,omitempty"`
	Chief002Address common.Address `json:"chief002Address,omitempty"`
	// chief.sol vsn 0.0.3
	Chief003Block   *big.Int       `json:"chief003Block,omitempty"`
	Chief003Address common.Address `json:"chief003Address,omitempty"`
	// chief.sol vsn 0.0.4
	Chief004Block   *big.Int       `json:"chief004Block,omitempty"`
	Chief004Address common.Address `json:"chief004Address,omitempty"`
	// chief.sol vsn 0.0.5
	Chief005Block   *big.Int       `json:"chief005Block,omitempty"`
	Chief005Address common.Address `json:"chief005Address,omitempty"`
	// chief.sol vsn 0.0.6
	Chief006Block   *big.Int       `json:"chief006Block,omitempty"`
	Chief006Address common.Address `json:"chief006Address,omitempty"`

	// extends 0.0.6 only for dev and test
	Chief007Block   *big.Int       `json:"chief007Block,omitempty"`
	Chief007Address common.Address `json:"chief007Address,omitempty"`
	// https://github.com/SmartMeshFoundation/Spectrum/wiki/%5BChinese%5D-v1.0.0-Standard
	Chief100Block   *big.Int       `json:"chief100Block,omitempty"`
	Chief100Address common.Address `json:"chief100Address,omitempty"`

	ChiefBaseAddress common.Address `json:"chiefBaseAddress,omitempty"`

	//PocBlock   *big.Int       `json:"PocBlock,omitempty"`
	PocAddress common.Address `json:"PocAddress,omitempty"`

	// <<< add by liangc : set chief start number <<<
	// add by liangc : new rules for chief.tx of gaspool
	SIP001Block *big.Int `json:"sip001Block,omitempty"` // SIP001 HF block
	// add by liangc : new rules for block period
	SIP002Block *big.Int `json:"sip002Block,omitempty"` // SIP002 HF block
	// add by liangc : 18-09-13 : incompatible HomesteadSigner begin at this number
	SIP003Block *big.Int `json:"sip003Block,omitempty"`
	// add by liangc : 19-03-27 : for smc-0.6.0
	// https://github.com/SmartMeshFoundation/Spectrum/wiki/%5BChinese%5D-v0.6.0-Standard
	SIP004Block *big.Int `json:"sip004Block,omitempty"`

	Meshbox001Block   *big.Int       `json:"meshbox001Block,omitempty"`
	Meshbox001Address common.Address `json:"meshbox001Address,omitempty"`
	// Note that : meshbox002 must is superset of meshbox001
	Meshbox002Block   *big.Int       `json:"meshbox002Block,omitempty"`
	Meshbox002Address common.Address `json:"meshbox002Address,omitempty"`

	Anmap001Block   *big.Int       `json:"anmap001Block,omitempty"`
	Anmap001Address common.Address `json:"anmap001Address,omitempty"`

	MinMinerBalance *big.Int `json:"minMinerBalance,omitempty"`
}

// EthashConfig is the consensus engine configs for proof-of-work based sealing.
type EthashConfig struct{}

// String implements the stringer interface, returning the consensus engine details.
func (c *EthashConfig) String() string {
	return "ethash"
}

// CliqueConfig is the consensus engine configs for proof-of-authority based sealing.
type CliqueConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
	Epoch  uint64 `json:"epoch"`  // Epoch length to reset votes and checkpoint
}

// String implements the stringer interface, returning the consensus engine details.
func (c *CliqueConfig) String() string {
	return "clique"
}

// TribeConfig is the consensus engine configs.
type TribeConfig struct {
	Period uint64 `json:"period"` // Number of seconds between blocks to enforce
}

// String implements the stringer interface, returning the consensus engine details.
func (c *TribeConfig) String() string {
	return "tribe"
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	case c.Tribe != nil:
		engine = c.Tribe
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Homestead: %v DAO: %v DAOSupport: %v EIP150: %v EIP155: %v EIP158: %v Byzantium: %v Engine: %v}",
		c.ChainId,
		c.HomesteadBlock,
		c.DAOForkBlock,
		c.DAOForkSupport,
		c.EIP150Block,
		c.EIP155Block,
		c.EIP158Block,
		c.ByzantiumBlock,
		engine,
	)
}

// IsHomestead returns whether num is either equal to the homestead block or greater.
func (c *ChainConfig) IsHomestead(num *big.Int) bool {
	return isForked(c.HomesteadBlock, num)
}

// IsDAO returns whether num is either equal to the DAO fork block or greater.
func (c *ChainConfig) IsDAOFork(num *big.Int) bool {
	return isForked(c.DAOForkBlock, num)
}

func (c *ChainConfig) IsEIP150(num *big.Int) bool {
	return isForked(c.EIP150Block, num)
}

func (c *ChainConfig) IsEIP155(num *big.Int) bool {
	return isForked(c.EIP155Block, num)
}

func (c *ChainConfig) IsEIP158(num *big.Int) bool {
	return isForked(c.EIP158Block, num)
}

func (c *ChainConfig) IsByzantium(num *big.Int) bool {
	// add by liangc : set default byzantium
	return isForked(c.ByzantiumBlock, num)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	if num == nil {
		return GasTableHomestead
	}
	switch {
	case c.IsEIP158(num):
		return GasTableEIP158
	case c.IsEIP150(num):
		return GasTableEIP150
	default:
		return GasTableHomestead
	}
}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	if isForkIncompatible(c.HomesteadBlock, newcfg.HomesteadBlock, head) {
		return newCompatError("Homestead fork block", c.HomesteadBlock, newcfg.HomesteadBlock)
	}
	if isForkIncompatible(c.DAOForkBlock, newcfg.DAOForkBlock, head) {
		return newCompatError("DAO fork block", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if c.IsDAOFork(head) && c.DAOForkSupport != newcfg.DAOForkSupport {
		return newCompatError("DAO fork support flag", c.DAOForkBlock, newcfg.DAOForkBlock)
	}
	if isForkIncompatible(c.EIP150Block, newcfg.EIP150Block, head) {
		return newCompatError("EIP150 fork block", c.EIP150Block, newcfg.EIP150Block)
	}
	if isForkIncompatible(c.EIP155Block, newcfg.EIP155Block, head) {
		return newCompatError("EIP155 fork block", c.EIP155Block, newcfg.EIP155Block)
	}
	if isForkIncompatible(c.EIP158Block, newcfg.EIP158Block, head) {
		return newCompatError("EIP158 fork block", c.EIP158Block, newcfg.EIP158Block)
	}
	if c.IsEIP158(head) && !configNumEqual(c.ChainId, newcfg.ChainId) {
		return newCompatError("EIP158 chain ID", c.EIP158Block, newcfg.EIP158Block)
	}
	if isForkIncompatible(c.ByzantiumBlock, newcfg.ByzantiumBlock, head) {
		return newCompatError("Byzantium fork block", c.ByzantiumBlock, newcfg.ByzantiumBlock)
	}
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainId                                   *big.Int
	IsHomestead, IsEIP150, IsEIP155, IsEIP158 bool
	IsByzantium                               bool
}

func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainId := c.ChainId
	if chainId == nil {
		chainId = new(big.Int)
	}
	return Rules{ChainId: new(big.Int).Set(chainId), IsHomestead: c.IsHomestead(num), IsEIP150: c.IsEIP150(num), IsEIP155: c.IsEIP155(num), IsEIP158: c.IsEIP158(num), IsByzantium: c.IsByzantium(num)}
}
