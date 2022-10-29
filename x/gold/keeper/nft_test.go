package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "m25/testutil/keeper"
	"m25/testutil/nullify"
	"m25/x/gold/keeper"
	"m25/x/gold/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNNft(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Nft {
	items := make([]types.Nft, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetNft(ctx, items[i])
	}
	return items
}

func TestNftGet(t *testing.T) {
	keeper, ctx := keepertest.GoldKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetNft(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestNftRemove(t *testing.T) {
	keeper, ctx := keepertest.GoldKeeper(t)
	items := createNNft(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveNft(ctx,
			item.Index,
		)
		_, found := keeper.GetNft(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestNftGetAll(t *testing.T) {
	keeper, ctx := keepertest.GoldKeeper(t)
	items := createNNft(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllNft(ctx)),
	)
}
