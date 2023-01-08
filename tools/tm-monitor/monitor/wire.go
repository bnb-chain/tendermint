package monitor

import (
	ctypes "github.com/bnb-chain/tendermint/rpc/core/types"
	amino "github.com/tendermint/go-amino"
)

var cdc = amino.NewCodec()

func init() {
	ctypes.RegisterAmino(cdc)
}
