package sidechain

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
)

type DocMsgCreateSpace struct {
	Name   string `bson:"name"`
	Uri    string `bson:"uri"`
	Sender string `bson:"sender"`
}

func (m *DocMsgCreateSpace) GetType() string {
	return MsgTypeCreateSpace
}

func (m *DocMsgCreateSpace) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateSpace)
	m.Name = msg.Name
	m.Uri = msg.Uri
	m.Sender = msg.Sender
}

func (m *DocMsgCreateSpace) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCreateSpace)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
