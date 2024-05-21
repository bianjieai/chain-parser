package group

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupPolicyDecisionPolicy struct {
		Admin              string      `bson:"admin"`
		GroupPolicyAddress string      `bson:"group_policy_address"`
		DecisionPolicy     interface{} `bson:"decision_policy"`
	}
)

func (m *DocTxMsgUpdateGroupPolicyDecisionPolicy) GetType() string {
	return MsgTypeUpdateGroupPolicyDecisionPolicy
}

func (m *DocTxMsgUpdateGroupPolicyDecisionPolicy) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupPolicyDecisionPolicy)
	m.Admin = msg.Admin
	m.GroupPolicyAddress = msg.GroupPolicyAddress
	m.DecisionPolicy = msg.DecisionPolicy
}

func (m *DocTxMsgUpdateGroupPolicyDecisionPolicy) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupPolicyDecisionPolicy)
	addrs = append(addrs, msg.Admin, msg.GroupPolicyAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
