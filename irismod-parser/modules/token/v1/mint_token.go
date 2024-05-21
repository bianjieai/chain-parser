package v1

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocMsgMintTokenV1 struct {
	Coin  models.Coin `bson:"coin"`
	To    string      `bson:"to"`
	Owner string      `bson:"owner"`
}

func (m *DocMsgMintTokenV1) GetType() string {
	return MsgTypeMintToken
}

func (m *DocMsgMintTokenV1) BuildMsg(v interface{}) {
	msg := v.(*MsgMintTokenV1)

	m.Coin = models.BuildDocCoin(msg.Coin)
	m.To = msg.Receiver
	m.Owner = msg.Owner
}

func (m *DocMsgMintTokenV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgMintTokenV1)
	addrs = append(addrs, msg.Owner, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
