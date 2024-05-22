package tibc_mod_parser

import (
	. "github.com/bianjieai/chain-parser/common-parser"
	"github.com/bianjieai/chain-parser/tibc-mod-parser/codec"
	"github.com/bianjieai/chain-parser/tibc-mod-parser/modules/tibc"
)

type MsgClient struct {
	Tibc Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Tibc: tibc.NewClient(),
	}
}
