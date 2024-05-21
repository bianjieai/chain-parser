package bank

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type BankClient struct {
}

func NewClient() BankClient {
	return BankClient{}
}

func (bank BankClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgSend:
		docMsg := DocMsgSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgMultiSend:
		docMsg := DocMsgMultiSend{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
