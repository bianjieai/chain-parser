package coinswap

import (
	"github.com/irisnet/irismod/modules/coinswap"
	"github.com/bianjieai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(coinswap.AppModuleBasic{})
}
