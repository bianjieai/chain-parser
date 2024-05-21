package auth

import (
	"github.com/cosmos/cosmos-sdk/x/auth"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(auth.AppModuleBasic{})
}
