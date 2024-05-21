package slashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
