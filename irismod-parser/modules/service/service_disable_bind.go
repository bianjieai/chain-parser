package service

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgDisableServiceBinding struct {
		ServiceName string `bson:"service_name" yaml:"service_name"`
		Provider    string `bson:"provider" yaml:"provider"`
		Owner       string `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgDisableServiceBinding) GetType() string {
	return MsgTypeDisableServiceBinding
}

func (m *DocMsgDisableServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgDisableServiceBinding)

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Owner = msg.Owner
}

func (m *DocMsgDisableServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDisableServiceBinding)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
