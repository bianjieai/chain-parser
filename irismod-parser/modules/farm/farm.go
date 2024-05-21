package farm

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type FarmClient struct {
}

func NewClient() FarmClient {
	return FarmClient{}
}

func (farm FarmClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgCreatePool:
		docMsg := DocTxMsgCreatePool{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgCreatePoolWithCommunityPool:
		docMsg := DocTxMsgCreatePoolWithCommunityPool{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgDestroyPool:
		docMsg := DocTxMsgDestroyPool{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgAdjustPool:
		docMsg := DocTxMsgAdjustPool{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgHarvest:
		docMsg := DocTxMsgHarvest{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStake:
		docMsg := DocTxMsgStake{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnstake:
		docMsg := DocTxMsgUnstake{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}

}
