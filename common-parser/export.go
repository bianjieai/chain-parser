package common_parser

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
