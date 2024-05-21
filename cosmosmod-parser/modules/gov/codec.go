package gov

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/gov"
)

func init() {
	codec.RegisterAppModules(
		gov.AppModuleBasic{},
	)
}
