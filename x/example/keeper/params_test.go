package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/yazhini/example/testutil/keeper"
	"github.com/yazhini/example/x/example/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.ExampleKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
