package farm

import (
	"github.com/irisnet/irismod/modules/farm"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(farm.AppModuleBasic{})
}
