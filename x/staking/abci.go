package staking

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/okxtuta/anatha-cli/x/staking/keeper"
	sdk "github.com/okxtuta/go-anatha/types"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
	k.DumpTickets(ctx)
	k.TrackHistoricalInfo(ctx)
}

// Called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	return k.BlockValidatorUpdates(ctx)
}
