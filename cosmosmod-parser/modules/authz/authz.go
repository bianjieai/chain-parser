package authz

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type AuthzClient struct {
}

func NewClient() AuthzClient {
	return AuthzClient{}
}

func (authz AuthzClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgExec:
		docMsg := DocMsgExec{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgRevoke:
		docMsg := DocMsgRevoke{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgGrant:
		docMsg := DocMsgGrant{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
