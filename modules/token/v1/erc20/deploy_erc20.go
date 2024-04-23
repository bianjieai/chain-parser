package erc20

import (
	. "gitlab.bianjie.ai/chain-parser/common-parser/modules"
	. "gitlab.bianjie.ai/chain-parser/irismod-parser/modules"
)

// MsgDeployERC20 defines an SDK message for DeployERC20
type DocMsgDeployERC20 struct {
	Symbol    string `bson:"symbol"`
	Name      string `bson:"name"`
	Scale     uint32 `bson:"scale"`
	MinUnit   string `bson:"min_unit"`
	Authority string `bson:"authority"`
}

func (m *DocMsgDeployERC20) GetType() string {
	return MsgTypeDeployERC20
}

func (m *DocMsgDeployERC20) BuildMsg(v interface{}) {
	msg := v.(*MsgDeployERC20)

	m.Symbol = msg.Symbol
	m.Name = msg.Name
	m.Scale = msg.Scale
	m.MinUnit = msg.MinUnit
	m.Authority = msg.Authority
}

func (m *DocMsgDeployERC20) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var addrs []string

	msg := v.(*MsgDeployERC20)
	addrs = append(addrs, msg.Authority)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
