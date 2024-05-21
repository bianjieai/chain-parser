package farm

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocTxMsgUnstake struct {
	PoolId string      `bson:"pool_id"`
	Amount models.Coin `bson:"amount"`
	Sender string      `bson:"sender"`
}

func (m *DocTxMsgUnstake) GetType() string {
	return MsgTypeUnstake
}

func (m *DocTxMsgUnstake) BuildMsg(v interface{}) {
	msg := v.(*MsgUnstake)
	m.PoolId = msg.PoolId
	m.Amount = models.BuildDocCoin(msg.Amount)
	m.Sender = msg.Sender
}

func (m *DocTxMsgUnstake) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUnstake)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
