package ibc_hooks

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/osmosis-labs/osmosis/osmoutils"
	"github.com/osmosis-labs/osmosis/x/ibc-hooks/types"
)

var WasmHookModuleAccountAddr sdk.AccAddress = address.Module(types.ModuleName, []byte("wasm-hook intermediary account"))

func IbcHooksInitGenesis(ctx sdk.Context, ak osmoutils.AccountKeeper) {
	err := osmoutils.CreateModuleAccount(ctx, ak, WasmHookModuleAccountAddr)
	if err != nil {
		panic(err)
	}
}
