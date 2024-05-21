package evidence

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
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
