package random

import (
	"github.com/irisnet/irismod/modules/random"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
