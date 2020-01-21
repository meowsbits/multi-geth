package eth

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params/types/ctypes"
)

type PrivateChainConfigAPI struct {
	e *Ethereum
}

func NewPrivateChainConfigAPI(e *Ethereum) *PrivateChainConfigAPI {
	return &PrivateChainConfigAPI{e}
}

func (api *PrivateChainConfigAPI) GetAccountStartNonce() *uint64 {
	return api.e.blockchain.Config().GetAccountStartNonce()
}

func (api *PrivateChainConfigAPI) SetAccountStartNonce(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetMaximumExtraDataSize() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetMaximumExtraDataSize(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetMinGasLimit() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetMinGasLimit(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetGasLimitBoundDivisor() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetGasLimitBoundDivisor(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetNetworkID() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetNetworkID(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetChainID() *big.Int {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetChainID(i *big.Int) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetMaxCodeSize() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetMaxCodeSize(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP7Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP7Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP150Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP150Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP152Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP152Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP160Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP160Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP161abcTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP161abcTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP161dTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP161dTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP170Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP170Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP155Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP155Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP140Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP140Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP198Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP198Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP211Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP211Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP212Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP212Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP213Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP213Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP214Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP214Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP658Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP658Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP145Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP145Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1014Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1014Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1052Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1052Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1283Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1283Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1283DisableTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1283DisableTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1108Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1108Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP2200Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP2200Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1344Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1344Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP1884Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP1884Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEIP2028Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEIP2028Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) IsForked(fn func() *uint64, n *big.Int) bool {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetForkCanonHash(n uint64) common.Hash {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetForkCanonHash(n uint64, h common.Hash) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetForkCanonHashes() map[uint64]common.Hash {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetConsensusEngineType() ctypes.ConsensusEngineT {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) MustSetConsensusEngineType(t ctypes.ConsensusEngineT) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashMinimumDifficulty() *big.Int {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashMinimumDifficulty(i *big.Int) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashDifficultyBoundDivisor() *big.Int {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashDifficultyBoundDivisor(i *big.Int) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashDurationLimit() *big.Int {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashDurationLimit(i *big.Int) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashHomesteadTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashHomesteadTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP2Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP2Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP779Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP779Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP649Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP649Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP1234Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP1234Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP2384Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP2384Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashECIP1010PauseTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashECIP1010PauseTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashECIP1010ContinueTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashECIP1010ContinueTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashECIP1017Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashECIP1017Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashECIP1017EraRounds() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashECIP1017EraRounds(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashEIP100BTransition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashEIP100BTransition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashECIP1041Transition() *uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashECIP1041Transition(n *uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashDifficultyBombDelaySchedule() ctypes.Uint64BigMapEncodesHex {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashDifficultyBombDelaySchedule(m ctypes.Uint64BigMapEncodesHex) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetEthashBlockRewardSchedule() ctypes.Uint64BigMapEncodesHex {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetEthashBlockRewardSchedule(m ctypes.Uint64BigMapEncodesHex) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetCliquePeriod() uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetCliquePeriod(n uint64) error {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) GetCliqueEpoch() uint64 {
	panic("implement me")
}

func (api *PrivateChainConfigAPI) SetCliqueEpoch(n uint64) error {
	panic("implement me")
}
