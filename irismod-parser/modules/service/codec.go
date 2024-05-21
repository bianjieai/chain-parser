package service

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/service"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
