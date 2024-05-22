package evm

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/tharsis/ethermint/x/evm"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
