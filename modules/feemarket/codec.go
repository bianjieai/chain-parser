package feemarket

import (
	"github.com/evmos/ethermint/x/feemarket"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		feemarket.AppModuleBasic{},
	)
}
