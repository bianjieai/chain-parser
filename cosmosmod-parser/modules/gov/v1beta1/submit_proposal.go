package v1beta1

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	cdc "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	commoncodec "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	models "gitlab.bianjie.ai/chain-parser/common-parser/types"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgSubmitProposal struct {
	Proposer       string        `bson:"proposer"`        //  Address of the proposer
	InitialDeposit []models.Coin `bson:"initial_deposit"` //  Initial deposit paid by sender. Must be strictly positive.
	Content        interface{}   `bson:"content"`
}

func (m *DocTxMsgSubmitProposal) GetType() string {
	return MsgTypeSubmitProposal
}

func (m *DocTxMsgSubmitProposal) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgSubmitProposal)

	marshalJSON, err := commoncodec.GetMarshaler().MarshalJSON(msg.Content)
	if err != nil {
		logrus.Errorf("DocTxMsgSubmitProposal commoncodec.GetMarshaler().MarshalJSON err:%v", err)
	}

	var contentInterface interface{}
	if err = json.Unmarshal(marshalJSON, &contentInterface); err != nil {
		logrus.Errorf("DocTxMsgSubmitProposal json.Unmarshal err:%v", err)
	}
	m.Content = contentInterface
	m.Proposer = msg.Proposer
	m.InitialDeposit = models.BuildDocCoins(msg.InitialDeposit)
}

func CovertContent(content GovContent) interface{} {
	switch content.ProposalType() {
	case ProposalTypeCancelSoftwareUpgrade:
		var data ContentCancelSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeSoftwareUpgrade:
		var data ContentSoftwareUpgradeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeCommunityPoolSpend:
		var data ContentCommunityPoolSpendProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeClientUpdate:
		var data ContentClientUpdateProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeText:
		var data ContentTextProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	case ProposalTypeParameterChange:
		var data ContentParameterChangeProposal
		utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(content), &data)
		return data
	}
	return content
}

func (m *DocTxMsgSubmitProposal) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var (
		addrs []string
		msg   MsgSubmitProposal
	)

	data, _ := cdc.GetMarshaler().MarshalJSON(v)
	cdc.GetMarshaler().UnmarshalJSON(data, &msg)

	content := msg.GetContent()
	if content != nil && ProposalTypeCommunityPoolSpend == content.ProposalType() {
		communityPoolSpend := CovertContent(content).(ContentCommunityPoolSpendProposal)
		addrs = append(addrs, communityPoolSpend.Recipient)
	}
	addrs = append(addrs, msg.Proposer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
