package service

import (
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type (
	DocMsgBindService struct {
		ServiceName string `bson:"service_name"`
		Provider    string `bson:"provider"`
		Deposit     Coins  `bson:"deposit"`
		Pricing     string `bson:"pricing"`
		QoS         int64  `bson:"qos"`
		Owner       string `bson:"owner"`
		Options     string `bson:"options"`
	}
)

func (m *DocMsgBindService) GetType() string {
	return MsgTypeBindService
}

func (m *DocMsgBindService) BuildMsg(v interface{}) {
	msg := v.(*MsgBindService)

	var coins Coins
	for _, one := range msg.Deposit {
		coins = append(coins, &Coin{Denom: one.Denom, Amount: one.Amount.String()})
	}
	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = coins
	m.Pricing = msg.Pricing
	m.QoS = int64(msg.QoS)
	m.Owner = msg.Owner
	m.Options = msg.Options
}

func (m *DocMsgBindService) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgBindService)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
