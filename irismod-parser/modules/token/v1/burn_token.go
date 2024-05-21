package v1

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocMsgBurnTokenV1 struct {
	Coin   models.Coin `bson:"coin"`
	Sender string      `bson:"sender"`
}

func (m *DocMsgBurnTokenV1) GetType() string {
	return MsgTypeBurnToken
}

func (m *DocMsgBurnTokenV1) BuildMsg(v interface{}) {
	msg := v.(*MsgBurnTokenV1)
	m.Coin = models.BuildDocCoin(msg.Coin)
	m.Sender = msg.Sender
}

func (m *DocMsgBurnTokenV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgBurnTokenV1)

	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
