package htlc

import (
	"github.com/irisnet/irismod/modules/htlc"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
