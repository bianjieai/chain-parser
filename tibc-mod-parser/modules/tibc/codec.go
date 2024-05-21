package tibc

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer"
	tibccore "github.com/bianjieai/tibc-go/modules/tibc/core"
)

func init() {
	codec.RegisterAppModules(
		tibctranfer.AppModuleBasic{},
		tibccore.AppModuleBasic{},
	)
}
