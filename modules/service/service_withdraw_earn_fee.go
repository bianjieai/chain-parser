package service

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgWithdrawEarnedFees struct {
		Owner    string
		Provider string
	}
)

func (m *DocMsgWithdrawEarnedFees) GetType() string {
	return MsgTypeWithdrawEarnedFees
}

func (m *DocMsgWithdrawEarnedFees) BuildMsg(v interface{}) {
	msg := v.(*MsgWithdrawEarnedFees)

	m.Owner = msg.Owner
	m.Provider = msg.Provider
}

func (m *DocMsgWithdrawEarnedFees) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgWithdrawEarnedFees)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
