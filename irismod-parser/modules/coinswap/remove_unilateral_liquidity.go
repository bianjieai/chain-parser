package coinswap

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocTxMsgRemoveUnilateralLiquidity struct {
	CounterpartyDenom string      `bson:"counterparty_denom"`
	MinToken          models.Coin `bson:"min_token"`
	ExactLiquidity    string      `bson:"exact_liquidity"`
	Deadline          int64       `bson:"deadline"`
	Sender            string      `bson:"sender"`
}

func (m *DocTxMsgRemoveUnilateralLiquidity) GetType() string {
	return MsgTypeRemoveUnilateralLiquidity
}

func (m *DocTxMsgRemoveUnilateralLiquidity) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgRemoveUnilateralLiquidity)
	m.CounterpartyDenom = msg.CounterpartyDenom
	m.MinToken = models.BuildDocCoin(msg.MinToken)
	m.ExactLiquidity = msg.ExactLiquidity.String()
	m.Deadline = msg.Deadline
	m.Sender = msg.Sender
}

func (m *DocTxMsgRemoveUnilateralLiquidity) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgRemoveUnilateralLiquidity)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
