package distribution

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgWithdrawTokenizeShareRecordReward struct {
	OwnerAddress string `bson:"owner_address"`
	RecordId     int64  `bson:"record_id"`
}

func (m *DocTxMsgWithdrawTokenizeShareRecordReward) GetType() string {
	return MsgTypeWithdrawTokenizeShareRecordReward
}

func (m *DocTxMsgWithdrawTokenizeShareRecordReward) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawTokenizeShareRecordReward)
	m.RecordId = int64(msg.RecordId)
	m.OwnerAddress = msg.OwnerAddress
}
func (m *DocTxMsgWithdrawTokenizeShareRecordReward) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgWithdrawTokenizeShareRecordReward)
	addrs = append(addrs, msg.OwnerAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
