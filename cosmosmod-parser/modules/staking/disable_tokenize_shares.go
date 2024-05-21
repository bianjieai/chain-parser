package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgDisableTokenizeShares struct {
	DelegatorAddress string `bson:"delegator_address"`
}

func (m *DocTxMsgDisableTokenizeShares) GetType() string {
	return MsgTypeDisableTokenizeShares
}

func (m *DocTxMsgDisableTokenizeShares) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgDisableTokenizeShares)
	m.DelegatorAddress = msg.DelegatorAddress
}
func (m *DocTxMsgDisableTokenizeShares) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgDisableTokenizeShares)
	addrs = append(addrs, msg.DelegatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
