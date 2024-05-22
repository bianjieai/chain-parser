package identity

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/identity"
)

func init() {
	codec.RegisterAppModules(identity.AppModuleBasic{})
}
