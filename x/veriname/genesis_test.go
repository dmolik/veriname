package veriname_test

import (
	"testing"

	keepertest "github.com/dmolik/veriname/testutil/keeper"
	"github.com/dmolik/veriname/x/veriname"
	"github.com/dmolik/veriname/x/veriname/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		IdentList: []types.Ident{
			{
				Index: "0",
			},
			{
				Index: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.VerinameKeeper(t)
	veriname.InitGenesis(ctx, *k, genesisState)
	got := veriname.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	require.Len(t, got.IdentList, len(genesisState.IdentList))
	require.Subset(t, genesisState.IdentList, got.IdentList)
	// this line is used by starport scaffolding # genesis/test/assert
}
