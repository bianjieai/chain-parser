package integration

import (
	"github.com/stretchr/testify/suite"
	irismod_parser "github.com/bianjieai/chain-parser/irismod-parser"
	"testing"
)

type IntegrationTestSuite struct {
	irismod_parser.MsgClient
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
	s.MsgClient = irismod_parser.NewMsgClient()
}
