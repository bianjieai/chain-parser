package record

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/record"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
