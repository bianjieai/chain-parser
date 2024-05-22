package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakingClient struct {
}

func NewClient() StakingClient {
	return StakingClient{}
}

func (staking StakingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgBeginRedelegate:
		docMsg := DocTxMsgBeginRedelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeBeginUnbonding:
		docMsg := DocTxMsgBeginUnbonding{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeDelegate:
		docMsg := DocTxMsgDelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeEdit:
		docMsg := DocMsgEditValidator{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeCreate:
		docMsg := DocTxMsgCreateValidator{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
