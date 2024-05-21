package slashing

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/slashing"
)

func init() {
	codec.RegisterAppModules(slashing.AppModuleBasic{})
}
