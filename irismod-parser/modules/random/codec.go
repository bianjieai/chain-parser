package random

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/random"
)

func init() {
	codec.RegisterAppModules(random.AppModuleBasic{})
}
