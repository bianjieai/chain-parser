package group

import (
	group "github.com/cosmos/cosmos-sdk/x/group/module"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(group.AppModuleBasic{})
}
