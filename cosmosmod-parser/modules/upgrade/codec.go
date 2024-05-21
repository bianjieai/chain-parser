package upgrade

import (
	"github.com/cosmos/cosmos-sdk/x/upgrade"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
