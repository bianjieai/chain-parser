package params

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/params"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
