package evidence

import (
	"github.com/cosmos/cosmos-sdk/x/evidence"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(evidence.AppModuleBasic{})
}
