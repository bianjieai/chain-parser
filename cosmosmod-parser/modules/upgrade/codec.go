package upgrade

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/upgrade"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
