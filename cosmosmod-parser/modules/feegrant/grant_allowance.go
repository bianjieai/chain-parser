package feegrant

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

type DocTxMsgGrantAllowance struct {
	Granter   string      `bson:"granter"`
	Grantee   string      `bson:"grantee"`
	Allowance interface{} `bson:"allowance"`
}

func (m *DocTxMsgGrantAllowance) GetType() string {
	return MsgTypeGrantAllowance
}

func (m *DocTxMsgGrantAllowance) BuildMsg(v interface{}) {
	msg := v.(*MsgGrantAllowance)
	m.Granter = msg.Granter
	m.Grantee = msg.Grantee
	marshalJSON, err := codec.GetMarshaler().MarshalJSON(msg.Allowance)
	if err != nil {
		logrus.Errorf("DocTxMsgGrantAllowance codec.GetMarshaler().MarshalJSON err:%v", err)
	}

	var msgInterface interface{}
	if err = json.Unmarshal(marshalJSON, &msgInterface); err != nil {
		logrus.Errorf("DocTxMsgGrantAllowance json.Unmarshal err:%v", err)
	}
	m.Allowance = msgInterface
}

func (m *DocTxMsgGrantAllowance) HandleTxMsg(v sdk.Msg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgGrantAllowance)
	addrs = append(addrs, msg.Granter, msg.Grantee)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
