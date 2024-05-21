package erc20

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	models "github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
)

// MsgSwapFromERC20 defines an SDK message for SwapFromERC20
type DocMsgSwapFromERC20 struct {
	WantedAmount models.Coin `bson:"wanted_amount"`
	Sender       string      `bson:"sender"`
	Receiver     string      `bson:"receiver"`
}

func (m *DocMsgSwapFromERC20) GetType() string {
	return MsgTypeSwapFromERC20
}

func (m *DocMsgSwapFromERC20) BuildMsg(v interface{}) {
	msg := v.(*MsgSwapFromERC20)

	m.WantedAmount = models.BuildDocCoin(msg.WantedAmount)
	m.Receiver = msg.Receiver
	m.Sender = msg.Sender
}

func (m *DocMsgSwapFromERC20) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgSwapFromERC20)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
