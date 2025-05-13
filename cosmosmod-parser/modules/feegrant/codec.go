package feegrant

import (
	feegrant "cosmossdk.io/x/feegrant/module"
	"github.com/bianjieai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(feegrant.AppModuleBasic{})
}
