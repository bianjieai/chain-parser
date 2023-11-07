package service

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgUpdateServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Pricing     string       `bson:"pricing" yaml:"pricing"`
		QoS         int64        `bson:"qos" yaml:"qos"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgUpdateServiceBinding) GetType() string {
	return MsgTypeUpdateServiceBinding
}

func (m *DocMsgUpdateServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateServiceBinding)

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = models.BuildDocCoins(msg.Deposit)
	m.Pricing = msg.Pricing
	m.QoS = int64(msg.QoS)
	m.Owner = msg.Owner
}

func (m *DocMsgUpdateServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgUpdateServiceBinding)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
