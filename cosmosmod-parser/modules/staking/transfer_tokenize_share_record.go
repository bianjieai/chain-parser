package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgTransferTokenizeShareRecord struct {
	TokenizeShareRecordId int64  `bson:"tokenize_share_record_id"`
	Sender                string `bson:"sender"`
	NewOwner              string `bson:"new_owner"`
}

func (m *DocTxMsgTransferTokenizeShareRecord) GetType() string {
	return MsgTypeTransferTokenizeShareRecord
}

func (m *DocTxMsgTransferTokenizeShareRecord) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgTransferTokenizeShareRecord)
	m.TokenizeShareRecordId = int64(msg.TokenizeShareRecordId)
	m.Sender = msg.Sender
	m.NewOwner = msg.NewOwner
}
func (m *DocTxMsgTransferTokenizeShareRecord) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgTransferTokenizeShareRecord)
	addrs = append(addrs, msg.Sender, msg.NewOwner)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
