package group

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type (
	DocTxMsgWithdrawProposal struct {
		ProposalId int64  `bson:"proposal_id"`
		Address    string `bson:"address"`
	}
)

func (m *DocTxMsgWithdrawProposal) GetType() string {
	return MsgTypeWithdrawProposal
}

func (m *DocTxMsgWithdrawProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgWithdrawProposal)
	m.ProposalId = int64(msg.ProposalId)
	m.Address = msg.Address
}

func (m *DocTxMsgWithdrawProposal) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgWithdrawProposal)
	addrs = append(addrs, msg.Address)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
