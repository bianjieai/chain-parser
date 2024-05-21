package farm

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocTxMsgHarvest struct {
	PoolId string `bson:"pool_id"`
	Sender string `bson:"sender"`
}

func (m *DocTxMsgHarvest) GetType() string {
	return MsgTypeHarvest
}

func (m *DocTxMsgHarvest) BuildMsg(v interface{}) {
	msg := v.(*MsgHarvest)
	m.PoolId = msg.PoolId
	m.Sender = msg.Sender
}

func (m *DocTxMsgHarvest) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgHarvest)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
