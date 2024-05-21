package v1

import (
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocMsgEditTokenV1 struct {
	Symbol    string `bson:"symbol"`
	Name      string `bson:"name"`
	MaxSupply string `bson:"max_supply"`
	Mintable  bool   `bson:"mintable"`
	Owner     string `bson:"owner"`
}

func (m *DocMsgEditTokenV1) GetType() string {
	return MsgTypeEditToken
}

func (m *DocMsgEditTokenV1) BuildMsg(v interface{}) {
	msg := v.(*MsgEditTokenV1)

	m.Symbol = msg.Symbol
	m.Owner = msg.Owner
	m.Name = msg.Name
	m.MaxSupply = fmt.Sprint(msg.MaxSupply)
	m.Mintable = msg.Mintable.ToBool()
}

func (m *DocMsgEditTokenV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgEditTokenV1)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
