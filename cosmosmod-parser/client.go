package cosmosmod_parser

import (
	. "github.com/bianjieai/chain-parser/common-parser"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/codec"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/auth"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/bank"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/crisis"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/distribution"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/evidence"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/feegrant"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/gov"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/params"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/slashing"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/staking"
	"github.com/bianjieai/chain-parser/cosmosmod-parser/modules/upgrade"
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
	}
}
