package oracle

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/oracle"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
