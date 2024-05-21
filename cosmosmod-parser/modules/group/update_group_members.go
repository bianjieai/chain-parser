package group

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupMembers struct {
		Admin         string          `bson:"admin"`
		GroupId       int64           `bson:"group_id"`
		MemberUpdates []MemberRequest `bson:"member_updates"`
	}
)

func (m *DocTxMsgUpdateGroupMembers) GetType() string {
	return MsgTypeUpdateGroupMembers
}

func (m *DocTxMsgUpdateGroupMembers) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupMembers)
	m.Admin = msg.Admin
	m.GroupId = int64(msg.GroupId)
	memberUpdates := make([]MemberRequest, 0, len(msg.MemberUpdates))
	for _, member := range msg.MemberUpdates {
		memberUpdates = append(memberUpdates, MemberRequest{
			member.Address, member.Weight, member.Metadata,
		})
	}
	m.MemberUpdates = memberUpdates

}
func (m *DocTxMsgUpdateGroupMembers) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupMembers)
	for _, member := range msg.MemberUpdates {
		addrs = append(addrs, member.Address)
	}

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
