package ibc

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocMsgTransfer struct {
	PacketId         string      `bson:"packet_id"`
	SourcePort       string      `bson:"source_port"`
	SourceChannel    string      `bson:"source_channel"`
	Token            models.Coin `bson:"token"`
	Sender           string      `bson:"sender"`
	Receiver         string      `bson:"receiver"`
	TimeoutHeight    Height      `bson:"timeout_height"`
	TimeoutTimestamp int64       `bson:"timeout_timestamp"`
	Memo             string      `bson:"memo"`
}

func (m *DocMsgTransfer) GetType() string {
	return MsgTypeIBCTransfer
}

func (m *DocMsgTransfer) BuildMsg(v interface{}) {
	msg := v.(*MsgTransfer)
	m.SourcePort = msg.SourcePort
	m.SourceChannel = msg.SourceChannel
	m.Sender = msg.Sender
	m.Receiver = msg.Receiver
	m.TimeoutTimestamp = int64(msg.TimeoutTimestamp)
	m.TimeoutHeight = loadHeight(msg.TimeoutHeight)
	m.Token = models.BuildDocCoin(msg.Token)
	m.Memo = msg.Memo
}

func (m *DocMsgTransfer) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransfer)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
