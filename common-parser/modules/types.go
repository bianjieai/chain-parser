package msgs

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin
)
