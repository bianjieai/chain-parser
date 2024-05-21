package feegrant

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DocTxMsgRevokeAllowance struct {
	Granter string `bson:"granter"`
	Grantee string `bson:"grantee"`
}

func (m *DocTxMsgRevokeAllowance) GetType() string {
	return MsgTypeRevokeAllowance
}

func (m *DocTxMsgRevokeAllowance) BuildMsg(v interface{}) {
	msg := v.(*MsgRevokeAllowance)

	m.Granter = msg.Granter
	m.Grantee = msg.Grantee
}

func (m *DocTxMsgRevokeAllowance) HandleTxMsg(v sdk.Msg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgRevokeAllowance)
	addrs = append(addrs, msg.Granter, msg.Grantee)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
