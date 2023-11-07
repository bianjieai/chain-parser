package service

import (
	"github.com/irisnet/irismod/modules/service"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(service.AppModuleBasic{})
}
