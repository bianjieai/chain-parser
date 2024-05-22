package msgs

import (
	coinswap "github.com/irisnet/irismod/modules/coinswap/types"
	farm "github.com/irisnet/irismod/modules/farm/types"
	htlc "github.com/irisnet/irismod/modules/htlc/types"
	mt "github.com/irisnet/irismod/modules/mt/types"
	nft "github.com/irisnet/irismod/modules/nft/types"
	oracle "github.com/irisnet/irismod/modules/oracle/types"
	random "github.com/irisnet/irismod/modules/random/types"
	record "github.com/irisnet/irismod/modules/record/types"
	service "github.com/irisnet/irismod/modules/service/types"
	token "github.com/irisnet/irismod/modules/token/types"
)

const (
	//coinswap
	MsgTypeAddLiquidity    = "add_liquidity"
	MsgTypeRemoveLiquidity = "remove_liquidity"
	MsgTypeSwapOrder       = "swap_order"
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
	MsgTypeTransferTokenOwner = "transfer_token_owner"
)

type (
	//coinswap
	MsgSwapOrder       = coinswap.MsgSwapOrder
	MsgAddLiquidity    = coinswap.MsgAddLiquidity
	MsgRemoveLiquidity = coinswap.MsgRemoveLiquidity

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

	//token
	MsgIssueToken         = token.MsgIssueToken
	MsgEditToken          = token.MsgEditToken
	MsgBurnToken          = token.MsgBurnToken
	MsgMintToken          = token.MsgMintToken
	MsgTransferTokenOwner = token.MsgTransferTokenOwner
)
