package oracle

import (
	"github.com/irisnet/irismod/modules/oracle"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
