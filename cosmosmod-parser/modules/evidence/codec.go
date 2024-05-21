package evidence

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/evidence"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
