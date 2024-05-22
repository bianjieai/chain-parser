package sidechain

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
)

type DocMsgTransferSpace struct {
	SpaceId   uint64 `bson:"space_id"`
	Recipient string `bson:"recipient"`
	Sender    string `bson:"sender"`
}

func (m *DocMsgTransferSpace) GetType() string {
	return MsgTypeTransferSpace
}

func (m *DocMsgTransferSpace) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferSpace)
	m.SpaceId = msg.SpaceId
	m.Recipient = msg.Recipient
	m.Sender = msg.Sender
}

func (m *DocMsgTransferSpace) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransferSpace)
	addrs = append(addrs, msg.Sender)
	addrs = append(addrs, msg.Recipient)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
