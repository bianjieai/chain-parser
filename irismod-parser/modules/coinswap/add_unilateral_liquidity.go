package coinswap

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocTxAddUnilateralLiquidity struct {
	CounterpartyDenom string      `bson:"counterparty_denom"`
	ExactToken        models.Coin `bson:"exact_token"`
	MinLiquidity      string      `bson:"min_liquidity"`
	Deadline          int64       `bson:"deadline"`
	Sender            string      `bson:"sender"`
}

func (m *DocTxAddUnilateralLiquidity) GetType() string {
	return MsgTypeAddUnilateralLiquidity
}

func (m *DocTxAddUnilateralLiquidity) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgAddUnilateralLiquidity)
	m.CounterpartyDenom = msg.CounterpartyDenom
	m.ExactToken = models.BuildDocCoin(msg.ExactToken)
	m.MinLiquidity = msg.MinLiquidity.String()
	m.Deadline = msg.Deadline
	m.Sender = msg.Sender
}

func (m *DocTxAddUnilateralLiquidity) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgAddUnilateralLiquidity)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
