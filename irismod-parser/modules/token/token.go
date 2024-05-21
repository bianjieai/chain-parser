package token

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
	v1 "gitlab.bianjie.ai/chain-parser/irismod-parser/modules/token/v1"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/token/v1/erc20"
	"gitlab.bianjie.ai/chain-parser/irismod-parser/modules/token/v1beta1"
)

type TokenClient struct {
}

func NewClient() TokenClient {
	return TokenClient{}
}
func (token TokenClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch msg := v.(type) {
	case *MsgMintToken:
		docMsg := v1beta1.DocMsgMintToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditToken:
		docMsg := v1beta1.DocMsgEditToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueToken:
		docMsg := v1beta1.DocMsgIssueToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwner:
		docMsg := v1beta1.DocMsgTransferTokenOwner{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnToken:
		docMsg := v1beta1.DocMsgBurnToken{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgMintTokenV1:
		docMsg := v1.DocMsgMintTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgEditTokenV1:
		docMsg := v1.DocMsgEditTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgIssueTokenV1:
		docMsg := v1.DocMsgIssueTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgTransferTokenOwnerV1:
		docMsg := v1.DocMsgTransferTokenOwnerV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgBurnTokenV1:
		docMsg := v1.DocMsgBurnTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgSwapFeeTokenV1:
		docMsg := v1.DocMsgSwapFeeTokenV1{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgSwapFromERC20:
		docMsg := erc20.DocMsgSwapFromERC20{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	case *MsgSwapToERC20:
		docMsg := erc20.DocMsgSwapToERC20{}
		msgDocInfo = docMsg.HandleTxMsg(msg)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
