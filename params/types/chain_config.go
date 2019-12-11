package paramtypes

import (
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	common2 "github.com/ethereum/go-ethereum/params/types/common"
	"github.com/ethereum/go-ethereum/params/types/goethereum"
)

// MultiGethChainConfig is the core config which determines the blockchain settings.
//
// MultiGethChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type MultiGethChainConfig struct {

	// Embedded ethereum/go-ethereum MultiGethChainConfig.
	// This is not a pointer value because it can be expected that there will
	// always be at least one (even zero-value) value desired from that data type, eg ChainID or engine.
	//goethereum.MultiGethChainConfig

	// Following fields are left commented because it's useful to see pairings,
	// both for reference and edification.

	NetworkID uint64   `json:"networkId"`
	ChainID   *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	// HF: Homestead
	//HomesteadBlock *big.Int `json:"homesteadBlock,omitempty"` // Homestead switch block (nil = no fork, 0 = already homestead)
	// "Homestead Hard-fork Changes"
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-2.md
	EIP2FBlock *big.Int `json:"eip2FBlock,omitempty"`
	// DELEGATECALL
	// https://eips.ethereum.org/EIPS/eip-7
	EIP7FBlock *big.Int `json:"eip7FBlock,omitempy"`
	// Note: EIP 8 was also included in this fork, but was not backwards-incompatible

	// HF: DAO
	DAOForkBlock *big.Int `json:"daoForkBlock,omitempty"` // TheDAO hard-fork switch block (nil = no fork)
	//DAOForkSupport bool     `json:"daoForkSupport,omitempty"` // Whether the nodes supports or opposes the DAO hard-fork

	// HF: Tangerine Whistle
	// EIP150 implements the Gas price changes (https://github.com/ethereum/EIPs/issues/150)
	EIP150Block *big.Int `json:"eip150Block,omitempty"` // EIP150 HF block (nil = no fork)
	//EIP150Hash  common.Hash `json:"eip150Hash,omitempty"`  // EIP150 HF hash (needed for header only clients as only gas pricing changed)

	// HF: Spurious Dragon
	EIP155Block *big.Int `json:"eip155Block,omitempty"` // EIP155 HF block
	//EIP158Block *big.Int `json:"eip158Block,omitempty"` // EIP158 HF block, includes implementations of 158/161, 160, and 170
	//
	// EXP cost increase
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-160.md
	// NOTE: this json tag:
	// (a.) varies from it's 'siblings', which have 'F's in them
	// (b.) without the 'F' will vary from ETH implementations if they choose to accept the proposed changes
	// with corresponding refactoring (https://github.com/ethereum/go-ethereum/pull/18401)
	EIP160FBlock *big.Int `json:"eip160Block,omitempty"`
	// State trie clearing (== EIP158 proper)
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-161.md
	EIP161FBlock *big.Int `json:"eip161FBlock,omitempty"`
	// Contract code size limit
	// https://github.com/ethereum/EIPs/blob/master/EIPS/eip-170.md
	EIP170FBlock *big.Int `json:"eip170FBlock,omitempty"`

	// HF: Byzantium
	//ByzantiumBlock *big.Int `json:"byzantiumBlock,omitempty"` // Byzantium switch block (nil = no fork, 0 = already on byzantium)

	// Difficulty adjustment to target mean block time including uncles
	// https://github.com/ethereum/EIPs/issues/100
	EIP100FBlock *big.Int `json:"eip100FBlock,omitempty"`
	// Opcode REVERT
	// https://eips.ethereum.org/EIPS/eip-140
	EIP140FBlock *big.Int `json:"eip140FBlock,omitempty"`
	// Precompiled contract for bigint_modexp
	// https://github.com/ethereum/EIPs/issues/198
	EIP198FBlock *big.Int `json:"eip198FBlock,omitempty"`
	// Opcodes RETURNDATACOPY, RETURNDATASIZE
	// https://github.com/ethereum/EIPs/issues/211
	EIP211FBlock *big.Int `json:"eip211FBlock,omitempty"`
	// Precompiled contract for pairing check
	// https://github.com/ethereum/EIPs/issues/212
	EIP212FBlock *big.Int `json:"eip212FBlock,omitempty"`
	// Precompiled contracts for addition and scalar multiplication on the elliptic curve alt_bn128
	// https://github.com/ethereum/EIPs/issues/213
	EIP213FBlock *big.Int `json:"eip213FBlock,omitempty"`
	// Opcode STATICCALL
	// https://github.com/ethereum/EIPs/issues/214
	EIP214FBlock *big.Int `json:"eip214FBlock,omitempty"`
	// Metropolis diff bomb delay and reducing block reward
	// https://github.com/ethereum/EIPs/issues/649
	// note that this is closely related to EIP100.
	// In fact, EIP100 is bundled in
	eip649FInferred bool     `json:"-"`
	EIP649FBlock    *big.Int `json:"-"`
	// Transaction receipt status
	// https://github.com/ethereum/EIPs/issues/658
	EIP658FBlock *big.Int `json:"eip658FBlock,omitempty"`
	// NOT CONFIGURABLE: prevent overwriting contracts
	// https://github.com/ethereum/EIPs/issues/684
	// EIP684FBlock *big.Int `json:"eip684BFlock,omitempty"`

	// HF: Constantinople
	//ConstantinopleBlock *big.Int `json:"constantinopleBlock,omitempty"` // Constantinople switch block (nil = no fork, 0 = already activated)
	//
	// Opcodes SHR, SHL, SAR
	// https://eips.ethereum.org/EIPS/eip-145
	EIP145FBlock *big.Int `json:"eip145FBlock,omitempty"`
	// Opcode CREATE2
	// https://eips.ethereum.org/EIPS/eip-1014
	EIP1014FBlock *big.Int `json:"eip1014FBlock,omitempty"`
	// Opcode EXTCODEHASH
	// https://eips.ethereum.org/EIPS/eip-1052
	EIP1052FBlock *big.Int `json:"eip1052FBlock,omitempty"`
	// Constantinople difficulty bomb delay and block reward adjustment
	// https://eips.ethereum.org/EIPS/eip-1234
	eip1234FInferred bool     `json:"-"`
	EIP1234FBlock    *big.Int `json:"-"`
	// Net gas metering
	// https://eips.ethereum.org/EIPS/eip-1283
	EIP1283FBlock *big.Int `json:"eip1283FBlock,omitempty"`

	PetersburgBlock *big.Int `json:"petersburgBlock,omitempty"` // Petersburg switch block (nil = same as Constantinople)

	// HF: Istanbul
	//IstanbulBlock *big.Int `json:"istanbulBlock,omitempty"` // Istanbul switch block (nil = no fork, 0 = already on istanbul)
	//
	// EIP-152: Add Blake2 compression function F precompile
	EIP152FBlock *big.Int `json:"eip152FBlock,omitempty"`
	// EIP-1108: Reduce alt_bn128 precompile gas costs
	EIP1108FBlock *big.Int `json:"eip1108FBlock,omitempty"`
	// EIP-1344: Add ChainID opcode
	EIP1344FBlock *big.Int `json:"eip1344FBlock,omitempty"`
	// EIP-1884: Repricing for trie-size-dependent opcodes
	EIP1884FBlock *big.Int `json:"eip1884FBlock,omitempty"`
	// EIP-2028: Calldata gas cost reduction
	EIP2028FBlock *big.Int `json:"eip2028FBlock,omitempty"`
	// EIP-2200: Rebalance net-metered SSTORE gas cost with consideration of SLOAD gas cost change
	EIP2200FBlock *big.Int `json:"eip2200FBlock,omitempty"`

	//EWASMBlock *big.Int `json:"ewasmBlock,omitempty"` // EWASM switch block (nil = no fork, 0 = already activated)

	ECIP1010PauseBlock *big.Int `json:"ecip1010PauseBlock,omitempty"` // ECIP1010 pause HF block
	ECIP1010Length     *big.Int `json:"ecip1010Length,omitempty"`     // ECIP1010 length
	ECIP1017FBlock     *big.Int `json:"ecip1017FBlock,omitempty"`
	ECIP1017EraRounds  *big.Int `json:"ecip1017EraRounds,omitempty"` // ECIP1017 era rounds
	DisposalBlock      *big.Int `json:"disposalBlock,omitempty"`     // Bomb disposal HF block
	SocialBlock        *big.Int `json:"socialBlock,omitempty"`       // Ethereum Social Reward block
	EthersocialBlock   *big.Int `json:"ethersocialBlock,omitempty"`  // Ethersocial Reward block
	
	// Various consensus engines
	Ethash *goethereum.EthashConfig `json:"ethash,omitempty"`
	Clique *goethereum.CliqueConfig `json:"clique,omitempty"`

	TrustedCheckpoint       *goethereum.TrustedCheckpoint      `json:"trustedCheckpoint"`
	TrustedCheckpointOracle *goethereum.CheckpointOracleConfig `json:"trustedCheckpointOracle"`

	DifficultyBombDelaySchedule common2.Uint64BigMapEncodesHex `json:"difficultyBombDelays,omitempty"'` // JSON tag matches Parity's
	BlockRewardSchedule         common2.Uint64BigMapEncodesHex `json:"blockReward,omitempty"`           // JSON tag matches Parity's

	RequireBlockHashes map[uint64]common.Hash `json:"requireBlockHashes"`
}

// String implements the fmt.Stringer interface.
func (c *MultiGethChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Ethash != nil:
		engine = c.Ethash
	case c.Clique != nil:
		engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("<FIXME!!> NetworkID: %v, ChainID: %v Engine: %v",
		c.NetworkID,
		c.ChainID,
		engine,
	)
}
