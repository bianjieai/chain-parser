package group

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgGroupExec struct {
		ProposalId int64  `bson:"proposal_id"`
		Executor   string `bson:"executor"`
	}
)

func (m *DocTxMsgGroupExec) GetType() string {
	return MsgTypeGroupExec
}

func (m *DocTxMsgGroupExec) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgGroupExec)
	m.ProposalId = int64(msg.ProposalId)
	m.Executor = msg.Executor
}

func (m *DocTxMsgGroupExec) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgGroupExec)
	addrs = append(addrs, msg.Executor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
