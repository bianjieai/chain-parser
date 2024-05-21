package params

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/params"
)

func init() {
	codec.RegisterAppModules(params.AppModuleBasic{})
}
