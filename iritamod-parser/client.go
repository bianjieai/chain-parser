package iritamod_parser

import (
	common "github.com/bianjieai/chain-parser/common-parser"
	"github.com/bianjieai/chain-parser/iritamod-parser/codec"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/identity"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/node"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/params"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/perm"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/sidechain"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/slashing"
	"github.com/bianjieai/chain-parser/iritamod-parser/modules/upgrade"
)

type MsgClient struct {
	Params    common.Client
	Slashing  common.Client
	Upgrade   common.Client
	Identity  common.Client
	Perm      common.Client
	Node      common.Client
	SideChain common.Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Params:    params.NewClient(),
		Slashing:  slashing.NewClient(),
		Upgrade:   upgrade.NewClient(),
		Identity:  identity.NewClient(),
		Perm:      perm.NewClient(),
		Node:      node.NewClient(),
		SideChain: sidechain.NewClient(),
	}
}
