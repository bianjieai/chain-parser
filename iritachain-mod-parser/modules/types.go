package msgs

import (
	evm "github.com/tharsis/ethermint/x/evm/types"
)

const (
	MsgTypeEthereumTx = "ethereum_tx"
)

type (
	//evm
	MsgEthereumTx = evm.MsgEthereumTx
)
