package coinswap

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

type DocTxMsgAddLiquidity struct {
	MaxToken     models.Coin `bson:"max_token"`      // coin to be deposited as liquidity with an upper bound for its amount
	ExactIrisAmt string      `bson:"exact_iris_amt"` // exact amount of native asset being add to the liquidity pool
	MinLiquidity string      `bson:"min_liquidity"`  // lower bound UNI sender is willing to accept for deposited coins
	Deadline     int64       `bson:"deadline"`
	Sender       string      `bson:"sender"`
}

func (doctx *DocTxMsgAddLiquidity) GetType() string {
	return MsgTypeAddLiquidity
}

func (doctx *DocTxMsgAddLiquidity) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgAddLiquidity)
	doctx.Sender = msg.Sender
	doctx.MinLiquidity = msg.MinLiquidity.String()
	doctx.ExactIrisAmt = msg.ExactStandardAmt.String()
	doctx.Deadline = msg.Deadline
	doctx.MaxToken = models.BuildDocCoin(msg.MaxToken)
}

func (m *DocTxMsgAddLiquidity) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgAddLiquidity)
	addrs = append(addrs, msg.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
