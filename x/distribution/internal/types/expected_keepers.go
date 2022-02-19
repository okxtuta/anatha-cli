package types

import sdk "github.com/okxtuta/go-anatha/types"

type NameHooks interface {
	AfterFirstNameCreated(ctx sdk.Context, address sdk.AccAddress) error
	AfterLastNameRemoved(ctx sdk.Context, address sdk.AccAddress) error
}
