package service

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgEnableServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgEnableServiceBinding) GetType() string {
	return MsgTypeEnableServiceBinding
}

func (m *DocMsgEnableServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgEnableServiceBinding)
	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = models.BuildDocCoins(msg.Deposit)
	m.Owner = msg.Owner
}

func (m *DocMsgEnableServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgEnableServiceBinding)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
