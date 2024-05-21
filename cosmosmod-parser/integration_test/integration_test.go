package integration

import (
	"github.com/stretchr/testify/suite"
	msg_parser "gitlab.bianjie.ai/chain-parser/cosmosmod-parser"
	"testing"
)

type IntegrationTestSuite struct {
	msg_parser.MsgClient
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
	s.MsgClient = msg_parser.NewMsgClient()
}
