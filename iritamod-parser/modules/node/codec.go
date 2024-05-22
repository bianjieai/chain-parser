package node

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/iritamod/modules/node"
)

func init() {
	codec.RegisterAppModules(
		node.AppModuleBasic{},
	)
}
