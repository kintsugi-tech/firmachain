package v035

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/firmachain/firmachain/app/keepers"

	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	ica "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts"
	icacontrollertypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/controller/types"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
)

func CreateV035UpgradeHandler(
	mm *module.Manager,
	cfg module.Configurator,
	_ *keepers.AppKeepers,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, _ module.VersionMap) (module.VersionMap, error) {
		// cosmos 0.44.5 consensus version (old)
		fromVM := map[string]uint64{
			"auth":         2,
			"authz":        1,
			"bank":         2,
			"capability":   1,
			"crisis":       1,
			"distribution": 2,
			"evidence":     1,
			"gov":          2,
			"mint":         1,
			"params":       1,
			"slashing":     2,
			"staking":      2,
			"upgrade":      1,
			"vesting":      1,
			"ibc":          1,
			"genutil":      1,
			"transfer":     1,
			"feegrant":     1,
			"contract":     1,
			"nft":          1,
			"token":        1,
			"burn":         1,
		}

		// Add Interchain Accounts host module
		// set the ICS27 consensus version so InitGenesis is not run
		fromVM[icatypes.ModuleName] = mm.Modules[icatypes.ModuleName].ConsensusVersion()

		// create ICS27 Controller submodule params, controller module not enabled.
		controllerParams := icacontrollertypes.Params{}

		// create ICS27 Host submodule params
		hostParams := icahosttypes.Params{
			HostEnabled: true,
			AllowMessages: []string{
				"/cosmos.bank.v1beta1.MsgSend",
				"/cosmos.bank.v1beta1.MsgMultiSend",
				"/cosmos.staking.v1beta1.MsgDelegate",
				"/cosmos.staking.v1beta1.MsgUndelegate",
				"/cosmos.staking.v1beta1.MsgBeginRedelegate",
				"/cosmos.staking.v1beta1.MsgCreateValidator",
				"/cosmos.staking.v1beta1.MsgEditValidator",
				"/cosmos.vesting.v1beta1.MsgCreateVestingAccount",
				"/cosmos.distribution.v1beta1.MsgWithdrawDelegatorReward",
				"/cosmos.distribution.v1beta1.MsgSetWithdrawAddress",
				"/cosmos.distribution.v1beta1.MsgWithdrawValidatorCommission",
				"/cosmos.distribution.v1beta1.MsgFundCommunityPool",
				"/cosmos.feegrant.v1beta1.MsgGrantAllowance",
				"/cosmos.feegrant.v1beta1.MsgRevokeAllowance",
				"/cosmos.gov.v1beta1.MsgVote",
				"/cosmos.gov.v1beta1.MsgVoteWeighted",
				"/cosmos.gov.v1beta1.MsgSubmitProposal",
				"/cosmos.gov.v1beta1.MsgDeposit",
				"/cosmos.authz.v1beta1.MsgExec",
				"/cosmos.authz.v1beta1.MsgGrant",
				"/cosmos.authz.v1beta1.MsgRevoke",
				"/cosmwasm.wasm.v1.MsgStoreCode",
				"/cosmwasm.wasm.v1.MsgInstantiateContract",
				"/cosmwasm.wasm.v1.InstantiateContract2",
				"/cosmwasm.wasm.v1.MsgExecuteContract",
				"/ibc.applications.transfer.v1.MsgTransfer",
			},
		}

		// InitModule will initialize the interchain accounts moudule. It should only be
		// called once and as an alternative to InitGenesis.
		// Use a type assertion to convert mm.Modules[icatypes.ModuleName] to ica.AppModule
		var icaModule ica.AppModule
		if mod, ok := mm.Modules[icatypes.ModuleName].(ica.AppModule); ok {
			icaModule = mod
		} else {
			// Handle the case where the assertion fails
			panic("Cannot initialize the ica module.")
		}
		icaModule.InitModule(ctx, controllerParams, hostParams)

		return mm.RunMigrations(ctx, cfg, fromVM)
	}
}
