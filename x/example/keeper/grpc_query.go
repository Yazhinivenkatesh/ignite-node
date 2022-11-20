package keeper

import (
	"github.com/yazhini/example/x/example/types"
)

var _ types.QueryServer = Keeper{}
