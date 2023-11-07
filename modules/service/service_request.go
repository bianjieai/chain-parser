package service

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type (
	DocMsgCallService struct {
		ServiceName       string       `bson:"service_name"`
		Providers         []string     `bson:"providers"`
		Consumer          string       `bson:"consumer"`
		Input             string       `bson:"input"`
		ServiceFeeCap     models.Coins `bson:"service_fee_cap"`
		Timeout           int64        `bson:"timeout"`
		Repeated          bool         `bson:"repeated"`
		RepeatedFrequency int64        `bson:"repeated_frequency"`
		RepeatedTotal     int64        `bson:"repeated_total"`
	}
)

func (m *DocMsgCallService) GetType() string {
	return MsgTypeCallService
}

func (m *DocMsgCallService) BuildMsg(msg interface{}) {
	v := msg.(*MsgCallService)

	m.ServiceName = v.ServiceName
	m.Providers = v.Providers
	m.Consumer = v.Consumer
	m.Input = v.Input
	m.ServiceFeeCap = models.BuildDocCoins(v.ServiceFeeCap)
	m.Timeout = v.Timeout
	//m.Input = hex.EncodeToString(v.Input)
	m.Repeated = v.Repeated
	m.RepeatedFrequency = int64(v.RepeatedFrequency)
	m.RepeatedTotal = v.RepeatedTotal
}

func (m *DocMsgCallService) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCallService)
	addrs = append(addrs, msg.Providers...)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
