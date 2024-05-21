package evidence

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type EvidenceClient struct {
}

func NewClient() EvidenceClient {
	return EvidenceClient{}
}

func (evidence EvidenceClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgSubmitEvidence:
		docMsg := DocMsgSubmitEvidence{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
