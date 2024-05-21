package group

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	group "github.com/cosmos/cosmos-sdk/x/group/module"
)

func init() {
	codec.RegisterAppModules(group.AppModuleBasic{})
}
