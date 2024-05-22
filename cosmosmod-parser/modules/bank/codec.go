package bank

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/cosmos/cosmos-sdk/x/bank"
)

func init() {
	codec.RegisterAppModules(bank.AppModuleBasic{})
}
