package record

import (
	"github.com/irisnet/irismod/modules/record"
	"gitlab.bianjie.ai/chain-parser/common-parser/codec"
)

func init() {
	codec.RegisterAppModules(record.AppModuleBasic{})
}
