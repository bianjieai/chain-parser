package integration

import (
	"encoding/hex"
	"fmt"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
	"gitlab.bianjie.ai/chain-parser/cosmosmod-parser/modules/ibc"
)

func (s IntegrationTestSuite) TestIbc() {
	cases := []SubTest{
		{
			"CreateClient",
			CreateClient,
		},
		{
			"GetIbcPacketDenom",
			GetIbcPacketDenom,
		}, {
			"MsgAcknowledgement",
			MsgAcknowledgement,
		},
		{
			"MsgRecvPacket",
			MsgRecvPacket,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func CreateClient(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0a90030a8d030a232f6962632e636f72652e636c69656e742e76312e4d7367437265617465436c69656e7412e5020aaa010a2b2f6962632e6c69676874636c69656e74732e74656e6465726d696e742e76312e436c69656e745374617465127b0a09626966726f73742d321204080110031a03088c0622030884072a0308d80432003a06080210a38c0d42190a090801180120012a0100120c0a02000110211804200c300142190a090801180120012a0100120c0a02000110201801200130014a07757067726164654a10757067726164656449424353746174651286010a2e2f6962632e6c69676874636c69656e74732e74656e6465726d696e742e76312e436f6e73656e737573537461746512540a0c08aadeb9800610bbadf08a0112220a2051c77a4f5ac9d60247b465bb42d81f3837bbe54a5e01e394dd2369fe089c26141a20b8c97bc5436f53aac003e94d194e75bc4167e6874ebbc2a8b327912bd6ee2f551a2d636f736d6f73313675726e79677a3472726a786c68793578386b6e336167377030756668756b6c79377464613212640a4e0a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a210265facc37ebf82d3732778610528478e402832ec5d1913efe13b363b8ec58054312040a02080112120a0c0a047562696712043234373410ff84061a40138e01c63715448b3a6a2546075c46edb138bbcda005fe9914c1fd45eccfdc175eeb5b4191faacd45674da4aa568e654e1d1bc338a9deef0623421c4cecc1922")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Ibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
func MsgAcknowledgement(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0adf070adc070a272f6962632e636f72652e6368616e6e656c2e76312e4d736741636b6e6f776c656467656d656e7412b0070ae001080112087472616e736665721a0a6368616e6e656c2d333122087472616e736665722a096368616e6e656c2d3032a1017b22616d6f756e74223a223130222c2264656e6f6d223a227472616e736665722f6368616e6e656c2d33312f7374616b65222c227265636569766572223a22696161316a7a386b74346e307566733475796730736d6536367368756c32716b68776476397a77643676222c2273656e646572223a22696161317376616e6e6876327a617865667138336d37747265673037387564666b33376c35726b787466227d3a0310d95940dcd396d0b1b2e0d81612117b22726573756c74223a2241513d3d227d1a86050a82030aff020a3261636b732f706f7274732f7472616e736665722f6368616e6e656c732f6368616e6e656c2d302f73657175656e6365732f31122008f7557ed51826fe18d84512bf24ec75001edbaf2123a477df72a0a9f3640a7c1a0d0801180120012a050002f6a301222d080112060204f6a301201a212073a0508aab36e7909b65372846f953a88a8d5aaccadfe83efda509fe09de41b7222d080112060406f6a301201a212066de425ace458527d718e2357799410d8c1603a978ef54a5ce0e24b18dcf9e30222d08011206060ef6a301201a2120840b5018133113e4b452d94ec6643d3f7c3402f82cca1ac1a8c9ad3571afcdce222d08011206081af6a301201a21205ff733bc3ddca5e265d7de48a72aa6842650d160202b17373080be57cadb5ab7222d080112060a2af6a301201a212087babc6e6d3f1a473af42a986bce716a1cd07dd00bf66bcc88b380f720d8ee84222d080112060e7cf6a301201a2120ae211f123ec9d5b21dec5ceaad2a6005bbe6624433ecd61c0686c5531a0a4e650afe010afb010a036962631220d1646e332046dcdcdd61ea2285fe24113b42e17a01288cea9a0c8049c5769a561a090801180120012a0100222508011221016c2870e99eaa2453e238dd6af9c7660109a8a9667cff72a26c031cb1594f7353222708011201011a205dfd19ed456064093ac720743ac192a01a3cd3687ae9b88a0fd5e134019d6a16222708011201011a20fa3672853e691148a59a717c6fb7c3f4c27ae5adf3d10f1609a4c9a3a9a37e4322250801122101128f4bee7ae00fd3197f860dcbece39442d264f8fae3fddee86c925656baa159222708011201011a2023a7499b9940678a0fc560d0b9b9d749919ab67ad81896765bf73b9fda33bb44220310fc512a2a696161316a7a386b74346e307566733475796730736d6536367368756c32716b68776476397a7764367612660a500a460a1f2f636f736d6f732e63727970746f2e736563703235366b312e5075624b657912230a2103adbd43a6a6d2e451d0378dc0bf765a21ebc3eac27c2605f1107794b64be4568312040a020801181a12120a0c0a05756e79616e120336303010e0a7121a40ee9e69fa27ae3d8e32001d365feaae3000afbbdb9e93c58bdd4c220104662ece1f86292deed96004bb611b6efe34556a15cdbf800271fb45557a43d586d58159")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Ibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func MsgRecvPacket(s IntegrationTestSuite) {
	codec.SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AE10C0ADE0C0A222F6962632E636F72652E6368616E6E656C2E76312E4D7367526563765061636B657412B70C0AED01080112087472616E736665721A0B6368616E6E656C2D32333922087472616E736665722A0B6368616E6E656C2D31363432A7017B22616D6F756E74223A2232222C2264656E6F6D223A227472616E736665722F6368616E6E656C2D3233302F7577773730313031222C227265636569766572223A22696161316E36797835376B3372703270666B3872656B616B666168793734306E79363073393463336339222C2273656E646572223A22636F736D6F733168667179396B6E70613335746530687A37663278707939396B7A386C6A68347375783932667A227D3A07080110BF80AC0240C0AAB4D286F4E58E1712910A0A8B080A88080A3B636F6D6D69746D656E74732F706F7274732F7472616E736665722F6368616E6E656C732F6368616E6E656C2D3233392F73657175656E6365732F3112204099DA84789F222A6F497990DFAD527DC411B1754A963E501C73E14D9F6D76821A0E0801180120012A060002F4C7CF09222C080112280204F4C7CF09202D684FF30F412025DE5777B60D63C235FBD35EE3DCAEA13CBE783B81F35BD25D20222E080112070408F4C7CF09201A2120A4F1BBDD2ABB1D9F288599945B6696D10D769CA718B1AB96B549FA482D44914B222C08011228060EF4C7CF09206366CA3FB365AD7C2F3F9A40A43455DCD2A6F44AA03286A3D2CB8E3F2530BD6720222E08011207081AF4C7CF09201A2120CBAD822071F43D24D3F480BF98BC9CCD590C5C8ACF0A09853C8EE45D550D6BCA222C080112280A2AF4C7CF0920A52368D1DAE63BEE347EF3F5C13F0CC983B07FAF50F832B4147091D48148D36320222E080112070C48F4C7CF09201A21201F1096068D652A09E7A0D955CE656F5D45D0116B9C0F67592492C724B6788AB8222F080112080E8001F4C7CF09201A21206048073BB50AD89ED7E083FB249CFD0B3387DEC1A3FF325DBEB2B4FABDC38B54222D0801122910E601F4C7CF0920EF9D31A76CA6709E79733F7A7093A15A972E72FE6EF5C935FAD9A88FF43BAB1B20222F0801120814A606F4C7CF09201A212037681819813505FA673960143172C06280B8B4995E01C81A154D19F84E50338A222F0801120818EC0DF4C7CF09201A21205E9265EBFF186DEC1617C3527C01B008E0546FA623630DDD3BC9753CA27B48D2222F080112081A961CF4C7CF09201A2120634B204C77EC145989873277008A3EE968B772F3D35B2DA3449E46C44A8A9D7C222F080112081CCE33F4C7CF09201A2120DAEF8845C999302B4F923F47720C37ADC863A4DFD274F954B49EC350DB6481FC222F080112081EB855F4C7CF09201A2120206024B5788417567552AF53CF7A480C2D0AC3555EB65FC4882B6828FC79FBB022300801120920BAAE01F4C7CF09201A2120CB970E4A09EF06DB70131CF644A62A4AEBB9B8F2DAC8D40C466DCC0F23D30C322230080112092286C702F4C7CF09201A2120F2F256CFA5595DE22BAF3913D67D0A8067E2BABAF3FA85D4FC53E951B6CF8B042230080112092486F205F4C7CF09201A2120E2B323F3804D77875AF50A162CE3829B8D5A117BBF1F24658018AF5F605C7D1522300801120926B8B509F4C7CF09201A2120CC7D140A1F9AD09868836D66D53E9AA46D68E251E31DFC3853C9020599F25C6F2230080112092AAAFC22F4C7CF09201A21205697E5350C466059A3015789CFAF3A858FFAC8B99A2DD6229C2601C8C6BB6B26222E0801122A2CA4B846F4C7CF09208DF88243FC712B997FAB3A010EDBC34681B4AAA870527896F93D90FF465DCBA6200A80020AFD010A03696263122075AE4C3968EF7BE20F564592288BB812C6486F3C909E1E88D703E9137E78D1E51A090801180120012A0100222708011201011A20BB5845E297960E4B4D4568A5BBB6B5BB3FFE247899EA26ED53B1485F290BDD9F222708011201011A2092DFCED16C0DE6DEE720EC0CDF691EB3A5C64AB5AE9534939AD90748CDEB7ACE222708011201011A20EF8156A0550E50596E511AF0A2B42D3D40BDC8C6628D25A2879A9BC3F39A54322225080112210199D1CEC3FE01EC38B6B2F6B915339E0BCF29688E1314101CE53AB0DC8A35D28E222708011201011A2070AF301F255EA35D9DDA0F98F34B5D5B49B849B373B23BCCED333E9A31C293B21A0510FDE3E704222A696161316E397775786B326436397874307139393672646574657771673675776438726D68647A30613312650A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102E81AD04ACBA1746112AB60C282871A8E8E4BC0CAEF2884752FA4DDDA555E34DB12040A02080118931512100A0A0A05756972697312013110D9A2081A4019DB58935F6463FC6FF9EB8027252758EC1DE24B359E6BE000B6FD7DF2CF922207F7D82F3A0B9E1C3A1A7FD694AD677AEDFDDB718D2D186EB5C89B050CEE9731")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Ibc.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func GetIbcPacketDenom(s IntegrationTestSuite) {
	packet := ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-9",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
		Data: ibc.PacketData{
			Denom:  "transfer/channel-9/umuon",
			Amount: "1",
		},
	}

	fmt.Println("Atom Iris => Cosmos: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))
	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-9",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-1",
		Data: ibc.PacketData{
			Denom:  "unyan",
			Amount: "1",
		},
	}

	fmt.Println("Iris Iris => Cosmos: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))

	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-1",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-9",
		Data: ibc.PacketData{
			Denom:  "umuon",
			Amount: "1",
		},
	}

	fmt.Println("Atom Cosmos => Iris: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))
	packet = ibc.Packet{
		SourcePort:         "transfer",
		SourceChannel:      "channel-1",
		DestinationPort:    "transfer",
		DestinationChannel: "channel-9",
		Data: ibc.PacketData{
			Denom:  "transfer/channel-1/unyan",
			Amount: "1",
		},
	}

	fmt.Println("Iris Cosmos => Iris: ", ibc.GetIbcPacketDenom(packet, packet.Data.Denom))

}
