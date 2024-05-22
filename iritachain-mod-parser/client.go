package iritachain_mod_parser

import (
	. "github.com/bianjieai/chain-parser/common-parser"
	"github.com/bianjieai/chain-parser/iritachain-mod-parser/codec"
	"github.com/bianjieai/chain-parser/iritachain-mod-parser/modules/evm"
)

type MsgClient struct {
	Evm Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Evm: evm.NewClient(),
	}
}
