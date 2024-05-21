package integration

import (
	"encoding/hex"
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestOracle() {
	cases := []SubTest{
		{
			"EditFeed",
			EditFeed,
		},
		{
			"PauseFeed",
			PauseFeed,
		},
		{
			"StartFeed",
			StartFeed,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func CreateFeed(s IntegrationTestSuite) {
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
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Oracle.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func EditFeed(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A99010A96010A1B2F697269736D6F642E6F7261636C652E4D7367456469744665656412770A03737A63180A222A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E3370280232100A057569726973120731303030303030380A40014A2A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E337012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210251CA6230319F0F129D2DE9CA5739DE9DCB4C685485B32AE742A8BBAA55E2CFD712040A020801180512150A0F0A057569726973120631303030303010C09A0C1A400ED623236B70863813C000FD740801367308AA6F1318FCE2F28D4EC76C56914E3B3AC0CFED57D733BD359083A8BE44459792B41659E231B2C03778B8DB278C50")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Oracle.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func PauseFeed(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A580A560A1C2F697269736D6F642E6F7261636C652E4D736750617573654665656412360A08666565646E616D65122A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E337012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210251CA6230319F0F129D2DE9CA5739DE9DCB4C685485B32AE742A8BBAA55E2CFD712040A020801180612150A0F0A057569726973120631303030303010C09A0C1A405A94BF92BE741C6D3DB320D67DE4ED66580BAB8B95F96D60F444F7A61B3ED21A1865009EDC5F7D5F1F39E2268D28651F81E6EB842F94E40A8E6EC957595FD512")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Oracle.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}

func StartFeed(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A580A560A1C2F697269736D6F642E6F7261636C652E4D736753746172744665656412360A08666565646E616D65122A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E337012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210251CA6230319F0F129D2DE9CA5739DE9DCB4C685485B32AE742A8BBAA55E2CFD712040A020801180712150A0F0A057569726973120631303030303010C09A0C1A4069A94508D54D9DC441C78F8192AFFE8014B56282FE5E25DF546FAF4D3F1F42FE087878A1727CC4381ABECFB024AAAEA4BF6622F520FAF24803ED800A7640BDA2")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Oracle.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
