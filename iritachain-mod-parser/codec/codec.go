package codec

import (
	commoncodec "github.com/bianjieai/chain-parser/common-parser/codec"
	amino "github.com/cosmos/cosmos-sdk/codec"
	types2 "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/auth/tx"
	enccodec "github.com/tharsis/ethermint/encoding/codec"
)

// 初始化账户地址前缀
func MakeEncodingConfig() {
	moduleBasics := module.NewBasicManager(commoncodec.AppModules...)
	cdc := amino.NewLegacyAmino()
	interfaceRegistry := types2.NewInterfaceRegistry()
	marshaler := amino.NewProtoCodec(interfaceRegistry)

	encodingConfig := params.EncodingConfig{
		InterfaceRegistry: interfaceRegistry,
		Marshaler:         marshaler,
		TxConfig:          tx.NewTxConfig(marshaler, tx.DefaultSignModes),
		Amino:             cdc,
	}

	enccodec.RegisterLegacyAminoCodec(encodingConfig.Amino)
	moduleBasics.RegisterLegacyAminoCodec(encodingConfig.Amino)
	enccodec.RegisterInterfaces(encodingConfig.InterfaceRegistry)
	moduleBasics.RegisterInterfaces(encodingConfig.InterfaceRegistry)

	commoncodec.Encodecfg = encodingConfig
}
