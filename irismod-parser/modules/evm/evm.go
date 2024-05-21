package evm

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
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
