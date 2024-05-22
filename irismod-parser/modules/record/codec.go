package record

import (
	"github.com/bianjieai/chain-parser/common-parser/codec"
	"github.com/irisnet/irismod/modules/record"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
