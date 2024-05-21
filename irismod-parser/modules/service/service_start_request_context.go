package service

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
	"strings"
)

type (
	DocMsgStartRequestContext struct {
		RequestContextID string `bson:"request_context_id" yaml:"request_context_id"`
		Consumer         string `bson:"consumer" yaml:"consumer"`
	}
)

func (m *DocMsgStartRequestContext) GetType() string {
	return MsgTypeStartRequestContext
}

func (m *DocMsgStartRequestContext) BuildMsg(v interface{}) {
	msg := v.(*MsgStartRequestContext)

	m.RequestContextID = strings.ToUpper(msg.RequestContextId)
	m.Consumer = msg.Consumer
}

func (m *DocMsgStartRequestContext) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgStartRequestContext)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
