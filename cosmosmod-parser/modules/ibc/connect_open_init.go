package ibc

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/bianjieai/chain-parser/common-parser/utils"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

type DocMsgConnectionOpenInit struct {
	ClientId     string       `bson:"client_id"`
	Counterparty Counterparty `bson:"counterparty"`
	Version      Version      `bson:"version"`
	DelayPeriod  int64        `bson:"delay_period"`
	Signer       string       `bson:"signer"`
}

func (m *DocMsgConnectionOpenInit) GetType() string {
	return MsgTypeConnectionOpenInit
}

func (m *DocMsgConnectionOpenInit) BuildMsg(v interface{}) {
	msg := v.(*MsgConnectionOpenInit)
	m.ClientId = msg.ClientId
	m.DelayPeriod = int64(msg.DelayPeriod)
	m.Signer = msg.Signer
	if msg.Version != nil {
		m.Version = Version{
			Identifier: msg.Version.Identifier,
			Features:   msg.Version.Features,
		}
	}
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(msg.Counterparty), &m.Counterparty)

}

func (m *DocMsgConnectionOpenInit) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgConnectionOpenInit)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
