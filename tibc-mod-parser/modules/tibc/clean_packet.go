package tibc

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/tibc-mod-parser/modules"
)

type DocMsgCleanPacket struct {
	CleanPacket CleanPacket `bson:"clean_packet"`
	Signer      string      `bson:"signer"`
}

func (m *DocMsgCleanPacket) GetType() string {
	return MsgTypeTIBCCleanPacket
}

func (m *DocMsgCleanPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgCleanPacket)
	m.Signer = msg.Signer
	m.CleanPacket = loadCleanPacket(msg.CleanPacket)
}

func (m *DocMsgCleanPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgCleanPacket)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
