package evidence

import (
	"cosmossdk.io/x/evidence"
	"github.com/bianjieai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
