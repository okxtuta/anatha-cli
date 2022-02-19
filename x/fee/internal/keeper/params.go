package keeper

import (
	"github.com/okxtuta/anatha-cli/x/fee/internal/types"
	sdk "github.com/okxtuta/go-anatha/types"
)

// FeePercentage
func (k Keeper) FeePercentage(ctx sdk.Context) (res sdk.Dec) {
	k.paramspace.Get(ctx, types.KeyFeePercentage, &res)
	return
}

// MinimumFee
func (k Keeper) MinimumFee(ctx sdk.Context) (res sdk.Coins) {
	k.paramspace.Get(ctx, types.KeyMinimumFee, &res)
	return
}

// Maximum Fee
func (k Keeper) MaximumFee(ctx sdk.Context) (res sdk.Coins) {
	k.paramspace.Get(ctx, types.KeyMaximumFee, &res)
	return
}

func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.paramspace.GetParamSet(ctx, &params)
	return params
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.paramspace.SetParamSet(ctx, &params)
}
