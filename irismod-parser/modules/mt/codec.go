package mt

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/mt"
)

func init() {
	codec.RegisterAppModules(mt.AppModuleBasic{})
}
