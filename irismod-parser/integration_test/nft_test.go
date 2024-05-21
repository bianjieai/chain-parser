package integration

import (
	"encoding/hex"
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestNft() {
	cases := []SubTest{
		{
			"IssueDenom",
			IssueDenom,
		},
		{
			"TransferDenom",
			TransferDenom,
		},
		{
			"NFTMint",
			NFTMint,
		},
		{
			"NFTTransfer",
			NFTTransfer,
		},
		{
			"NFTEdit",
			NFTEdit,
		},
		{
			"NFTBurn",
			NFTBurn,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func IssueDenom(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AC4020AB2010A1A2F697269736D6F642E6E66742E4D7367497373756544656E6F6D1293010A2675707469636B3832643836343233366463626361363562623762626633366666643465373938121CE5A4A7E5A3B0E782B9E79A845F313636393737383035345F327279371A1F7B2274797065223A222F75707469636B2E736F7576656E697243617264227D222A696161316871397A717A773061683335756463766D38797039746A3675637277736771343864327279370A89010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412690A2A696161316871397A717A773061683335756463766D38797039746A367563727773677134386432727937122A69616131737A6B6861713337767379687536396B677363786D6E307377666C376174786B7A716C3464671A0F0A05756E79616E1206313030303030120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21037E013FFC7D81CC195D4F80F1D3DCC0E5E021DADE6AA65CE21628683404225F3C12040A020801188C0D12150A0F0A05756E79616E120635363030303010C08B111A40AE66721FD36B308AD4197ED2076BD15BE2F543781A13F9825E798CEF1CF76CD26D6916BDE77ECE6B868EB362457158310C9D6062ED0E06578F3CA14B5C75F1C2")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func TransferDenom(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A89010A86010A1D2F697269736D6F642E6E66742E4D73675472616E7366657244656E6F6D12650A0B6E667464656E6F6D696432122A696161316571766B667468747272393367347039717370703534773664746A74726E32376172377270771A2A696161317376616E6E6876327A617865667138336D37747265673037387564666B33376C35726B78746612680A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801188A0212130A0D0A05756E79616E12043130303010C09A0C1A40D3DA8AF5F9349051586C0B05B628C82FFF6CC97A28895E6E351010C6052C59696C2C9D1C24A2B24AEE8BF7C7743CF759658FD40D210921D3D0B3B82C54C8F959")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func NFTMint(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0ADB010AD8010A172F697269736D6F642E6E66742E4D73674D696E744E465412BC010A0C617564696F696D6167657332120B64656E6F6D7374616765311A09E4BAABE9A398E9A398223C687474703A2F2F766964656F2E6368696E616E6577732E636F6D2F666C762F323031392F30342F32332F3430302F3131313737335F7765622E6D7034322A6961613163377933373073766D6D7539796877673766333063666D3730676471756874336C37793277673A2A696161316D643276326C757A68386E7733326C3465746D7067786B723939376679356B736E337676786D12670A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210206B3F0799D528B2288FF7F3C1EB25C28CB1640A8AA787107CB1583BEDC4578D812040A02080118D50212120A0C0A05756E79616E120334303010C09A0C1A40938CA0EDFBB85AC304EE61D5EC84270E97980461D8F6DBD9F68C60BC0BA8368F14278790B7A53105CE7734CB27B475926F89E3009CFEADE0C58F181937E266AF")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func NFTTransfer(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AE2010ADC010A1B2F697269736D6F642E6E66742E4D73675472616E736665724E465412BC010A11797A303135353639737377657772657772120B64656E6F6D7374616765311A0F5B646F2D6E6F742D6D6F646966795D220F5B646F2D6E6F742D6D6F646966795D2A0F5B646F2D6E6F742D6D6F646966795D322A696161316D643276326C757A68386E7733326C3465746D7067786B723939376679356B736E337676786D3A2A69616131366D7636686D7836677779656E3578756A327A6E7867753432667A637830793772676436386A420F5B646F2D6E6F742D6D6F646966795D120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A2102D2FF0141817477CA836A4F4159C1DF5D7F4462E8640F4426F88121FE8A5D6D2A12040A02080118900212150A0F0A05756E79616E12063830303030301080B5181A407AB890E594ABAE081D304688CA21F5E7FF5E9B4AB0FB906F80000172C03D18D95D352852AF68F63FE5D65F17F0836A7492378DB86B7B4F7DD3635BE3C1ED0703")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func NFTEdit(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A8E010A8B010A172F697269736D6F642E6E66742E4D7367456469744E465412700A066E6674696431120B6E667464656E6F6D6964311A096E6674206E616D653122086E667420757269312A096E6674206461746131322A696161316571766B667468747272393367347039717370703534773664746A74726E32376172377270773A0D6E66742075726920686173683112680A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A02080118880212130A0D0A05756E79616E12043130303010C09A0C1A40D44FC76AB318F536791B8A1B717984F04F91504E4355A8A38C8AB783DBF81AF611CDF9B58BFAD5ADE603FA26AFF0798699C37DBC8A1895868B70EE8BA38A4779")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func NFTBurn(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A5E0A5C0A172F697269736D6F642E6E66742E4D73674275726E4E465412410A066E6674696432120B6E667464656E6F6D6964311A2A696161316571766B667468747272393367347039717370703534773664746A74726E323761723772707712680A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801188B0212130A0D0A05756E79616E12043130303010C09A0C1A404788020041949BFF00885FA40E854589113DFB37A0D94CC3DF35D7F5C36030007053972DD53C55886E50A25CFF6D9EE5B6072E4735081FC1A47A0D20CC62D6F6")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
