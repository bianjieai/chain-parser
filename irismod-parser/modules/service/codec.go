package service

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"mods.irisnet.org/modules/service"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
