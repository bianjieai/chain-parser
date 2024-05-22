package distribution

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/distribution"
)

func init() {
	codec.RegisterAppModules(distribution.AppModuleBasic{})
}
