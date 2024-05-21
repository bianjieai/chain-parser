package staking

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

// MsgBeginUnbonding - struct for unbonding transactions
type DocTxMsgBeginUnbonding struct {
	DelegatorAddress string      `bson:"delegator_address"`
	ValidatorAddress string      `bson:"validator_address"`
	Amount           models.Coin `bson:"amount"`
}

func (m *DocTxMsgBeginUnbonding) GetType() string {
	return MsgTypeStakeBeginUnbonding
}

func (m *DocTxMsgBeginUnbonding) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgStakeBeginUnbonding)
	m.ValidatorAddress = msg.ValidatorAddress
	m.DelegatorAddress = msg.DelegatorAddress
	m.Amount = models.BuildDocCoin(msg.Amount)
}
func (m *DocTxMsgBeginUnbonding) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgStakeBeginUnbonding)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
