package feegrant

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
)

func init() {
	codec.RegisterAppModules(feegrant.AppModuleBasic{})
}
