package erc20

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

// MsgSwapToERC20 defines an SDK message for SwapToERC20
type DocMsgSwapToERC20 struct {
	Amount   models.Coin `bson:"amount"`
	Sender   string      `bson:"sender"`
	Receiver string      `bson:"receiver"`
}

func (m *DocMsgSwapToERC20) GetType() string {
	return MsgTypeSwapToERC20
}

func (m *DocMsgSwapToERC20) BuildMsg(v interface{}) {
	msg := v.(*MsgSwapToERC20)

	m.Amount = models.BuildDocCoin(msg.Amount)
	m.Receiver = msg.Receiver
	m.Sender = msg.Sender
}

func (m *DocMsgSwapToERC20) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgSwapToERC20)
	addrs = append(addrs, msg.Sender, msg.Receiver)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
