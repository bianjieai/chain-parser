package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgRedeemTokensForShares struct {
	DelegatorAddress string `bson:"delegator_address"`
	Amount           Coin   `bson:"amount"`
}

func (m *DocTxMsgRedeemTokensForShares) GetType() string {
	return MsgTypeRedeemTokensForShares
}

func (m *DocTxMsgRedeemTokensForShares) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgRedeemTokensForShares)
	m.DelegatorAddress = msg.DelegatorAddress
	m.Amount = Coin(models.BuildDocCoin(msg.Amount))
}
func (m *DocTxMsgRedeemTokensForShares) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string
	msg := v.(*MsgRedeemTokensForShares)
	addrs = append(addrs, msg.DelegatorAddress)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
