package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgTokenizeShares struct {
	DelegatorAddress    string `bson:"delegator_address"`
	ValidatorAddress    string `bson:"validator_address"`
	Amount              Coin   `bson:"amount"`
	TokenizedShareOwner string `bson:"tokenized_share_owner"`
}

func (m *DocTxMsgTokenizeShares) GetType() string {
	return MsgTypeTokenizeShares
}

func (m *DocTxMsgTokenizeShares) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgTokenizeShares)
	m.DelegatorAddress = msg.DelegatorAddress
	m.ValidatorAddress = msg.ValidatorAddress
	m.TokenizedShareOwner = msg.TokenizedShareOwner
	m.Amount = Coin(models.BuildDocCoin(msg.Amount))
}
func (m *DocTxMsgTokenizeShares) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgTokenizeShares)
	addrs = append(addrs, msg.DelegatorAddress, msg.ValidatorAddress, msg.TokenizedShareOwner)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
