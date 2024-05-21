package auth

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type AuthClient struct {
}

func NewClient() AuthClient {
	return AuthClient{}
}

func (auth AuthClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	return msgDocInfo, false
}
