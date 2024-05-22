package oracle

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type OracleClient struct {
}

func NewClient() OracleClient {
	return OracleClient{}
}

func (oracle OracleClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgStartFeed:
		docMsg := DocMsgStartFeed{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgPauseFeed:
		docMsg := DocMsgPauseFeed{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgEditFeed:
		docMsg := DocMsgEditFeed{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgCreateFeed:
		docMsg := DocMsgCreateFeed{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}

}
