package feegrant

import (
	feegrant "github.com/cosmos/cosmos-sdk/x/feegrant/module"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(feegrant.AppModuleBasic{})
}
