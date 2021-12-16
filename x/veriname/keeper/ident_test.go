package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/dmolik/veriname/testutil/keeper"
	"github.com/dmolik/veriname/x/veriname/keeper"
	"github.com/dmolik/veriname/x/veriname/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNIdent(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Ident {
	items := make([]types.Ident, n)
	for i := range items {
		items[i].Index = strconv.Itoa(i)

		keeper.SetIdent(ctx, items[i])
	}
	return items
}

func TestIdentGet(t *testing.T) {
	keeper, ctx := keepertest.VerinameKeeper(t)
	items := createNIdent(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetIdent(ctx,
			item.Index,
		)
		require.True(t, found)
		require.Equal(t, item, rst)
	}
}
func TestIdentRemove(t *testing.T) {
	keeper, ctx := keepertest.VerinameKeeper(t)
	items := createNIdent(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveIdent(ctx,
			item.Index,
		)
		_, found := keeper.GetIdent(ctx,
			item.Index,
		)
		require.False(t, found)
	}
}

func TestIdentGetAll(t *testing.T) {
	keeper, ctx := keepertest.VerinameKeeper(t)
	items := createNIdent(keeper, ctx, 10)
	require.ElementsMatch(t, items, keeper.GetAllIdent(ctx))
}
