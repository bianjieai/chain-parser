package sidechain

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/iritamod-parser/modules"
)

type DocMsgCreateBlockHeader struct {
	SpaceId uint64 `bson:"space_id"`
	Height  uint64 `bson:"height"`
	Header  string `bson:"header"`
	Sender  string `bson:"sender"`
}

func (m *DocMsgCreateBlockHeader) GetType() string {
	return MsgTypeCreateBlockHeader
}

func (m *DocMsgCreateBlockHeader) BuildMsg(v interface{}) {
	msg := v.(*MsgCreateBlockHeader)
	m.SpaceId = msg.SpaceId
	m.Header = msg.Header
	m.Height = msg.Height
	m.Sender = msg.Sender
}

func (m *DocMsgCreateBlockHeader) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCreateBlockHeader)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
