package bank

import (
	"github.com/cosmos/cosmos-sdk/x/bank"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
