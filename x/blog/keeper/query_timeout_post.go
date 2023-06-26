package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"planet/x/blog/types"
)

func (k Keeper) TimeoutPostAll(goCtx context.Context, req *types.QueryAllTimeoutPostRequest) (*types.QueryAllTimeoutPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timeoutPosts []types.TimeoutPost
	ctx := sdk.UnwrapSDKContext(goCtx)

	store := ctx.KVStore(k.storeKey)
	timeoutPostStore := prefix.NewStore(store, types.KeyPrefix(types.TimeoutPostKey))

	pageRes, err := query.Paginate(timeoutPostStore, req.Pagination, func(key []byte, value []byte) error {
		var timeoutPost types.TimeoutPost
		if err := k.cdc.Unmarshal(value, &timeoutPost); err != nil {
			return err
		}

		timeoutPosts = append(timeoutPosts, timeoutPost)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimeoutPostResponse{TimeoutPost: timeoutPosts, Pagination: pageRes}, nil
}

func (k Keeper) TimeoutPost(goCtx context.Context, req *types.QueryGetTimeoutPostRequest) (*types.QueryGetTimeoutPostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(goCtx)
	timeoutPost, found := k.GetTimeoutPost(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimeoutPostResponse{TimeoutPost: timeoutPost}, nil
}
