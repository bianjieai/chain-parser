package integration

import (
	irismod_parser "github.com/bianjieai/chain-parser/irismod-parser"
	"github.com/stretchr/testify/suite"
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
