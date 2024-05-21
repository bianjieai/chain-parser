package farm

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/farm"
)

func init() {
	codec.RegisterAppModules(farm.AppModuleBasic{})
}
