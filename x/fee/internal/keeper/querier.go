package keeper

import (
	"github.com/okxtuta/anatha-cli/x/fee/internal/types"
	"github.com/okxtuta/go-anatha/codec"
	sdk "github.com/okxtuta/go-anatha/types"
	sdkerrors "github.com/okxtuta/go-anatha/types/errors"
	abci "github.com/tendermint/tendermint/abci/types"
)

const (
	QueryParameters          = "parameters"
	QueryFeeExcludedMessages = "excluded-messages"
)

func NewQuerier(k Keeper) sdk.Querier {
	return func(ctx sdk.Context, path []string, req abci.RequestQuery) ([]byte, error) {
		switch path[0] {
		case QueryParameters:
			return queryParams(ctx, k)

		case QueryFeeExcludedMessages:
			return queryFeeExcludedMessages(ctx, k)

		default:
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, "unknown distribution query endpoint")
		}
	}
}

func queryParams(ctx sdk.Context, k Keeper) ([]byte, error) {
	params := k.GetParams(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, params)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}

func queryFeeExcludedMessages(ctx sdk.Context, k Keeper) ([]byte, error) {
	feeExcludedMessages := k.GetFeeExcludedMessages(ctx)

	res, err := codec.MarshalJSONIndent(types.ModuleCdc, feeExcludedMessages)

	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrJSONMarshal, err.Error())
	}

	return res, nil
}
