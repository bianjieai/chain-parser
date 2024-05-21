package authz

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	authz "github.com/cosmos/cosmos-sdk/x/authz/module"
)

func init() {
	codec.RegisterAppModules(authz.AppModuleBasic{})
}
