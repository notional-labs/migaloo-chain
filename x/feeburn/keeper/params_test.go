package keeper_test

import (
	"testing"

	testkeeper "github.com/White-Whale-Defi-Platform/migaloo-chain/testutil/keeper"

	"github.com/White-Whale-Defi-Platform/migaloo-chain/x/feeburn/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.FeeburnKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
