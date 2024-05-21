package evm

import (
	"github.com/evmos/ethermint/x/evm"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		evm.AppModuleBasic{},
	)
}
