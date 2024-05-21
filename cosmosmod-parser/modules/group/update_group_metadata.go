package group

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupMetadata struct {
		Admin    string `bson:"admin"`
		GroupId  int64  `bson:"group_id"`
		Metadata string `bson:"metadata"`
	}
)

func (m *DocTxMsgUpdateGroupMetadata) GetType() string {
	return MsgTypeUpdateGroupMetadata
}

func (m *DocTxMsgUpdateGroupMetadata) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupMetadata)
	m.Admin = msg.Admin
	m.GroupId = int64(msg.GroupId)
	m.Metadata = msg.Metadata
}

func (m *DocTxMsgUpdateGroupMetadata) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupMetadata)

	addrs = append(addrs, msg.Admin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
