package gold_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "m25/testutil/keeper"
	"m25/testutil/nullify"
	"m25/x/gold"
	"m25/x/gold/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.GoldKeeper(t)
	gold.InitGenesis(ctx, *k, genesisState)
	got := gold.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
