package slashing

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
)

type DocTxMsgUnjailValidator struct {
	Id       string `bson:"id" json:"id"`
	Operator string `bson:"operator" json:"operator"`
}

func (m *DocTxMsgUnjailValidator) GetType() string {
	return MsgTypeUnjailValidator
}

func (m *DocTxMsgUnjailValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUnjailValidator)
	m.Id = msg.Id
	m.Operator = msg.Operator
}

func (m *DocTxMsgUnjailValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUnjailValidator)
	addrs = append(addrs, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
