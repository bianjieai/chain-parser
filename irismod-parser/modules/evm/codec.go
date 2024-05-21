package evm

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/evmos/ethermint/x/evm"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
