package slashing

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type SlashingClient struct {
}

func NewClient() SlashingClient {
	return SlashingClient{}
}

func (slashing SlashingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgUnjailValidator:
		docMsg := DocTxMsgUnjailValidator{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
