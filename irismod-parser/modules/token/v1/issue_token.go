package v1

import (
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocMsgIssueTokenV1 struct {
	Symbol        string `bson:"symbol"`
	Name          string `bson:"name"`
	Scale         uint32 `bson:"scale"`
	MinUnit       string `bson:"min_unit"`
	InitialSupply string `bson:"initial_supply"`
	MaxSupply     string `bson:"max_supply"`
	Mintable      bool   `bson:"mintable"`
	Owner         string `bson:"owner"`
}

func (m *DocMsgIssueTokenV1) GetType() string {
	return MsgTypeIssueToken
}

func (m *DocMsgIssueTokenV1) BuildMsg(v interface{}) {
	msg := v.(*MsgIssueTokenV1)

	m.Symbol = msg.Symbol
	m.Name = msg.Name
	m.Scale = msg.Scale
	m.MinUnit = msg.MinUnit
	m.InitialSupply = fmt.Sprint(msg.InitialSupply)
	m.Owner = msg.Owner
	m.MaxSupply = fmt.Sprint(msg.MaxSupply)
	m.Mintable = msg.Mintable
}

func (m *DocMsgIssueTokenV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgIssueTokenV1)
	addrs = append(addrs, msg.Owner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
