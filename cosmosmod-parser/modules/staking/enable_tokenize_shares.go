package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgEnableTokenizeShares struct {
	DelegatorAddress string `bson:"delegator_address"`
}

func (m *DocTxMsgEnableTokenizeShares) GetType() string {
	return MsgTypeEnableTokenizeShares
}

func (m *DocTxMsgEnableTokenizeShares) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgEnableTokenizeShares)
	m.DelegatorAddress = msg.DelegatorAddress
}
func (m *DocTxMsgEnableTokenizeShares) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgEnableTokenizeShares)
	addrs = append(addrs, msg.DelegatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
