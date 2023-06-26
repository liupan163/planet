package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/blog/types"
)

// GetTimeoutPostCount get the total number of timeoutPost
func (k Keeper) GetTimeoutPostCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimeoutPostCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimeoutPostCount set the total number of timeoutPost
func (k Keeper) SetTimeoutPostCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimeoutPostCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimeoutPost appends a timeoutPost in the store with a new id and update the count
func (k Keeper) AppendTimeoutPost(
	ctx sdk.Context,
	timeoutPost types.TimeoutPost,
) uint64 {
	// Create the timeoutPost
	count := k.GetTimeoutPostCount(ctx)

	// Set the ID of the appended value
	timeoutPost.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutPostKey))
	appendedValue := k.cdc.MustMarshal(&timeoutPost)
	store.Set(GetTimeoutPostIDBytes(timeoutPost.Id), appendedValue)

	// Update timeoutPost count
	k.SetTimeoutPostCount(ctx, count+1)

	return count
}

// SetTimeoutPost set a specific timeoutPost in the store
func (k Keeper) SetTimeoutPost(ctx sdk.Context, timeoutPost types.TimeoutPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutPostKey))
	b := k.cdc.MustMarshal(&timeoutPost)
	store.Set(GetTimeoutPostIDBytes(timeoutPost.Id), b)
}

// GetTimeoutPost returns a timeoutPost from its id
func (k Keeper) GetTimeoutPost(ctx sdk.Context, id uint64) (val types.TimeoutPost, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutPostKey))
	b := store.Get(GetTimeoutPostIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimeoutPost removes a timeoutPost from the store
func (k Keeper) RemoveTimeoutPost(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutPostKey))
	store.Delete(GetTimeoutPostIDBytes(id))
}

// GetAllTimeoutPost returns all timeoutPost
func (k Keeper) GetAllTimeoutPost(ctx sdk.Context) (list []types.TimeoutPost) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimeoutPostKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimeoutPost
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimeoutPostIDBytes returns the byte representation of the ID
func GetTimeoutPostIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimeoutPostIDFromBytes returns ID in uint64 format from a byte array
func GetTimeoutPostIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
