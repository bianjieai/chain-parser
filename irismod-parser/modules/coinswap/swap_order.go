package coinswap

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type DocTxMsgSwapOrder struct {
	Input      Input  `bson:"input"`        // the amount the sender is trading
	Output     Output `bson:"output"`       // the amount the sender is receiving
	Deadline   int64  `bson:"deadline"`     // deadline for the transaction to still be considered valid
	IsBuyOrder bool   `bson:"is_buy_order"` // boolean indicating whether the order should be treated as a buy or sell
}

type Input struct {
	Address string      `bson:"address"`
	Coin    models.Coin `bson:"coin"`
}

type Output struct {
	Address string      `bson:"address"`
	Coin    models.Coin `bson:"coin"`
}

func (doctx *DocTxMsgSwapOrder) GetType() string {
	return MsgTypeSwapOrder
}

func (doctx *DocTxMsgSwapOrder) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSwapOrder)
	doctx.Deadline = msg.Deadline
	doctx.IsBuyOrder = msg.IsBuyOrder
	doctx.Input = Input{
		Address: msg.Input.Address,
		Coin:    models.BuildDocCoin(msg.Input.Coin),
	}
	doctx.Output = Output{
		Address: msg.Output.Address,
		Coin:    models.BuildDocCoin(msg.Output.Coin),
	}
}
func (m *DocTxMsgSwapOrder) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgSwapOrder)
	addrs = append(addrs, msg.Output.Address, msg.Input.Address)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
