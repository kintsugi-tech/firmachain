package keepers

import (
	"path/filepath"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisiskeeper "github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distr "github.com/cosmos/cosmos-sdk/x/distribution"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencekeeper "github.com/cosmos/cosmos-sdk/x/evidence/keeper"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"

	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"

	"github.com/cosmos/cosmos-sdk/x/feegrant"
	feegrantkeeper "github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	govkeeper "github.com/cosmos/cosmos-sdk/x/gov/keeper"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	mintkeeper "github.com/cosmos/cosmos-sdk/x/mint/keeper"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	paramproposal "github.com/cosmos/cosmos-sdk/x/params/types/proposal"
	slashingkeeper "github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingkeeper "github.com/cosmos/cosmos-sdk/x/staking/keeper"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	upgradekeeper "github.com/cosmos/cosmos-sdk/x/upgrade/keeper"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	"github.com/cosmos/ibc-go/v4/modules/apps/transfer"
	ibctransferkeeper "github.com/cosmos/ibc-go/v4/modules/apps/transfer/keeper"
	ibctransfertypes "github.com/cosmos/ibc-go/v4/modules/apps/transfer/types"
	ibcclient "github.com/cosmos/ibc-go/v4/modules/core/02-client"
	ibcclienttypes "github.com/cosmos/ibc-go/v4/modules/core/02-client/types"
	ibcporttypes "github.com/cosmos/ibc-go/v4/modules/core/05-port/types"
	ibchost "github.com/cosmos/ibc-go/v4/modules/core/24-host"
	ibckeeper "github.com/cosmos/ibc-go/v4/modules/core/keeper"

	contractmodulekeeper "github.com/firmachain/firmachain/x/contract/keeper"
	contractmoduletypes "github.com/firmachain/firmachain/x/contract/types"
	nftmodulekeeper "github.com/firmachain/firmachain/x/nft/keeper"
	nftmoduletypes "github.com/firmachain/firmachain/x/nft/types"

	tokenmodulekeeper "github.com/firmachain/firmachain/x/token/keeper"
	tokenmoduletypes "github.com/firmachain/firmachain/x/token/types"

	burnmodulekeeper "github.com/firmachain/firmachain/x/burn/keeper"
	burnmoduletypes "github.com/firmachain/firmachain/x/burn/types"

	"github.com/CosmWasm/wasmd/x/wasm"

	icahost "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host"
	icahostkeeper "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/keeper"
	icahosttypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/host/types"
	icatypes "github.com/cosmos/ibc-go/v4/modules/apps/27-interchain-accounts/types"
)

type AppKeepers struct {
	// keys to access the substores
	keys    map[string]*sdk.KVStoreKey
	tkeys   map[string]*sdk.TransientStoreKey
	memKeys map[string]*sdk.MemoryStoreKey

	// keepers
	AccountKeeper    authkeeper.AccountKeeper
	BankKeeper       bankkeeper.Keeper
	CapabilityKeeper *capabilitykeeper.Keeper
	StakingKeeper    stakingkeeper.Keeper
	SlashingKeeper   slashingkeeper.Keeper
	MintKeeper       mintkeeper.Keeper
	DistrKeeper      distrkeeper.Keeper
	GovKeeper        govkeeper.Keeper
	CrisisKeeper     crisiskeeper.Keeper
	UpgradeKeeper    upgradekeeper.Keeper
	ParamsKeeper     paramskeeper.Keeper
	AuthzKeeper      authzkeeper.Keeper
	IBCKeeper        *ibckeeper.Keeper // IBC Keeper must be a pointer in the app, so we can SetRouter on it correctly
	EvidenceKeeper   evidencekeeper.Keeper
	TransferKeeper   ibctransferkeeper.Keeper
	FeeGrantKeeper   feegrantkeeper.Keeper
	ICAHostKeeper    icahostkeeper.Keeper
	WasmKeeper       wasm.Keeper

	// custom modules
	NftKeeper      nftmodulekeeper.Keeper
	ContractKeeper contractmodulekeeper.Keeper
	TokenKeeper    tokenmodulekeeper.Keeper
	BurnKeeper     burnmodulekeeper.Keeper

	// make scoped keepers public for test purposes
	ScopedIBCKeeper      capabilitykeeper.ScopedKeeper
	ScopedICAHostKeeper  capabilitykeeper.ScopedKeeper
	ScopedTransferKeeper capabilitykeeper.ScopedKeeper
	ScopedWasmKeeper     capabilitykeeper.ScopedKeeper
}

// module account permissions
var maccPerms = map[string][]string{
	authtypes.FeeCollectorName:     nil,
	distrtypes.ModuleName:          nil,
	minttypes.ModuleName:           {authtypes.Minter},
	stakingtypes.BondedPoolName:    {authtypes.Burner, authtypes.Staking},
	stakingtypes.NotBondedPoolName: {authtypes.Burner, authtypes.Staking},
	govtypes.ModuleName:            {authtypes.Burner},
	ibctransfertypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	tokenmoduletypes.ModuleName:    {authtypes.Minter, authtypes.Burner},
	burnmoduletypes.ModuleName:     {authtypes.Burner},
	wasm.ModuleName:                {authtypes.Burner},
	icatypes.ModuleName:            nil,
}

func NewAppKeepers(
	appCodec codec.Codec,
	bApp *baseapp.BaseApp,
	cdc *codec.LegacyAmino,
	maccPerms map[string][]string,
	enabledProposals []wasm.ProposalType,
	appOpts servertypes.AppOptions,
	wasmOpts []wasm.Option,
	homePath string,
	invCheckPeriod uint,
	skipUpgradeHeights map[int64]bool,
) AppKeepers {
	appKeepers := AppKeepers{}

	keys := sdk.NewKVStoreKeys(
		authtypes.StoreKey, banktypes.StoreKey, stakingtypes.StoreKey,
		minttypes.StoreKey, distrtypes.StoreKey, slashingtypes.StoreKey,
		govtypes.StoreKey, paramstypes.StoreKey, ibchost.StoreKey, upgradetypes.StoreKey, feegrant.StoreKey,
		evidencetypes.StoreKey, ibctransfertypes.StoreKey, capabilitytypes.StoreKey, authzkeeper.StoreKey,
		// this line is used by starport scaffolding # stargate/app/storeKey
		nftmoduletypes.StoreKey,
		contractmoduletypes.StoreKey,
		tokenmoduletypes.StoreKey,
		burnmoduletypes.StoreKey,
		wasm.StoreKey,
		icahosttypes.StoreKey,
	)
	tkeys := sdk.NewTransientStoreKeys(paramstypes.TStoreKey)
	memKeys := sdk.NewMemoryStoreKeys(capabilitytypes.MemStoreKey)

	appKeepers.ParamsKeeper = initParamsKeeper(appCodec, cdc, keys[paramstypes.StoreKey], tkeys[paramstypes.TStoreKey])

	// set the BaseApp's parameter store
	bApp.SetParamStore(appKeepers.ParamsKeeper.Subspace(baseapp.Paramspace).WithKeyTable(paramskeeper.ConsensusParamsKeyTable()))

	// add capability keeper and ScopeToModule for ibc module
	appKeepers.CapabilityKeeper = capabilitykeeper.NewKeeper(appCodec, keys[capabilitytypes.StoreKey], memKeys[capabilitytypes.MemStoreKey])

	// grant capabilities for the ibc and ibc-transfer modules
	scopedIBCKeeper := appKeepers.CapabilityKeeper.ScopeToModule(ibchost.ModuleName)
	scopedICAHostKeeper := appKeepers.CapabilityKeeper.ScopeToModule(icahosttypes.SubModuleName)
	scopedTransferKeeper := appKeepers.CapabilityKeeper.ScopeToModule(ibctransfertypes.ModuleName)
	scopedWasmKeeper := appKeepers.CapabilityKeeper.ScopeToModule(wasm.ModuleName)

	// add keepers
	appKeepers.AccountKeeper = authkeeper.NewAccountKeeper(
		appCodec, keys[authtypes.StoreKey], appKeepers.GetSubspace(authtypes.ModuleName), authtypes.ProtoBaseAccount, maccPerms,
	)

	appKeepers.BankKeeper = bankkeeper.NewBaseKeeper(
		appCodec, keys[banktypes.StoreKey], appKeepers.AccountKeeper, appKeepers.GetSubspace(banktypes.ModuleName), appKeepers.ModuleAccountAddrsForBankModule(),
	)
	stakingKeeper := stakingkeeper.NewKeeper(
		appCodec, keys[stakingtypes.StoreKey], appKeepers.AccountKeeper, appKeepers.BankKeeper, appKeepers.GetSubspace(stakingtypes.ModuleName),
	)
	appKeepers.MintKeeper = mintkeeper.NewKeeper(
		appCodec, keys[minttypes.StoreKey], appKeepers.GetSubspace(minttypes.ModuleName), &stakingKeeper,
		appKeepers.AccountKeeper, appKeepers.BankKeeper, authtypes.FeeCollectorName,
	)
	appKeepers.DistrKeeper = distrkeeper.NewKeeper(
		appCodec, keys[distrtypes.StoreKey], appKeepers.GetSubspace(distrtypes.ModuleName), appKeepers.AccountKeeper, appKeepers.BankKeeper,
		&stakingKeeper, authtypes.FeeCollectorName, appKeepers.ModuleAccountAddrs(),
	)
	appKeepers.SlashingKeeper = slashingkeeper.NewKeeper(
		appCodec, keys[slashingtypes.StoreKey], &stakingKeeper, appKeepers.GetSubspace(slashingtypes.ModuleName),
	)
	appKeepers.CrisisKeeper = crisiskeeper.NewKeeper(
		appKeepers.GetSubspace(crisistypes.ModuleName), invCheckPeriod, appKeepers.BankKeeper, authtypes.FeeCollectorName,
	)

	appKeepers.FeeGrantKeeper = feegrantkeeper.NewKeeper(appCodec, keys[feegrant.StoreKey], appKeepers.AccountKeeper)
	appKeepers.UpgradeKeeper = upgradekeeper.NewKeeper(skipUpgradeHeights, keys[upgradetypes.StoreKey], appCodec, homePath, bApp)

	// register the staking hooks
	// NOTE: stakingKeeper above is passed by reference, so that it will contain these hooks
	appKeepers.StakingKeeper = *stakingKeeper.SetHooks(
		stakingtypes.NewMultiStakingHooks(appKeepers.DistrKeeper.Hooks(), appKeepers.SlashingKeeper.Hooks()),
	)

	appKeepers.AuthzKeeper = authzkeeper.NewKeeper(keys[authzkeeper.StoreKey], appCodec, bApp.MsgServiceRouter())

	// ... other modules keepers

	// Create IBC Keeper
	appKeepers.IBCKeeper = ibckeeper.NewKeeper(
		appCodec, keys[ibchost.StoreKey], appKeepers.GetSubspace(ibchost.ModuleName), appKeepers.StakingKeeper, appKeepers.UpgradeKeeper, scopedIBCKeeper,
	)

	// register the proposal types
	govRouter := govtypes.NewRouter()
	govRouter.AddRoute(govtypes.RouterKey, govtypes.ProposalHandler).
		AddRoute(paramproposal.RouterKey, params.NewParamChangeProposalHandler(appKeepers.ParamsKeeper)).
		AddRoute(distrtypes.RouterKey, distr.NewCommunityPoolSpendProposalHandler(appKeepers.DistrKeeper)).
		AddRoute(upgradetypes.RouterKey, upgrade.NewSoftwareUpgradeProposalHandler(appKeepers.UpgradeKeeper)).
		AddRoute(ibcclienttypes.RouterKey, ibcclient.NewClientProposalHandler(appKeepers.IBCKeeper.ClientKeeper))

	// Create Transfer Keepers
	appKeepers.TransferKeeper = ibctransferkeeper.NewKeeper(
		appCodec, keys[ibctransfertypes.StoreKey], appKeepers.GetSubspace(ibctransfertypes.ModuleName),
		appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCKeeper.ChannelKeeper, &appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper, appKeepers.BankKeeper, scopedTransferKeeper,
	)

	transferIBCModule := transfer.NewIBCModule(appKeepers.TransferKeeper)

	appKeepers.ICAHostKeeper = icahostkeeper.NewKeeper(
		appCodec,
		keys[icahosttypes.StoreKey],
		appKeepers.GetSubspace(icahosttypes.SubModuleName),
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		appKeepers.AccountKeeper,
		scopedICAHostKeeper,
		bApp.MsgServiceRouter(),
	)

	icaHostIBCModule := icahost.NewIBCModule(appKeepers.ICAHostKeeper)

	// Create evidence Keeper for to register the IBC light client misbehaviour evidence route
	evidenceKeeper := evidencekeeper.NewKeeper(
		appCodec, keys[evidencetypes.StoreKey], &appKeepers.StakingKeeper, appKeepers.SlashingKeeper,
	)
	// If evidence needs to be handled for the app, set routes in router here and seal
	appKeepers.EvidenceKeeper = *evidenceKeeper

	wasmDir := filepath.Join(homePath, "data")
	wasmConfig, err := wasm.ReadWasmConfig(appOpts)

	if err != nil {
		panic("error : read wasm config: " + err.Error())
	}

	supportedFeatures := "iterator,staking,stargate"

	appKeepers.WasmKeeper = wasm.NewKeeper(
		appCodec,
		keys[wasm.StoreKey],
		appKeepers.GetSubspace(wasm.ModuleName),
		appKeepers.AccountKeeper,
		appKeepers.BankKeeper,
		appKeepers.StakingKeeper,
		appKeepers.DistrKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		appKeepers.IBCKeeper.ChannelKeeper,
		&appKeepers.IBCKeeper.PortKeeper,
		scopedWasmKeeper,
		appKeepers.TransferKeeper,
		bApp.MsgServiceRouter(),
		bApp.GRPCQueryRouter(),
		wasmDir,
		wasmConfig,
		supportedFeatures,
		wasmOpts...,
	)

	// register wasm gov proposal types
	if len(enabledProposals) != 0 {
		govRouter.AddRoute(wasm.RouterKey, wasm.NewWasmProposalHandler(appKeepers.WasmKeeper, enabledProposals))
	}

	appKeepers.GovKeeper = govkeeper.NewKeeper(
		appCodec, keys[govtypes.StoreKey], appKeepers.GetSubspace(govtypes.ModuleName), appKeepers.AccountKeeper, appKeepers.BankKeeper,
		&stakingKeeper, govRouter,
	)

	// Create static IBC router, add transfer route, then set and seal it
	ibcRouter := ibcporttypes.NewRouter()
	ibcRouter.AddRoute(ibctransfertypes.ModuleName, transferIBCModule)
	ibcRouter.AddRoute(wasm.ModuleName, wasm.NewIBCHandler(appKeepers.WasmKeeper, appKeepers.IBCKeeper.ChannelKeeper, appKeepers.IBCKeeper.ChannelKeeper))
	ibcRouter.AddRoute(icahosttypes.SubModuleName, icaHostIBCModule)
	// this line is used by starport scaffolding # ibc/app/router
	appKeepers.IBCKeeper.SetRouter(ibcRouter)

	appKeepers.ContractKeeper = *contractmodulekeeper.NewKeeper(
		appCodec,
		keys[contractmoduletypes.StoreKey],
		keys[contractmoduletypes.MemStoreKey],
	)

	appKeepers.NftKeeper = *nftmodulekeeper.NewKeeper(
		appCodec,
		keys[nftmoduletypes.StoreKey],
		keys[nftmoduletypes.MemStoreKey],
	)

	appKeepers.TokenKeeper = *tokenmodulekeeper.NewKeeper(
		appCodec,
		keys[tokenmoduletypes.StoreKey],
		keys[tokenmoduletypes.MemStoreKey],
		appKeepers.BankKeeper,
	)

	appKeepers.BurnKeeper = *burnmodulekeeper.NewKeeper(
		appCodec,
		keys[burnmoduletypes.StoreKey],
		keys[burnmoduletypes.MemStoreKey],
		appKeepers.BankKeeper,
		appKeepers.AccountKeeper,
	)

	appKeepers.ScopedIBCKeeper = scopedIBCKeeper
	appKeepers.ScopedICAHostKeeper = scopedICAHostKeeper
	appKeepers.ScopedTransferKeeper = scopedTransferKeeper
	appKeepers.ScopedWasmKeeper = scopedWasmKeeper

	return appKeepers
}

// initParamsKeeper init params keeper and its subspaces
func initParamsKeeper(appCodec codec.BinaryCodec, legacyAmino *codec.LegacyAmino, key, tkey sdk.StoreKey) paramskeeper.Keeper {
	paramsKeeper := paramskeeper.NewKeeper(appCodec, legacyAmino, key, tkey)

	paramsKeeper.Subspace(authtypes.ModuleName)
	paramsKeeper.Subspace(banktypes.ModuleName)
	paramsKeeper.Subspace(stakingtypes.ModuleName)
	paramsKeeper.Subspace(minttypes.ModuleName)
	paramsKeeper.Subspace(distrtypes.ModuleName)
	paramsKeeper.Subspace(slashingtypes.ModuleName)
	paramsKeeper.Subspace(govtypes.ModuleName).WithKeyTable(govtypes.ParamKeyTable())
	paramsKeeper.Subspace(crisistypes.ModuleName)
	paramsKeeper.Subspace(ibctransfertypes.ModuleName)
	paramsKeeper.Subspace(ibchost.ModuleName)
	paramsKeeper.Subspace(nftmoduletypes.ModuleName)
	paramsKeeper.Subspace(contractmoduletypes.ModuleName)
	paramsKeeper.Subspace(tokenmoduletypes.ModuleName)
	paramsKeeper.Subspace(burnmoduletypes.ModuleName)
	paramsKeeper.Subspace(wasm.ModuleName)
	paramsKeeper.Subspace(icahosttypes.SubModuleName)

	return paramsKeeper
}

// GetSubspace returns a param subspace for a given module name.
//
// NOTE: This is solely to be used for testing purposes.
func (appKeepers *AppKeepers) GetSubspace(moduleName string) paramstypes.Subspace {
	subspace, _ := appKeepers.ParamsKeeper.GetSubspace(moduleName)
	return subspace
}

func (appKeepers *AppKeepers) ModuleAccountAddrsForBankModule() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	modAccAddrs[authtypes.NewModuleAddress(burnmoduletypes.ModuleName).String()] = false
	return modAccAddrs
}

// ModuleAccountAddrs returns all the app's module account addresses.
func (appKeepers *AppKeepers) ModuleAccountAddrs() map[string]bool {
	modAccAddrs := make(map[string]bool)
	for acc := range maccPerms {
		modAccAddrs[authtypes.NewModuleAddress(acc).String()] = true
	}

	return modAccAddrs
}

// GetMaccPerms returns a copy of the module account permissions
func GetMaccPerms() map[string][]string {
	dupMaccPerms := make(map[string][]string)
	for k, v := range maccPerms {
		dupMaccPerms[k] = v
	}

	return dupMaccPerms
}

func (appKeepers *AppKeepers) GetKeys() map[string]*sdk.KVStoreKey {
	return appKeepers.keys
}

func (appKeepers *AppKeepers) GetTKeys() map[string]*sdk.TransientStoreKey {
	return appKeepers.tkeys
}

func (appKeepers *AppKeepers) GetMemKeys() map[string]*sdk.MemoryStoreKey {
	return appKeepers.memKeys
}
