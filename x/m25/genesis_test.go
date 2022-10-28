package m25_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "m25/testutil/keeper"
	"m25/testutil/nullify"
	"m25/x/m25"
	"m25/x/m25/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.M25Keeper(t)
	m25.InitGenesis(ctx, *k, genesisState)
	got := m25.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
