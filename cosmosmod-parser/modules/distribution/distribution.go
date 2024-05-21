package distribution

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DistributionClient struct {
}

func NewClient() DistributionClient {
	return DistributionClient{}
}

func (distribution DistributionClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgStakeSetWithdrawAddress:
		docMsg := DocTxMsgSetWithdrawAddress{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawDelegatorReward:
		docMsg := DocTxMsgWithdrawDelegatorReward{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawValidatorCommission:
		docMsg := DocTxMsgWithdrawValidatorCommission{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgFundCommunityPool:
		docMsg := DocTxMsgFundCommunityPool{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawAllTokenizeShareRecordReward:
		docMsg := DocTxMsgWithdrawAllTokenizeShareRecordReward{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgWithdrawTokenizeShareRecordReward:
		docMsg := DocTxMsgWithdrawTokenizeShareRecordReward{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
