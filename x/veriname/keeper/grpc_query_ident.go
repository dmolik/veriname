package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/dmolik/veriname/x/veriname/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) IdentAll(c context.Context, req *types.QueryAllIdentRequest) (*types.QueryAllIdentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var idents []types.Ident
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	identStore := prefix.NewStore(store, types.KeyPrefix(types.IdentKeyPrefix))

	pageRes, err := query.Paginate(identStore, req.Pagination, func(key []byte, value []byte) error {
		var ident types.Ident
		if err := k.cdc.Unmarshal(value, &ident); err != nil {
			return err
		}

		idents = append(idents, ident)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllIdentResponse{Ident: idents, Pagination: pageRes}, nil
}

func (k Keeper) Ident(c context.Context, req *types.QueryGetIdentRequest) (*types.QueryGetIdentResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetIdent(
		ctx,
		req.Index,
	)
	if !found {
		return nil, status.Error(codes.InvalidArgument, "not found")
	}

	return &types.QueryGetIdentResponse{Ident: val}, nil
}
