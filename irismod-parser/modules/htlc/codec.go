package htlc

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/htlc"
)

func init() {
	codec.RegisterAppModules(htlc.AppModuleBasic{})
}
