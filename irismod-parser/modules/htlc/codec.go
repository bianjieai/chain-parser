package htlc

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/htlc"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
