package test

import (
	"testing"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store"
	storetypes "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramskeeper "github.com/cosmos/cosmos-sdk/x/params/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	porttypes "github.com/cosmos/ibc-go/v5/modules/core/05-port/types"
	"github.com/golang/mock/gomock"
	"github.com/strangelove-ventures/packet-forward-middleware/v5/router"
	"github.com/strangelove-ventures/packet-forward-middleware/v5/router/keeper"
	"github.com/strangelove-ventures/packet-forward-middleware/v5/router/types"
	"github.com/strangelove-ventures/packet-forward-middleware/v5/test/mock"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmdb "github.com/tendermint/tm-db"
)

func NewTestSetup(t *testing.T, ctl *gomock.Controller) *TestSetup {
	initializer := newInitializer()

	transferKeeperMock := mock.NewMockTransferKeeper(ctl)
	channelKeeperMock := mock.NewMockChannelKeeper(ctl)
	distributionKeeperMock := mock.NewMockDistributionKeeper(ctl)
	bankKeeperMock := mock.NewMockBankKeeper(ctl)
	ibcModuleMock := mock.NewMockIBCModule(ctl)
	ics4WrapperMock := mock.NewMockICS4Wrapper(ctl)
	portKeeperMock := mock.NewMockPortKeeper(ctl)

	paramsKeeper := initializer.paramsKeeper()
	routerKeeper := initializer.routerKeeper(paramsKeeper, transferKeeperMock, channelKeeperMock, distributionKeeperMock, bankKeeperMock, portKeeperMock, ics4WrapperMock)
	//routerModule := initializer.routerModule(routerKeeper)

	require.NoError(t, initializer.StateStore.LoadLatestVersion())

	routerKeeper.SetParams(initializer.Ctx, types.DefaultParams())

	return &TestSetup{
		Initializer: initializer,

		Keepers: &testKeepers{
			ParamsKeeper: &paramsKeeper,
			RouterKeeper: routerKeeper,
		},

		Mocks: &testMocks{
			TransferKeeperMock:     transferKeeperMock,
			ChannelKeeperMock:      channelKeeperMock,
			DistributionKeeperMock: distributionKeeperMock,
			IBCModuleMock:          ibcModuleMock,
			PortKeeperMock:         portKeeperMock,
			ICS4WrapperMock:        ics4WrapperMock,
		},

		ForwardMiddleware: initializer.forwardMiddleware(ibcModuleMock, routerKeeper, 0, keeper.DefaultForwardTransferPacketTimeoutTimestamp, keeper.DefaultRefundTransferPacketTimeoutTimestamp),
	}
}

type TestSetup struct {
	Initializer initializer

	Keepers *testKeepers
	Mocks   *testMocks

	ForwardMiddleware router.IBCMiddleware
}

type testKeepers struct {
	ParamsKeeper *paramskeeper.Keeper
	RouterKeeper *keeper.Keeper
}

type testMocks struct {
	TransferKeeperMock     *mock.MockTransferKeeper
	ChannelKeeperMock      *mock.MockChannelKeeper
	DistributionKeeperMock *mock.MockDistributionKeeper
	IBCModuleMock          *mock.MockIBCModule
	PortKeeperMock         *mock.MockPortKeeper
	ICS4WrapperMock        *mock.MockICS4Wrapper
}

type initializer struct {
	DB         *tmdb.MemDB
	StateStore store.CommitMultiStore
	Ctx        sdk.Context
	Marshaler  codec.Codec
	Amino      *codec.LegacyAmino
}

// Create an initializer with in memory database and default codecs
func newInitializer() initializer {
	logger := log.TestingLogger()
	logger.Debug("initializing test setup")

	db := tmdb.NewMemDB()
	stateStore := store.NewCommitMultiStore(db)

	ctx := sdk.NewContext(stateStore, tmproto.Header{}, false, logger)
	interfaceRegistry := cdctypes.NewInterfaceRegistry()
	amino := codec.NewLegacyAmino()
	marshaler := codec.NewProtoCodec(interfaceRegistry)

	return initializer{
		DB:         db,
		StateStore: stateStore,
		Ctx:        ctx,
		Marshaler:  marshaler,
		Amino:      amino,
	}
}

func (i initializer) paramsKeeper() paramskeeper.Keeper {
	storeKey := sdk.NewKVStoreKey(paramstypes.StoreKey)
	transientStoreKey := sdk.NewTransientStoreKey(paramstypes.TStoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)
	i.StateStore.MountStoreWithDB(transientStoreKey, storetypes.StoreTypeTransient, i.DB)

	paramsKeeper := paramskeeper.NewKeeper(i.Marshaler, i.Amino, storeKey, transientStoreKey)

	return paramsKeeper
}

func (i initializer) routerKeeper(
	paramsKeeper paramskeeper.Keeper,
	transferKeeper types.TransferKeeper,
	channelKeeper types.ChannelKeeper,
	distributionKeeper types.DistributionKeeper,
	bankKeeper types.BankKeeper,
	portKeeper types.PortKeeper,
	ics4Wrapper porttypes.ICS4Wrapper,
) *keeper.Keeper {
	storeKey := sdk.NewKVStoreKey(types.StoreKey)
	i.StateStore.MountStoreWithDB(storeKey, storetypes.StoreTypeIAVL, i.DB)

	subspace := paramsKeeper.Subspace(types.ModuleName)
	routerKeeper := keeper.NewKeeper(
		i.Marshaler,
		storeKey,
		subspace,
		transferKeeper,
		channelKeeper,
		distributionKeeper,
		bankKeeper,
		portKeeper,
		ics4Wrapper,
	)

	return routerKeeper
}

func (i initializer) routerModule(routerKeeper *keeper.Keeper) router.AppModule {
	return router.NewAppModule(routerKeeper)
}

func (i initializer) forwardMiddleware(app porttypes.IBCModule, k *keeper.Keeper, retriesOnTimeout uint8, forwardTimeout time.Duration, refundTimeout time.Duration) router.IBCMiddleware {
	return router.NewIBCMiddleware(app, k, retriesOnTimeout, forwardTimeout, refundTimeout)
}
