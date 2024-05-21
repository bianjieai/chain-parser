package feemarket

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/evmos/ethermint/x/feemarket"
)

func init() {
	codec.RegisterAppModules(
		feemarket.AppModuleBasic{},
	)
}
