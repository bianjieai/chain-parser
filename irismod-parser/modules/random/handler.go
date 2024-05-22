package random

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type RandomClient struct {
}

func NewClient() RandomClient {
	return RandomClient{}
}

func (random RandomClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgRequestRandom:
		txMsg := DocTxMsgRequestRand{}
		return txMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}
}
