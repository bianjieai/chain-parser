package nft

import (
	nft "github.com/irisnet/irismod/modules/nft/module"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(nft.AppModuleBasic{})
}
