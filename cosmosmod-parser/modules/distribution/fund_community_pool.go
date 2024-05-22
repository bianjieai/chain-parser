package distribution

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/bianjieai/chain-parser/common-parser/types"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
)

// msg struct for delegation withdraw for all of the delegator's delegations
type DocTxMsgFundCommunityPool struct {
	Amount    []types.Coin `bson:"amount"`
	Depositor string       `bson:"depositor"`
}

func (m *DocTxMsgFundCommunityPool) GetType() string {
	return MsgTypeMsgFundCommunityPool
}

func (m *DocTxMsgFundCommunityPool) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgFundCommunityPool)
	m.Depositor = msg.Depositor
	m.Amount = types.BuildDocCoins(msg.Amount)
}
func (m *DocTxMsgFundCommunityPool) HandleTxMsg(v SdkMsg) MsgDocInfo {

	var addrs []string

	msg := v.(*MsgFundCommunityPool)
	addrs = append(addrs, msg.Depositor)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
