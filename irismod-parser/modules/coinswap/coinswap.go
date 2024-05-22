package coinswap

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/irismod-parser/modules"
	"github.com/cosmos/cosmos-sdk/types"
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
	default:
		return MsgDocInfo{}, false
	}
}
