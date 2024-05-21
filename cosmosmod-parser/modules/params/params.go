package params

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
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
