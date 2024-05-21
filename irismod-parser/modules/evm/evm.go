package evm

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

type EvmClient struct {
}

func NewClient() EvmClient {
	return EvmClient{}
}

func (evm EvmClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {

	switch msg := v.(type) {
	case *MsgEthereumTx:
		docMsg := DocMsgEthereumTx{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}

}
