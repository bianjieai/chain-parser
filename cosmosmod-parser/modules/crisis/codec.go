package crisis

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/crisis"
)

func init() {
	codec.RegisterAppModules(crisis.AppModuleBasic{})
}
