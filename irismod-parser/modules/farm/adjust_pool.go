package farm

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocTxMsgAdjustPool struct {
	PoolId           string        `bson:"pool_id"`
	AdditionalReward []models.Coin `bson:"additional_reward"`
	RewardPerBlock   []models.Coin `bson:"reward_per_block"`
	Creator          string        `bson:"creator"`
}

func (m *DocTxMsgAdjustPool) GetType() string {
	return MsgTypeAdjustPool
}

func (m *DocTxMsgAdjustPool) BuildMsg(v interface{}) {
	msg := v.(*MsgAdjustPool)
	m.PoolId = msg.PoolId
	m.RewardPerBlock = models.BuildDocCoins(msg.RewardPerBlock)
	m.AdditionalReward = models.BuildDocCoins(msg.AdditionalReward)
	m.Creator = msg.Creator

}

func (m *DocTxMsgAdjustPool) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgAdjustPool)
	addrs = append(addrs, msg.Creator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
