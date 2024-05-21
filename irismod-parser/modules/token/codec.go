package token

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/token"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
