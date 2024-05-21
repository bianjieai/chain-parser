package params

import (
	"github.com/cosmos/cosmos-sdk/x/params"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
