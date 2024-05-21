package irismod_parser

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/codec"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/coinswap"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/evm"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/farm"
	_ "gitlab.bianjie.ai/chain-parser/irismod-parser/modules/feemarket"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/htlc"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/mt"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/nft"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/oracle"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/random"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/record"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/service"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/token"
)

type MsgClient struct {
	Coinswap Client
	Farm     Client
	Htlc     Client
	Mt       Client
	Nft      Client
	Oracle   Client
	Random   Client
	Record   Client
	Service  Client
	Token    Client
	Evm      Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Coinswap: coinswap.NewClient(),
		Farm:     farm.NewClient(),
		Htlc:     htlc.NewClient(),
		Mt:       mt.NewClient(),
		Nft:      nft.NewClient(),
		Oracle:   oracle.NewClient(),
		Random:   random.NewClient(),
		Record:   record.NewClient(),
		Service:  service.NewClient(),
		Token:    token.NewClient(),
		Evm:      evm.NewClient(),
	}
}
