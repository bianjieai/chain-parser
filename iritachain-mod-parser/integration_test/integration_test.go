package integration

import (
	iritachain_mod_parser "github.com/bianjieai/chain-parser/iritachain-mod-parser"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationTestSuite struct {
	iritachain_mod_parser.MsgClient
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.MsgClient = iritachain_mod_parser.NewMsgClient()
}
