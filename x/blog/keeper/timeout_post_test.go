package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/blog/keeper"
	"planet/x/blog/types"
)

func createNTimeoutPost(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimeoutPost {
	items := make([]types.TimeoutPost, n)
	for i := range items {
		items[i].Id = keeper.AppendTimeoutPost(ctx, items[i])
	}
	return items
}

func TestTimeoutPostGet(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimeoutPost(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimeoutPost(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTimeoutPostRemove(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimeoutPost(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimeoutPost(ctx, item.Id)
		_, found := keeper.GetTimeoutPost(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimeoutPostGetAll(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimeoutPost(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTimeoutPost(ctx)),
	)
}

func TestTimeoutPostCount(t *testing.T) {
	keeper, ctx := keepertest.BlogKeeper(t)
	items := createNTimeoutPost(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimeoutPostCount(ctx))
}
