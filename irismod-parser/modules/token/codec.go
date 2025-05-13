package token

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/token"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
