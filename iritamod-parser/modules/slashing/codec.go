package slashing

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
)

func init() {
	codec.RegisterAppModules(iritaslashing.AppModuleBasic{})
}
