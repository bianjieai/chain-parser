package random

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
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
