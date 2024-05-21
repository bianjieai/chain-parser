package group

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgCreateGroupPolicy struct {
		Admin          string      `bson:"admin"`
		GroupId        int64       `bson:"group_id"`
		Metadata       string      `bson:"metadata"`
		DecisionPolicy interface{} `bson:"decision_policy"`
	}
)

func (m *DocTxMsgCreateGroupPolicy) GetType() string {
	return MsgTypeCreateGroupPolicy
}

func (m *DocTxMsgCreateGroupPolicy) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgCreateGroupPolicy)
	m.Admin = msg.Admin
	m.GroupId = int64(msg.GroupId)
	m.Metadata = msg.Metadata
	m.DecisionPolicy = msg.DecisionPolicy
}

func (m *DocTxMsgCreateGroupPolicy) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgCreateGroupPolicy)

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
