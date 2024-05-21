package token

import (
	"github.com/irisnet/irismod/modules/token"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(token.AppModuleBasic{})
}
