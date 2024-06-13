package msgs

import (
	models "github.com/bianjieai/chain-parser/common-parser/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type (
	MsgDocInfo struct {
		DocTxMsg models.TxMsg
		Addrs    []string
		Signers  []string
		Msgs     []sdk.Msg
	}
	SdkMsg sdk.Msg
	Msg    models.Msg

	Coin models.Coin

	Coins []*Coin
)
