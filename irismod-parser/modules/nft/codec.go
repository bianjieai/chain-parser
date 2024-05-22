package nft

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/nft"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
