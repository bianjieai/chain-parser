package staking

import (
	. "github.com/bianjieai/chain-parser/common-parser/modules"
	. "github.com/bianjieai/chain-parser/cosmosmod-parser/modules"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StakingClient struct {
}

func NewClient() StakingClient {
	return StakingClient{}
}

func (staking StakingClient) HandleTxMsg(v sdk.Msg) (MsgDocInfo, bool) {
	switch msg := v.(type) {
	case *MsgBeginRedelegate:
		docMsg := DocTxMsgBeginRedelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeBeginUnbonding:
		docMsg := DocTxMsgBeginUnbonding{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeDelegate:
		docMsg := DocTxMsgDelegate{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeEdit:
		docMsg := DocMsgEditValidator{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgStakeCreate:
		docMsg := DocTxMsgCreateValidator{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgCancelUnbondingDelegation:
		docMsg := DocTxMsgCancelUnbondingDelegation{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgValidatorBond:
		docMsg := DocTxMsgValidatorBond{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgUnbondValidator:
		docMsg := DocTxMsgUnbondValidator{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgTokenizeShares:
		docMsg := DocTxMsgTokenizeShares{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgDisableTokenizeShares:
		docMsg := DocTxMsgDisableTokenizeShares{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgEnableTokenizeShares:
		docMsg := DocTxMsgEnableTokenizeShares{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgRedeemTokensForShares:
		docMsg := DocTxMsgRedeemTokensForShares{}
		return docMsg.HandleTxMsg(msg), true
	case *MsgTransferTokenizeShareRecord:
		docMsg := DocTxMsgTransferTokenizeShareRecord{}
		return docMsg.HandleTxMsg(msg), true
	}
	return MsgDocInfo{}, false
}
