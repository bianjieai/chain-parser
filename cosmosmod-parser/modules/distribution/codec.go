package distribution

import (
	"github.com/cosmos/cosmos-sdk/x/distribution"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(distribution.AppModuleBasic{})
}
