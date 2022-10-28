package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "m25/testutil/keeper"
	"m25/x/m25/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.M25Keeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
