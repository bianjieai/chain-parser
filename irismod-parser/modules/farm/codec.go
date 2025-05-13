package farm

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/farm"
)

func init() {
	codec.RegisterAppModules(farm.AppModuleBasic{})
}
