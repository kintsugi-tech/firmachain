package v05

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/firmachain/firmachain/v05/app/keepers"

	packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v7/router/types"
	icqtypes "github.com/cosmos/ibc-apps/modules/async-icq/v7/types"
	ibchookstypes "github.com/cosmos/ibc-apps/modules/ibc-hooks/v7/types"
	icacontrollertypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/host/types"
	ibcfeetypes "github.com/cosmos/ibc-go/v7/modules/apps/29-fee/types"
	ibcexported "github.com/cosmos/ibc-go/v7/modules/core/exported"

	// SDK v47 modules
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v7/modules/apps/transfer/types"
)

func CreateV0_5_0UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	keepers *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger().With("upgrade", UpgradeName)

		// === Params migration ===
		// https://github.com/cosmos/cosmos-sdk/pull/12363/files
		// Set param key table for params module migration
		for _, subspace := range keepers.ParamsKeeper.GetSubspaces() {
			subspace := subspace

			var keyTable paramstypes.KeyTable
			switch subspace.Name() {
			// SDK
			case authtypes.ModuleName:
				keyTable = authtypes.ParamKeyTable() //nolint:staticcheck
			case banktypes.ModuleName:
				keyTable = banktypes.ParamKeyTable() //nolint:staticcheck
			case crisistypes.ModuleName:
				keyTable = crisistypes.ParamKeyTable() //nolint:staticcheck
			case distrtypes.ModuleName:
				keyTable = distrtypes.ParamKeyTable() //nolint:staticcheck
			case govtypes.ModuleName:
				keyTable = govv1.ParamKeyTable() //nolint:staticcheck
			case minttypes.ModuleName:
				keyTable = minttypes.ParamKeyTable() //nolint:staticcheck
			case slashingtypes.ModuleName:
				keyTable = slashingtypes.ParamKeyTable() //nolint:staticcheck
			case stakingtypes.ModuleName:
				keyTable = stakingtypes.ParamKeyTable() //nolint:staticcheck

			// IBC
			case ibctransfertypes.ModuleName:
				keyTable = ibctransfertypes.ParamKeyTable()
			case icahosttypes.SubModuleName:
				keyTable = icahosttypes.ParamKeyTable()

			// Wasm
			case wasmtypes.ModuleName:
				keyTable = wasmtypes.ParamKeyTable() //nolint:staticcheck

			default:
				continue
			}
			if !subspace.HasKeyTable() {
				subspace.WithKeyTable(keyTable)
			}
		}
		// Migrate Tendermint consensus parameters from x/params module to a deprecated x/consensus module.
		// The old params module is required to still be imported in your app.go in order to handle this migration.
		baseAppLegacySS := keepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramstypes.ConsensusParamsKeyTable())
		baseapp.MigrateParams(ctx, baseAppLegacySS, &keepers.ConsensusParamsKeeper)

		// === New params ===
		// https://github.com/cosmos/ibc-go/blob/v7.1.0/docs/migrations/v7-to-v7_1.md
		// explicitly update the IBC 02-client params, adding the localhost client type
		newIBCCoreParams := keepers.IBCKeeper.ClientKeeper.GetParams(ctx)
		newIBCCoreParams.AllowedClients = append(newIBCCoreParams.AllowedClients, ibcexported.Localhost)
		// ICA Host
		newIcaHostParams := icahosttypes.Params{
			HostEnabled: true,
			// https://github.com/cosmos/ibc-go/blob/v4.2.0/docs/apps/interchain-accounts/parameters.md#allowmessages
			AllowMessages: []string{"*"},
		}
		// ICA Controller
		newIcaControllerParams := icacontrollertypes.Params{ControllerEnabled: true}
		// IBC PFM
		newPFMParams := packetforwardtypes.DefaultParams()
		// ICQ
		newICQParams := icqtypes.NewParams(true, nil)
		// IBC Fee
		logger.Info(fmt.Sprintf("ibcfee module version %s set", fmt.Sprint(vm[ibcfeetypes.ModuleName])))
		// IBC Hooks
		logger.Info(fmt.Sprintf("ibchooks module version %s set", fmt.Sprint(vm[ibchookstypes.ModuleName])))

		// ==== Run migration ====
		logger.Info(fmt.Sprintf("pre migrate version map: %v", vm))
		versionMap, err := mm.RunMigrations(ctx, cfg, vm)
		logger.Info(fmt.Sprintf("post migrate version map: %v", versionMap))

		// New modules run AFTER the migrations, so to set the correct params after the default.

		// ==== Set Params ====
		keepers.IBCKeeper.ClientKeeper.SetParams(ctx, newIBCCoreParams)
		logger.Info(fmt.Sprintf("ibc core: ICQKeeper params set"))
		keepers.ICAHostKeeper.SetParams(ctx, newIcaHostParams)
		logger.Info(fmt.Sprintf("icahost: ICAHostKeeper params set"))
		keepers.ICAControllerKeeper.SetParams(ctx, newIcaControllerParams)
		logger.Info(fmt.Sprintf("icacontroller: ICAControllerKeeper params set"))
		keepers.PacketForwardKeeper.SetParams(ctx, newPFMParams)
		logger.Info(fmt.Sprintf("packetforward: PacketForwardKeeper params set"))
		keepers.ICQKeeper.SetParams(ctx, newICQParams)
		logger.Info(fmt.Sprintf("icq: ICQKeeper params set"))

		return versionMap, err
	}
}
