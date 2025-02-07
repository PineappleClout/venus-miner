// +build calibnet

package build

import (
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"
	"github.com/filecoin-project/venus-miner/chain/actors/policy"
)

const UpgradeBreezeHeight = -1
const BreezeGasTampingDuration = 120

const UpgradeSmokeHeight = -2

const UpgradeIgnitionHeight = -3
const UpgradeRefuelHeight = -4

var UpgradeActorsV2Height = abi.ChainEpoch(30)

const UpgradeTapeHeight = 60

const UpgradeLiftoffHeight = -5

const UpgradeKumquatHeight = 90

const UpgradeCalicoHeight = 92000
const UpgradePersianHeight = UpgradeCalicoHeight + (builtin2.EpochsInHour * 60)

// 2020-12-17T19:00:00Z
const UpgradeClausHeight = 161386

// 2021-01-17T19:00:00Z
const UpgradeOrangeHeight = 250666

// 2021-01-28T21:00:00Z
const UpgradeActorsV3Height = 282586

func init() {
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(10 << 30))
	policy.SetSupportedProofTypes(
		abi.RegisteredSealProof_StackedDrg512MiBV1,
		abi.RegisteredSealProof_StackedDrg32GiBV1,
		abi.RegisteredSealProof_StackedDrg64GiBV1,
	)

	SetAddressNetwork(address.Testnet)

	Devnet = true

	BuildType = BuildCalibnet
}

const BlockDelaySecs = uint64(builtin2.EpochDurationSeconds)

const PropagationDelaySecs = uint64(12)
