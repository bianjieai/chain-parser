package upgrade

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/upgrade"
)

func init() {
	codec.RegisterAppModules(
		upgrade.AppModuleBasic{},
	)
}
