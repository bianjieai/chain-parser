package integration

import (
	"encoding/hex"
	"fmt"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
	"gitlab.bianjie.ai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestBank() {
	cases := []SubTest{
		{
			"TestSend",
			send,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func send(s IntegrationTestSuite) {
	txBytes, err := hex.DecodeString("0AC4020AB2010A1A2F697269736D6F642E6E66742E4D7367497373756544656E6F6D1293010A2675707469636B3832643836343233366463626361363562623762626633366666643465373938121CE5A4A7E5A3B0E782B9E79A845F313636393737383035345F327279371A1F7B2274797065223A222F75707469636B2E736F7576656E697243617264227D222A696161316871397A717A773061683335756463766D38797039746A3675637277736771343864327279370A89010A1C2F636F736D6F732E62616E6B2E763162657461312E4D736753656E6412690A2A696161316871397A717A773061683335756463766D38797039746A367563727773677134386432727937122A69616131737A6B6861713337767379687536396B677363786D6E307377666C376174786B7A716C3464671A0F0A05756E79616E1206313030303030120120126A0A510A460A1F2F636F736D6F732E63727970746F2E736563703235366B312E5075624B657912230A21037E013FFC7D81CC195D4F80F1D3DCC0E5E021DADE6AA65CE21628683404225F3C12040A020801188C0D12150A0F0A05756E79616E120635363030303010C08B111A40AE66721FD36B308AD4197ED2076BD15BE2F543781A13F9825E798CEF1CF76CD26D6916BDE77ECE6B868EB362457158310C9D6062ED0E06578F3CA14B5C75F1C2")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := codec.GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Bank.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
