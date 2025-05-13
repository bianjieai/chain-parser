package random

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/random"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
