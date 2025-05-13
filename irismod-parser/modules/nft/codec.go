package nft

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	nft "mods.irisnet.org/modules/nft"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
