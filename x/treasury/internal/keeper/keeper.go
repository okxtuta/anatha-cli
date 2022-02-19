package keeper

import (
	"fmt"
	"github.com/okxtuta/go-anatha/x/auth"
	"github.com/okxtuta/go-anatha/x/bank"
	"github.com/okxtuta/go-anatha/x/params"
	"github.com/okxtuta/go-anatha/x/supply"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/okxtuta/anatha-cli/x/treasury/internal/types"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
)

type Keeper struct {
	storeKey      sdk.StoreKey
	cdc           *codec.Codec
	paramspace    params.Subspace
	supplyKeeper  supply.Keeper
	AccountKeeper auth.AccountKeeper
	BankKeeper    bank.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, paramspace params.Subspace, supplyKeeper supply.Keeper, accountKeeper auth.AccountKeeper, bankKeeper bank.Keeper) Keeper {
	keeper := Keeper{
		storeKey:      key,
		cdc:           cdc,
		paramspace:    paramspace.WithKeyTable(types.ParamKeyTable()),
		supplyKeeper:  supplyKeeper,
		AccountKeeper: accountKeeper,
		BankKeeper:    bankKeeper,
	}
	return keeper
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) GetTreasury(ctx sdk.Context) (minter types.Treasury) {
	store := ctx.KVStore(k.storeKey)
	b := store.Get(types.TreasuryKey)
	if b == nil {
		panic("stored treasury should not have been nil")
	}

	k.cdc.MustUnmarshalBinaryBare(b, &minter)
	return
}

func (k Keeper) SetTreasury(ctx sdk.Context, minter types.Treasury) {
	store := ctx.KVStore(k.storeKey)
	b := k.cdc.MustMarshalBinaryBare(minter)
	store.Set(types.TreasuryKey, b)
}
