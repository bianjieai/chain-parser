package staking

import (
	"github.com/cosmos/cosmos-sdk/x/staking"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(staking.AppModuleBasic{})
}
