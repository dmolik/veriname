package veriname

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dmolik/veriname/x/veriname/keeper"
	"github.com/dmolik/veriname/x/veriname/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the ident
	for _, elem := range genState.IdentList {
		k.SetIdent(ctx, elem)
	}
	// this line is used by starport scaffolding # genesis/module/init
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	genesis.IdentList = k.GetAllIdent(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
