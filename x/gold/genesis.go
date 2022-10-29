package gold

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"m25/x/gold/keeper"
	"m25/x/gold/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the nft
	for _, elem := range genState.NftList {
		k.SetNft(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.NftList = k.GetAllNft(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
