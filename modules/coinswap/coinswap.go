package coinswap

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/common-parser/modules"
	. "github.com/kaifei-bianjie/irismod-parser/modules"
)

type CoinswapClient struct {
}

func NewClient() CoinswapClient {
	return CoinswapClient{}
}

func (coinswap CoinswapClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgAddLiquidity:
		docMsg := DocTxMsgAddLiquidity{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgRemoveLiquidity:
		docMsg := DocTxMsgRemoveLiquidity{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgSwapOrder:
		docMsg := DocTxMsgSwapOrder{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgAddUnilateralLiquidity:
		docMsg := DocTxAddUnilateralLiquidity{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgRemoveUnilateralLiquidity:
		docMsg := DocTxMsgRemoveUnilateralLiquidity{}
		return docMsg.HandleTxMsg(msg), true
	default:
		return MsgDocInfo{}, false
	}
}
