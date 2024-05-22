package perm

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/perm"
)

func init() {
	codec.RegisterAppModules(perm.AppModuleBasic{})
}
