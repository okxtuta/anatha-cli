package keeper

import (
	"github.com/okxtuta/anatha-cli/x/hra/internal/types"
	sdk "github.com/okxtuta/go-anatha/types"
)

var _ types.NameHooks = Keeper{}

func (k Keeper) AfterFirstNameCreated(ctx sdk.Context, address sdk.AccAddress) error {
	if k.hooks != nil {
		err := k.hooks.AfterFirstNameCreated(ctx, address)
		if err != nil {
			return err
		}
	}
	return nil
}

func (k Keeper) AfterLastNameRemoved(ctx sdk.Context, address sdk.AccAddress) error {
	if k.hooks != nil {
		err := k.hooks.AfterLastNameRemoved(ctx, address)
		if err != nil {
			return err
		}
	}
	return nil
}
