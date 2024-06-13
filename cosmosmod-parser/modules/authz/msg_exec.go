package authz

import (
	"encoding/json"
	commoncodec "github.com/bianjieai/chain-parser/common-parser/codec"
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
)

type (
	DocMsgExec struct {
		Grantee string        `bson:"grantee"`
		Msgs    []interface{} `bson:"msgs"`
	}
)

func (m *DocMsgExec) GetType() string {
	return MsgTypeAuthzExec
}

func (m *DocMsgExec) BuildMsg(v interface{}) {
	msg := v.(*MsgExec)
	m.Grantee = msg.Grantee
	msgs, _ := UnmarshalExecData(msg)
	m.Msgs = msgs
}

func (m *DocMsgExec) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgExec)
	addrs = append(addrs, msg.Grantee)
	_, execMsgs := UnmarshalExecData(msg)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	msgDocInfo := CreateMsgDocInfo(v, handler)
	msgDocInfo.Msgs = append(msgDocInfo.Msgs, execMsgs...)
	return msgDocInfo
}

func UnmarshalExecData(msg *MsgExec) ([]interface{}, []sdk.Msg) {
	msgs := make([]interface{}, 0, len(msg.Msgs))
	var execMsgs []sdk.Msg
	for _, message := range msg.Msgs {
		marshalJSON, err := commoncodec.GetMarshaler().MarshalJSON(message)
		if err != nil {
			logrus.Errorf("DocMsgExec commoncodec.GetMarshaler().MarshalJSON err:%v", err)
		}

		var msgInterface interface{}
		if err = json.Unmarshal(marshalJSON, &msgInterface); err != nil {
			logrus.Errorf("DocMsgExec json.Unmarshal err:%v", err)
		}

		msgs = append(msgs, msgInterface)
		var data sdk.Msg
		if err := commoncodec.GetMarshaler().UnpackAny(message, &data); err == nil {
			execMsgs = append(execMsgs, data)
		}
	}
	return msgs, execMsgs
}
