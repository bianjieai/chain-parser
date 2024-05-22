package mt

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/mt"
)

func init() {
	codec.RegisterAppModules(mt.AppModuleBasic{})
}
