package msgs

import (
	evm "github.com/evmos/ethermint/x/evm/types"
	coinswap "github.com/irisnet/irismod/modules/coinswap/types"
	farm "github.com/irisnet/irismod/modules/farm/types"
	htlc "github.com/irisnet/irismod/modules/htlc/types"
	mt "github.com/irisnet/irismod/modules/mt/types"
	nft "github.com/irisnet/irismod/modules/nft/types"
	oracle "github.com/irisnet/irismod/modules/oracle/types"
	random "github.com/irisnet/irismod/modules/random/types"
	record "github.com/irisnet/irismod/modules/record/types"
	service "github.com/irisnet/irismod/modules/service/types"
	tokenv1 "github.com/irisnet/irismod/modules/token/types/v1"
	tokenv1beta1 "github.com/irisnet/irismod/modules/token/types/v1beta1"
)

const (
	//coinswap
	MsgTypeAddLiquidity              = "add_liquidity"
	MsgTypeRemoveLiquidity           = "remove_liquidity"
	MsgTypeSwapOrder                 = "swap_order"
	MsgTypeAddUnilateralLiquidity    = "add_unilateral_liquidity"
	MsgTypeRemoveUnilateralLiquidity = "remove_unilateral_liquidity"
	//farm
	MsgTypeCreatePool     = "create_pool"
	MsgTypeCreateProposal = "create_pool_with_community_pool"
	MsgTypeDestroyPool    = "destroy_pool"
	MsgTypeAdjustPool     = "adjust_pool"
	MsgTypeStake          = "stake"
	MsgTypeUnstake        = "unstake"
	MsgTypeHarvest        = "harvest"
	//htlc
	MsgTypeCreateHTLC = "create_htlc"
	MsgTypeClaimHTLC  = "claim_htlc"
	MsgTypeRefundHTLC = "refund_htlc"
	//mt
	MsgTypeMTIssueDenom    = "mt_issue_denom"
	MsgTypeMTTransferDenom = "mt_transfer_denom"
	MsgTypeMintMT          = "mint_mt"
	MsgTypeTransferMT      = "transfer_mt"
	MsgTypeEditMT          = "edit_mt"
	MsgTypeBurnMT          = "burn_mt"
	//nft
	MsgTypeIssueDenom    = "issue_denom"
	MsgTypeTransferDenom = "transfer_denom"
	MsgTypeNFTMint       = "mint_nft"
	MsgTypeNFTTransfer   = "transfer_nft"
	MsgTypeNFTEdit       = "edit_nft"
	MsgTypeNFTBurn       = "burn_nft"
	//oracle
	TxTypeCreateFeed = "create_feed"
	TxTypeEditFeed   = "edit_feed"
	TxTypePauseFeed  = "pause_feed"
	TxTypeStartFeed  = "start_feed"
	//random
	TxTypeRequestRand = "request_rand"
	//record
	MsgTypeRecordCreate = "create_record"
	//service
	MsgTypeDefineService             = "define_service"
	MsgTypeBindService               = "bind_service"
	MsgTypeUpdateServiceBinding      = "update_service_binding"
	MsgTypeServiceSetWithdrawAddress = "service/set_withdraw_address"
	MsgTypeDisableServiceBinding     = "disable_service_binding"
	MsgTypeEnableServiceBinding      = "enable_service_binding"
	MsgTypeRefundServiceDeposit      = "refund_service_deposit"
	MsgTypeCallService               = "call_service"
	MsgTypeRespondService            = "respond_service"
	MsgTypePauseRequestContext       = "pause_request_context"
	MsgTypeStartRequestContext       = "start_request_context"
	MsgTypeKillRequestContext        = "kill_request_context"
	MsgTypeUpdateRequestContext      = "update_request_context"
	MsgTypeWithdrawEarnedFees        = "withdraw_earned_fees"
	//token
	MsgTypeMintToken          = "mint_token"
	MsgTypeBurnToken          = "burn_token"
	MsgTypeEditToken          = "edit_token"
	MsgTypeIssueToken         = "issue_token"
	MsgTypeSwapFeeToken       = "swap_fee_token"
	MsgTypeTransferTokenOwner = "transfer_token_owner"
	MsgTypeSwapFromERC20      = "swap_from_erc20"
	MsgTypeSwapToERC20        = "swap_to_erc20"

	MsgTypeEthereumTx = "ethereum_tx"
)

type (
	//coinswap
	MsgSwapOrder                 = coinswap.MsgSwapOrder
	MsgAddLiquidity              = coinswap.MsgAddLiquidity
	MsgRemoveLiquidity           = coinswap.MsgRemoveLiquidity
	MsgAddUnilateralLiquidity    = coinswap.MsgAddUnilateralLiquidity
	MsgRemoveUnilateralLiquidity = coinswap.MsgRemoveUnilateralLiquidity

	//farm
	MsgUnstake                      = farm.MsgUnstake
	MsgStake                        = farm.MsgStake
	MsgCreatePool                   = farm.MsgCreatePool
	MsgCreatePoolWithCommunityPool  = farm.MsgCreatePoolWithCommunityPool
	MsgDestroyPool                  = farm.MsgDestroyPool
	MsgAdjustPool                   = farm.MsgAdjustPool
	MsgHarvest                      = farm.MsgHarvest
	CommunityPoolCreateFarmProposal = farm.CommunityPoolCreateFarmProposal

	//htlc
	MsgClaimHTLC  = htlc.MsgClaimHTLC
	MsgCreateHTLC = htlc.MsgCreateHTLC

	//mt
	MsgMTMint          = mt.MsgMintMT
	MsgMTEdit          = mt.MsgEditMT
	MsgMTTransfer      = mt.MsgTransferMT
	MsgMTBurn          = mt.MsgBurnMT
	MsgMTIssueDenom    = mt.MsgIssueDenom
	MsgMTTransferDenom = mt.MsgTransferDenom

	//nft
	MsgNFTMint       = nft.MsgMintNFT
	MsgNFTEdit       = nft.MsgEditNFT
	MsgNFTTransfer   = nft.MsgTransferNFT
	MsgNFTBurn       = nft.MsgBurnNFT
	MsgIssueDenom    = nft.MsgIssueDenom
	MsgTransferDenom = nft.MsgTransferDenom

	//oracle
	MsgCreateFeed = oracle.MsgCreateFeed
	MsgEditFeed   = oracle.MsgEditFeed
	MsgPauseFeed  = oracle.MsgPauseFeed
	MsgStartFeed  = oracle.MsgStartFeed

	//random
	MsgRequestRandom = random.MsgRequestRandom

	//record
	MsgRecordCreate = record.MsgCreateRecord

	//service
	MsgDefineService         = service.MsgDefineService
	MsgBindService           = service.MsgBindService
	MsgCallService           = service.MsgCallService
	MsgRespondService        = service.MsgRespondService
	MsgUpdateServiceBinding  = service.MsgUpdateServiceBinding
	MsgSetWithdrawAddress    = service.MsgSetWithdrawAddress
	MsgDisableServiceBinding = service.MsgDisableServiceBinding
	MsgEnableServiceBinding  = service.MsgEnableServiceBinding
	MsgRefundServiceDeposit  = service.MsgRefundServiceDeposit
	MsgPauseRequestContext   = service.MsgPauseRequestContext
	MsgStartRequestContext   = service.MsgStartRequestContext
	MsgKillRequestContext    = service.MsgKillRequestContext
	MsgUpdateRequestContext  = service.MsgUpdateRequestContext
	MsgWithdrawEarnedFees    = service.MsgWithdrawEarnedFees

	//token v1beta1
	MsgIssueToken         = tokenv1beta1.MsgIssueToken
	MsgEditToken          = tokenv1beta1.MsgEditToken
	MsgBurnToken          = tokenv1beta1.MsgBurnToken
	MsgMintToken          = tokenv1beta1.MsgMintToken
	MsgTransferTokenOwner = tokenv1beta1.MsgTransferTokenOwner

	//token v1
	MsgIssueTokenV1         = tokenv1.MsgIssueToken
	MsgEditTokenV1          = tokenv1.MsgEditToken
	MsgBurnTokenV1          = tokenv1.MsgBurnToken
	MsgMintTokenV1          = tokenv1.MsgMintToken
	MsgTransferTokenOwnerV1 = tokenv1.MsgTransferTokenOwner
	MsgSwapFeeTokenV1       = tokenv1.MsgSwapFeeToken
	MsgSwapFromERC20        = tokenv1.MsgSwapFromERC20
	MsgSwapToERC20          = tokenv1.MsgSwapToERC20

	//evm
	MsgEthereumTx = evm.MsgEthereumTx
)
