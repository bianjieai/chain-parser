package irismod_parser

import (
	. "github.com/bianjieai/chain-parser/common-parser"
	"github.com/bianjieai/chain-parser/irismod-parser/codec"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/coinswap"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/evm"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/farm"
	_ "github.com/bianjieai/chain-parser/irismod-parser/modules/feemarket"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/htlc"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/mt"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/nft"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/oracle"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/random"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/record"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/service"
	"github.com/bianjieai/chain-parser/irismod-parser/modules/token"
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
