package v1

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocMsgTransferTokenOwnerV1 struct {
	SrcOwner string `bson:"src_owner"`
	DstOwner string `bson:"dst_owner"`
	Symbol   string `bson:"symbol"`
}

func (m *DocMsgTransferTokenOwnerV1) GetType() string {
	return MsgTypeTransferTokenOwner
}

func (m *DocMsgTransferTokenOwnerV1) BuildMsg(v interface{}) {
	msg := v.(*MsgTransferTokenOwnerV1)

	m.Symbol = msg.Symbol
	m.SrcOwner = msg.SrcOwner
	m.DstOwner = msg.DstOwner
}

func (m *DocMsgTransferTokenOwnerV1) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgTransferTokenOwnerV1)
	addrs = append(addrs, msg.SrcOwner, msg.DstOwner)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
