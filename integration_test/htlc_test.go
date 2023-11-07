package integration

import (
	"encoding/hex"
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestHtlc() {
	cases := []SubTest{
		{
			"createHTLC",
			createHTLC,
		},
		{
			"claimHTLC",
			claimHTLC,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func createHTLC(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0ACC020ABB020A1B2F697269736D6F642E68746C632E4D736743726561746548544C43129B020A2A696161316571766B667468747272393367347039717370703534773664746A74726E3237617237727077122A6961613137636A64673633746879327666717676676A356C667635647033333974306C726132677138751A2B74626E6231726C6B71397A7836637A6536756D63387235377235677277773774756E78736C376432347773222B74626E623171776735353361726B773567783934706770793838656B7134756437356D7576717A6A7967612A1B0A0A68746C74626362757364120D31303030303030303030303030324037303962346636656331363664393439393539613866613232386437386130336264373930653036623836366363313135376437653839343331353531613033388E90B3900640644801120C68746C63207478206D656D6F126A0A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801180612160A100A0575697269731207313030303030301080B5181A404CEB594F7FA339500A82C7028A40B10C86D34E75CC2C8B055335C94D28B92A1B765B1CD5703455BBC2B7C58E14969F2C960F83BC40BC2F71031A5A015CB46D27")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if bankDoc, ok := s.Htlc.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func claimHTLC(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AE3010ACF010A1A2F697269736D6F642E68746C632E4D7367436C61696D48544C4312B0010A2A696161316571766B667468747272393367347039717370703534773664746A74726E32376172377270771240323333414132323541433632413545303030463132463245313543434331304241393434423745454646454543453031313938374439394444443234343135441A4034454542314539433745384534383935363643303843383742453443463836313736314144323035394345434346414444314142424331464442323334383030120F636C61696D2068746C63206D656D6F126A0A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801180712160A100A05756972697312073130303030303010C09A0C1A40E00E34FA3743767A4C55E501CD92AA6858EA540CECE86B9D8E660C83C865E78C6E5733C55DD788185E34583B2C08D34F265C56EEDAB3E8847AE38D5C80B9D37D")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if bankDoc, ok := s.Htlc.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
