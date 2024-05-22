package staking

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/staking"
)

func init() {
	codec.RegisterAppModules(staking.AppModuleBasic{})
}
