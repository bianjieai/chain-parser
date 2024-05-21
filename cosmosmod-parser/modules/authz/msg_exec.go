package authz

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	commoncodec "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
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
	msgs := make([]interface{}, 0, len(msg.Msgs))
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
	}
	m.Msgs = msgs
}

func (m *DocMsgExec) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgExec)
	addrs = append(addrs, msg.Grantee)

	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
