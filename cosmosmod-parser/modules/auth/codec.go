package auth

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
