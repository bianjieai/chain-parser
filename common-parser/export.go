package common_parser

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type Client interface {
	HandleTxMsg(v types.Msg) (MsgDocInfo, bool)
}
