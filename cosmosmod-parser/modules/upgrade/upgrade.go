package upgrade

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
)

type UpgradeClient struct {
}

func NewClient() UpgradeClient {
	return UpgradeClient{}
}

func (upgrade UpgradeClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	return msgDocInfo, false
}
