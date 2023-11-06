package integration

import (
	"encoding/hex"
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestRandom() {
	cases := []SubTest{
		{
			"RequestRand",
			RequestRand,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func RequestRand(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A680A660A202F697269736D6F642E72616E646F6D2E4D73675265717565737452616E646F6D1242080A122A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E3370180122100A05756972697312073130303030303012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210251CA6230319F0F129D2DE9CA5739DE9DCB4C685485B32AE742A8BBAA55E2CFD712040A020801180212150A0F0A057569726973120631303030303010C09A0C1A40510D3B6F38315F3B452CCB3BEB9BECF13F174A2E02C81B0FCA9E67C3FF1D45D1607E4C44B10123D15B019E843442FCD83F942009E2B4D020FAD4F53CCDBD68B0")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Random.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
