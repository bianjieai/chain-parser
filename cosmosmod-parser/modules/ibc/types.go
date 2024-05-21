package ibc

import (
	"fmt"
	icoreclient "github.com/cosmos/ibc-go/v7/modules/core/02-client/types"
	icorechannel "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	cdc "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
	. "gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules"
)

func loadPacket(packet icorechannel.Packet) Packet {
	var data interface{}
	switch packet.SourcePort {
	case "nft-transfer":
		//todo The premise is irishub, not all chains.
		// This may not be very accurate, as we will need to implement the requirement of multiple ports in the future, which means that multiple different ports can transfer nf
		data = UnmarshalNftPacketData(packet.GetData())
	default:
		data = UnmarshalPacketData(packet.GetData())
	}
	return Packet{
		Sequence:           int64(packet.Sequence),
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               data,
		TimeoutTimestamp:   int64(packet.TimeoutTimestamp),
		TimeoutHeight:      loadHeight(packet.TimeoutHeight)}
}

func UnmarshalNftPacketData(bytesdata []byte) NftPacketData {
	var (
		packetData NonFungibleTokenIbcPacketData
		data       NftPacketData
	)
	err := cdc.GetMarshaler().UnmarshalJSON(bytesdata, &packetData)
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(packetData), &data)
	return data
}

// NonFungibleTokenPacketData
type NftPacketData struct {
	ClassId   string   `bson:"class_id" json:"class_id"`
	ClassUri  string   `bson:"class_uri" json:"class_uri"`
	ClassData string   `bson:"class_data" json:"class_data"`
	TokenIds  []string `bson:"token_ids" json:"token_ids"`
	TokenUris []string `bson:"token_uris" json:"token_uris"`
	TokenData []string `bson:"token_data" json:"token_data"`
	Sender    string   `bson:"sender" json:"sender"`
	Receiver  string   `bson:"receiver" json:"receiver"`
	Memo      string   `bson:"memo" json:"memo"`
}

func UnmarshalPacketData(bytesdata []byte) PacketData {
	var (
		packetData FungibleTokenPacketData
		data       PacketData
	)
	err := cdc.GetMarshaler().UnmarshalJSON(bytesdata, &packetData)
	if err != nil {
		fmt.Println(err.Error())
	}
	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(packetData), &data)
	return data
}

func loadHeight(height icoreclient.Height) Height {
	return Height{
		RevisionNumber: int64(height.RevisionNumber),
		RevisionHeight: int64(height.RevisionHeight)}
}

func loadChannel(channel icorechannel.Channel) Channel {
	return Channel{
		State:          int32(channel.State),
		Ordering:       int32(channel.State),
		Version:        channel.Version,
		ConnectionHops: channel.ConnectionHops,
		Counterparty: ChannelCounterparty{
			ChannelId: channel.Counterparty.ChannelId,
			PortId:    channel.Counterparty.PortId,
		},
	}
}

type Channel struct {
	State          int32               `bson:"state"`
	Ordering       int32               `bson:"ordering"`
	Counterparty   ChannelCounterparty `bson:"counterparty"`
	ConnectionHops []string            `bson:"connection_hops"`
	Version        string              `bson:"version"`
}
type ChannelCounterparty struct {
	PortId    string `bson:"port_id"`
	ChannelId string `bson:"channel_id"`
}

type Height struct {
	RevisionNumber int64 `bson:"revision_number"`
	RevisionHeight int64 `bson:"revision_height"`
}

// Packet defines a type that carries data across different chains through IBC
type Packet struct {
	Sequence           int64       `bson:"sequence"`
	SourcePort         string      `bson:"source_port"`
	SourceChannel      string      `bson:"source_channel"`
	DestinationPort    string      `bson:"destination_port"`
	DestinationChannel string      `bson:"destination_channel"`
	Data               interface{} `bson:"data"`
	TimeoutHeight      Height      `bson:"timeout_height"`
	TimeoutTimestamp   int64       `bson:"timeout_timestamp"`
}

// FungibleTokenPacketData
type PacketData struct {
	Denom    string `bson:"denom" json:"denom"`
	Amount   string `bson:"amount" json:"amount"`
	Sender   string `bson:"sender" json:"sender"`
	Receiver string `bson:"receiver" json:"receiver"`
}

func GetIbcPacketDenom(packet Packet, packetMsgDenom string) string {
	if utils.ReceiverChainIsSource(packet.SourcePort, packet.SourceChannel, packetMsgDenom) {
		// remove prefix added by sender chain
		voucherPrefix := utils.GetDenomPrefix(packet.SourcePort, packet.SourceChannel)
		unprefixedDenom := packetMsgDenom[len(voucherPrefix):]
		denom := unprefixedDenom
		denomTrace := utils.ParseDenomTrace(unprefixedDenom)
		if denomTrace.Path != "" {
			denom = utils.IBCDenom(fmt.Sprintf("%v/%v", denomTrace.Path, denomTrace.BaseDenom))
		}
		return denom
	}

	// since SendPacket did not prefix the denomination, we must prefix denomination here
	sourcePrefix := utils.GetDenomPrefix(packet.DestinationPort, packet.DestinationChannel)
	// NOTE: sourcePrefix contains the trailing "/"
	prefixedDenom := sourcePrefix + packetMsgDenom
	denomTrace := utils.ParseDenomTrace(prefixedDenom)
	if denomTrace.Path != "" {
		return utils.IBCDenom(fmt.Sprintf("%v/%v", denomTrace.Path, denomTrace.BaseDenom))
	}

	return denomTrace.BaseDenom
}
