package service

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgSetWithdrawAddress struct {
		Owner           string `bson:"owner" yaml:"owner"`
		WithdrawAddress string `bson:"withdraw_address" yaml:"withdraw_address"`
	}
)

func (m *DocMsgSetWithdrawAddress) GetType() string {
	return MsgTypeServiceSetWithdrawAddress
}

func (m *DocMsgSetWithdrawAddress) BuildMsg(v interface{}) {
	msg := v.(*MsgSetWithdrawAddress)

	m.Owner = msg.Owner
	m.WithdrawAddress = msg.WithdrawAddress
}

func (m *DocMsgSetWithdrawAddress) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgSetWithdrawAddress)
	addrs = append(addrs, msg.Owner, msg.WithdrawAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
