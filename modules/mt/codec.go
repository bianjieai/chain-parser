package mt

import (
	"github.com/irisnet/irismod/modules/mt"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(mt.AppModuleBasic{})
}
