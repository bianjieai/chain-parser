package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestNode() {
	cases := []SubTest{
		{
			"TestCreate",
			remove,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func remove(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A96010A93010A212F69726974616D6F642E6E6F64652E4D736752656D6F766556616C696461746F72126E0A4030384444333335453636434330344531463734453637393743373446304546304131423836313043444533463145414243393239424134373730414430443046122A696161317A6476657333783977716D6C34726C6771797A797833657238786C6E6E6A6471757A6667716512640A4C0A400A192F636F736D6F732E63727970746F2E736D322E5075624B657912230A210273EAD3067D3F355A8B945FCA1F8E25D92FA9B8FA095B58CC407AB7FC8E97C07312040A02080118CCE80312140A0E0A0475676173120632303030303010C09A0C1A4020B788D6B5D2E51ED8E75403A59116ECB1FE14E546576323488027B94B4555222A638F7E00E087B7C32F5AC7A1B412EF5C850A94EEEC589677C03B8A70B6F121")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if identityDoc, ok := s.Node.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(identityDoc))
		}
	}
}
