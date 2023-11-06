package coinswap

import (
	"github.com/irisnet/irismod/modules/coinswap"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
