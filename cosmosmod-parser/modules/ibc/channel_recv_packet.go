package ibc

import (
	"fmt"
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	"github.com/bianjieai/chain-parser/common-parser/utils"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
)

type DocMsgRecvPacket struct {
	PacketId        string `bson:"packet_id"`
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(msg.ProofCommitment)
	m.Packet = loadPacket(msg.Packet)
	m.PacketId = fmt.Sprintf("%v%v%v%v%v", msg.Packet.SourcePort, msg.Packet.SourceChannel,
		msg.Packet.DestinationPort, msg.Packet.DestinationChannel, msg.Packet.Sequence)

}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs   []string
		icaMsgs []sdk.Msg
	)

	msg := v.(*MsgRecvPacket)
	switch msg.Packet.DestinationPort {
	case _dstPortMap[NFTTransferDscPortKey]:
		packetData := UnmarshalNftPacketData(msg.Packet.GetData())
		addrs = append(addrs, packetData.Receiver)
	case icatypes.HostPortID:
		_, icaMsgs = UnmarshalICAPacketData(msg.Packet.GetData())
	case _dstPortMap[IBCTransferDscPortKey]:
		packetData := UnmarshalPacketData(msg.Packet.GetData())
		addrs = append(addrs, packetData.Receiver)
	default:
		packetData := UnmarshalPacketData(msg.Packet.GetData())
		addrs = append(addrs, packetData.Receiver)
	}
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	msgDocInfo := CreateMsgDocInfo(v, handler)
	msgDocInfo.Msgs = append(msgDocInfo.Msgs, icaMsgs...)
	return msgDocInfo
}
