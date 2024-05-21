package v1

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

// MsgVote
type DocTxMsgVoteWeightedV1 struct {
	ProposalID int64                `bson:"proposal_id"` // ID of the proposal
	Voter      string               `bson:"voter"`       //  address of the voter
	Options    []WeightedVoteOption `bson:"options"`     //  option from OptionSet chosen by the voter
	Metadata   string               `bson:"metadata"`
}

type WeightedVoteOption struct {
	Option int32  `bson:"option"`
	Weight string `bson:"weight"`
}

func (m *DocTxMsgVoteWeightedV1) GetType() string {
	return MsgTypeVoteWeighted
}

func (m *DocTxMsgVoteWeightedV1) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgVoteWeightedV1)
	m.Voter = msg.Voter

	for _, v := range msg.Options {
		m.Options = append(m.Options, WeightedVoteOption{
			Option: int32(v.Option),
			Weight: v.Weight,
		})
	}
	m.ProposalID = int64(msg.ProposalId)
	m.Metadata = msg.Metadata
}

func (m *DocTxMsgVoteWeightedV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgVoteWeightedV1)
	addrs = append(addrs, msg.Voter)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
