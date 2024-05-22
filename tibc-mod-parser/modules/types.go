package msgs

import (
	tibctranfer "github.com/bianjieai/tibc-go/modules/tibc/apps/nft_transfer/types"
	tibcclient "github.com/bianjieai/tibc-go/modules/tibc/core/02-client/types"
	tibcpacket "github.com/bianjieai/tibc-go/modules/tibc/core/04-packet/types"
)

const (
	//tibc
	MsgTypeTIBCNftTransfer     = "tibc_nft_transfer"
	MsgTypeTIBCRecvPacket      = "tibc_recv_packet"
	MsgTypeTIBCUpdateClient    = "tibc_update_client"
	MsgTypeTIBCAcknowledgement = "tibc_acknowledge_packet"
	MsgTypeTIBCCleanPacket     = "clean_packet"
	MsgTypeTIBCRecvCleanPacket = "recv_clean_packet"
)

type (
	//tibc
	TIBCAcknowledgement        = tibcpacket.Acknowledgement
	NonFungibleTokenPacketData = tibctranfer.NonFungibleTokenPacketData
	MsgTIBCNftTransfer         = tibctranfer.MsgNftTransfer
	MsgTIBCUpdateClient        = tibcclient.MsgUpdateClient
	MsgTIBCRecvPacket          = tibcpacket.MsgRecvPacket
	MsgTIBCAcknowledgement     = tibcpacket.MsgAcknowledgement
	MsgCleanPacket             = tibcpacket.MsgCleanPacket
	MsgRecvCleanPacket         = tibcpacket.MsgRecvCleanPacket
)
