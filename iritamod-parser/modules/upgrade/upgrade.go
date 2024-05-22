package upgrade

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type UpgradeClient struct {
}

func NewClient() UpgradeClient {
	return UpgradeClient{}
}

func (upgrade UpgradeClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)

	return msgDocInfo, false
}
