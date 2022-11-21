package registry

import (
	"github.com/cerc-io/laconicd/x/registry/keeper"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// BeginBlocker will persist the current header and validator set as a historical entry
// and prune the oldest entry based on the HistoricalEntries parameter
func BeginBlocker(ctx sdk.Context, k keeper.Keeper) {
}

// EndBlocker Called every block, update validator set
func EndBlocker(ctx sdk.Context, k keeper.Keeper) []abci.ValidatorUpdate {
	k.ProcessRecordExpiryQueue(ctx)
	k.ProcessAuthorityExpiryQueue(ctx)

	return []abci.ValidatorUpdate{}
}
