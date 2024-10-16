package v0_3_5

import (
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"

	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/firmachain/firmachain/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Firmachain v4 upgrade.
const UpgradeName = "v0_3_5"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV0_3_5UpgradeHandler,
	StoreUpgrades: store.StoreUpgrades{
		// NOTICE: https://github.com/cosmos/ibc-go/issues/1088
		Added: []string{"wasm", icacontrollertypes.StoreKey, icahosttypes.StoreKey},
	},
}
