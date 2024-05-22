package msgs

import (
	cdc "github.com/bianjieai/chain-parser/common-parser/codec"
)

func UnmarshalTibcAcknowledgement(bytesdata []byte) string {
	var result TIBCAcknowledgement
	if err := cdc.GetMarshaler().Unmarshal(bytesdata, &result); err == nil {
		return result.String()
	}
	return ""
}
