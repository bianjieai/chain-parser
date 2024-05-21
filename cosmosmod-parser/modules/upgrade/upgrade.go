package upgrade

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
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
