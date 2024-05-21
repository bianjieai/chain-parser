package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgValidatorBond struct {
	DelegatorAddress string `bson:"delegator_address"`
	ValidatorAddress string `bson:"validator_address"`
}

func (m *DocTxMsgValidatorBond) GetType() string {
	return MsgTypeValidatorBond
}

func (m *DocTxMsgValidatorBond) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgValidatorBond)
	m.DelegatorAddress = msg.DelegatorAddress
	m.ValidatorAddress = msg.ValidatorAddress
}
func (m *DocTxMsgValidatorBond) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgValidatorBond)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
