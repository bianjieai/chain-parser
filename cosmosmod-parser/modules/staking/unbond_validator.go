package staking

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgUnbondValidator struct {
	ValidatorAddress string `bson:"validator_address"`
}

func (m *DocTxMsgUnbondValidator) GetType() string {
	return MsgTypeUnbondValidator
}

func (m *DocTxMsgUnbondValidator) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgUnbondValidator)
	m.ValidatorAddress = msg.ValidatorAddress
}
func (m *DocTxMsgUnbondValidator) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgUnbondValidator)
	addrs = append(addrs, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
