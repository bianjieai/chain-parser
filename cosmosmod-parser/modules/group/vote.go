package group

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgGroupVote struct {
		ProposalId int64  `bson:"proposal_id"`
		Voter      string `bson:"voter"`
		Option     int32  `bson:"option"`
		Metadata   string `bson:"metadata"`
		Exec       int32  `bson:"exec"`
	}
)

func (m *DocTxMsgGroupVote) GetType() string {
	return MsgTypeGroupVote
}

func (m *DocTxMsgGroupVote) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgGroupVote)
	m.ProposalId = int64(msg.ProposalId)
	m.Voter = msg.Voter
	m.Option = int32(msg.Option)
	m.Metadata = msg.Metadata
	m.Exec = int32(msg.Exec)
}

func (m *DocTxMsgGroupVote) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgGroupVote)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
