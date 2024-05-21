package crisis

import (
	"github.com/cosmos/cosmos-sdk/x/crisis"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(crisis.AppModuleBasic{})
}
