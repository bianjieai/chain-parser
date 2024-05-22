package coinswap

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/coinswap"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
