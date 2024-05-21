package cosmosmod_parser

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/codec"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/auth"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/authz"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/bank"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/crisis"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/distribution"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/evidence"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/feegrant"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/gov"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/group"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/ibc"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/params"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/slashing"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/staking"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/upgrade"
)

type MsgClient struct {
	Auth         Client
	Bank         Client
	Crisis       Client
	Distribution Client
	Evidence     Client
	Feegrant     Client
	Gov          Client
	Params       Client
	Slashing     Client
	Staking      Client
	Upgrade      Client
	Ibc          Client
	Authz        Client
	Group        Client
}

func NewMsgClient() MsgClient {
	codec.MakeEncodingConfig()
	return MsgClient{
		Auth:         auth.NewClient(),
		Bank:         bank.NewClient(),
		Crisis:       crisis.NewClient(),
		Distribution: distribution.NewClient(),
		Evidence:     evidence.NewClient(),
		Feegrant:     feegrant.NewClient(),
		Gov:          gov.NewClient(),
		Params:       params.NewClient(),
		Slashing:     slashing.NewClient(),
		Upgrade:      upgrade.NewClient(),
		Staking:      staking.NewClient(),
		Ibc:          ibc.NewClient(),
		Authz:        authz.NewClient(),
		Group:        group.NewClient(),
	}
}
