package params

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
)

type ParamsClient struct {
}

func NewClient() ParamsClient {
	return ParamsClient{}
}

func (params ParamsClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)

	return msgDocInfo, false
}
