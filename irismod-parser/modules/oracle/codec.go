package oracle

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/oracle"
)

func init() {
	codec.RegisterAppModules(oracle.AppModuleBasic{})
}
