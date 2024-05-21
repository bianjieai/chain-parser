package distribution

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgWithdrawAllTokenizeShareRecordReward struct {
	OwnerAddress string `bson:"owner_address"`
}

func (m *DocTxMsgWithdrawAllTokenizeShareRecordReward) GetType() string {
	return MsgTypeWithdrawAllTokenizeShareRecordReward
}

func (m *DocTxMsgWithdrawAllTokenizeShareRecordReward) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawAllTokenizeShareRecordReward)
	m.OwnerAddress = msg.OwnerAddress
}
func (m *DocTxMsgWithdrawAllTokenizeShareRecordReward) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgWithdrawAllTokenizeShareRecordReward)
	addrs = append(addrs, msg.OwnerAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
