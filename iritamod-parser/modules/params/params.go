package params

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type ParamsClient struct {
}

func NewClient() ParamsClient {
	return ParamsClient{}
}

func (params ParamsClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	return msgDocInfo, false
}
