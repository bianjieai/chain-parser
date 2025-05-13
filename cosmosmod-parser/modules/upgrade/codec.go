package upgrade

import (
	"cosmossdk.io/x/upgrade"
	"github.com/bianjieai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
