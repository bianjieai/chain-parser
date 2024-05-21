package ibc

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	nfttransfer "github.com/bianjieai/nft-transfer"
	ibctransfer "github.com/cosmos/ibc-go/v7/modules/apps/transfer"
	ibc "github.com/cosmos/ibc-go/v7/modules/core"
	ibcclients "github.com/cosmos/ibc-go/v7/modules/light-clients/07-tendermint"
)

func init() {
	codec.RegisterAppModules(
		ibc.AppModuleBasic{},
		ibctransfer.AppModuleBasic{},
		nfttransfer.AppModuleBasic{},
		ibcclients.AppModuleBasic{},
	)
}
