package crisis

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CrisisClient struct {
}

func NewClient() CrisisClient {
	return CrisisClient{}
}

func (crisis CrisisClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgVerifyInvariant:
		docMsg := DocMsgVerifyInvariant{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
