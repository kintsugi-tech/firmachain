package v4

import (
	store "github.com/cosmos/cosmos-sdk/store/types"

	"github.com/firmachain/firmachain/app/upgrades"
)

// UpgradeName defines the on-chain upgrade name for the Firmachain v4 upgrade.
const UpgradeName = "v4"

var Upgrade = upgrades.Upgrade{
	UpgradeName:          UpgradeName,
	CreateUpgradeHandler: CreateV4UpgradeHandler,
	StoreUpgrades:        store.StoreUpgrades{},
}
