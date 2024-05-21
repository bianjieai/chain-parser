package group

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgUpdateGroupAdmin struct {
		Admin    string `bson:"admin"`
		GroupId  int64  `bson:"group_id"`
		NewAdmin string `bson:"new_admin"`
	}
)

func (m *DocTxMsgUpdateGroupAdmin) GetType() string {
	return MsgTypeUpdateGroupAdmin
}

func (m *DocTxMsgUpdateGroupAdmin) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUpdateGroupAdmin)
	m.Admin = msg.Admin
	m.GroupId = int64(msg.GroupId)
	m.NewAdmin = msg.NewAdmin
}
func (m *DocTxMsgUpdateGroupAdmin) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgUpdateGroupAdmin)

	addrs = append(addrs, msg.Admin, msg.NewAdmin)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
