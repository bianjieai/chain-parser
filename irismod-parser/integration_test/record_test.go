package integration

import (
	"encoding/hex"
	"fmt"
	. "gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestRecord() {
	cases := []SubTest{
		{
			"RecordCreate",
			RecordCreate,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func RecordCreate(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A80010A7E0A1F2F697269736D6F642E7265636F72642E4D73674372656174655265636F7264125B0A2D0A256B736C64666A6F697364666B6C6572706F646C6B636C3233343332346C6B6C657270323339120473686169122A69616131646E6E786A636E6B63377672356B70356E3365796A7235703564393978636E783665786E337012690A500A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A210251CA6230319F0F129D2DE9CA5739DE9DCB4C685485B32AE742A8BBAA55E2CFD712040A020801180412150A0F0A057569726973120631303030303010C09A0C1A40FD1C7176394CE618C8230C3E456D2683D3C46F21AF4967C509BB4531F08D04A542137A3670378FD3F8F039D256B58170A60D842275ABF7FB1CB01720FB2FAE91")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Record.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
