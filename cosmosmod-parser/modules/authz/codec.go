package authz

import (
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(authz.AppModuleBasic{})
}
