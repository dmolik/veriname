package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/dmolik/veriname/x/veriname/types"
)

// SetIdent set a specific ident in the store from its index
func (k Keeper) SetIdent(ctx sdk.Context, ident types.Ident) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentKeyPrefix))
	b := k.cdc.MustMarshal(&ident)
	store.Set(types.IdentKey(
		ident.Index,
	), b)
}

// GetIdent returns a ident from its index
func (k Keeper) GetIdent(
	ctx sdk.Context,
	index string,

) (val types.Ident, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentKeyPrefix))

	b := store.Get(types.IdentKey(
		index,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveIdent removes a ident from the store
func (k Keeper) RemoveIdent(
	ctx sdk.Context,
	index string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentKeyPrefix))
	store.Delete(types.IdentKey(
		index,
	))
}

// GetAllIdent returns all ident
func (k Keeper) GetAllIdent(ctx sdk.Context) (list []types.Ident) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.IdentKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Ident
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
