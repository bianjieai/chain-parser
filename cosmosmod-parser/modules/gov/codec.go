package gov

import (
	"github.com/cosmos/cosmos-sdk/x/gov"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
	)
}
