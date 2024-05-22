package perm

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
)

type DocMsgUnblockAccount struct {
	Address  string `bson:"address" json:"address"`
	Operator string `bson:"operator" json:"operator"`
}

func (m *DocMsgUnblockAccount) GetType() string {
	return DocTypeUnlockAccount
}

func (m *DocMsgUnblockAccount) BuildMsg(v interface{}) {
	msg := v.(*MsgUnblockAccount)
	m.Address = msg.Address
	m.Operator = msg.Operator
}

func (m *DocMsgUnblockAccount) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUnblockAccount)
	addrs = append(addrs, msg.Address, msg.Operator)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
