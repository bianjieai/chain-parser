package nft

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	nft "github.com/irisnet/irismod/modules/nft/module"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
