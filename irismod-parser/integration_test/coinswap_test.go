package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestCoinswap() {
	cases := []SubTest{
		{
			"swapOrder",
			swapOrder,
		},
		{
			"addLiquidity",
			addLiquidity,
		},
		{
			"removeLiquidity",
			removeLiquidity,
		},
		{
			"addUnilateralLiquidity",
			addUnilateralLiquidity,
		},
		{
			"removeUnilateralLiquidity",
			removeUnilateralLiquidity,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func swapOrder(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0AF3010AED010A1E2F697269736D6F642E636F696E737761702E4D7367537761704F7264657212CA010A3F0A2A696161316571766B667468747272393367347039717370703534773664746A74726E323761723772707712110A05756972697312083535343537333937127F0A2A696161316571766B667468747272393367347039717370703534773664746A74726E323761723772707712510A446962632F45344638304437433934383943344231333146463933344335373734443545333843414333343144314144314637393931353738413932324638463946433935120931313030303030303018BEE98897062001120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801189C1112150A0F0A05756972697312063430303030301080B5181A4041C1A9F5FE6FD4F8529B2F0040699BD4D82883BB1B8DCE99D9073A372B27D07866F8B6507B544388F0E5DE7A6875EE8D4CBC145253E911D26719B3188960EDC9")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if bankDoc, ok := s.Coinswap.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func addLiquidity(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0ABF010AB9010A212F697269736D6F642E636F696E737761702E4D73674164644C69717569646974791293010A510A446962632F453446383044374339343839433442313331464639333443353737344435453338434143333431443141443146373939313537384139323246384639464339351209313030303030303030120835303030393938351A0235302094C78397062A2A696161316571766B667468747272393367347039717370703534773664746A74726E3237617237727077120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A020801189B1112150A0F0A05756972697312063430303030301080B5181A40488A4681917A48608D1C21C980C785BDDD1233469D31B18D1D70A67A56D34D8A718A7E419F6D5793B2A98F85F02E747A657AA5F236DEDB9B5C5E2493F9C70FD1")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if bankDoc, ok := s.Coinswap.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func removeLiquidity(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A8B010A85010A242F697269736D6F642E636F696E737761702E4D736752656D6F76654C6971756964697479125D0A130A066C70742D3231120931303030303530303012093139393939303036311A0931303030313530303020CF8FFF96062A2A696161316571766B667468747272393367347039717370703534773664746A74726E3237617237727077120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21038BE4539785F0621A19066EB2DA45C11DCC5FCCFF5D58E89C670BB80D251CC1B712040A02080118951112150A0F0A05756972697312063430303030301080B5181A406D936A57E5FCE5C0317B9C726AA1F1DA530B34880586AD6E148F3342665EBEBA4EC0A2FC484CE930844DB998C548B3F3E29E51A5DEEC258BE80BD468A27681BE")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if bankDoc, ok := s.Coinswap.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func addUnilateralLiquidity(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if doc, ok := s.Coinswap.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(doc))
		}
	}
}

func removeUnilateralLiquidity(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, irismod := range authTx.GetMsgs() {
		if doc, ok := s.Coinswap.HandleTxMsg(irismod); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(doc))
		}
	}
}
