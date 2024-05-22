package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/bianjieai/chain-parser/common-parser/utils"
)

func (s IntegrationTestSuite) TestIdentity() {
	cases := []SubTest{
		{
			"TestCreate",
			create,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func create(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0A660A640A242F69726974616D6F642E6964656E746974792E4D73674372656174654964656E74697479123C0A0E33373330353533333443344234452A2A696161317638786768303471376178676E396B737475646B656737726A303863636734656B733863387212640A4B0A400A192F636F736D6F732E63727970746F2E736D322E5075624B657912230A2102F39AB53A0984A46C94AFE1F16C687C6667762304101FFEC810F1FC822029861D12040A02080118D91E12150A0F0A047567617312073130303030303010E0A7121A40B6A6FD85C5EDCB70F0F982F1789055A592F35D3F741E17ACB91BDF6DF2989523026748F18F31F891B7DC2B737CB35641486445DC69D5D23AC9D4D339D86B983A")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if identityDoc, ok := s.Identity.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(identityDoc))
		}
	}
}
