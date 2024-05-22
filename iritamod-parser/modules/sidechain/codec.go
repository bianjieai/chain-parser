package sidechain

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/side-chain"
)

func init() {
	codec.RegisterAppModules(sidechain.AppModuleBasic{})
}
