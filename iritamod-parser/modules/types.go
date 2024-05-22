package msgs

import (
	"github.com/bianjieai/iritamod/modules/identity"
	"github.com/bianjieai/iritamod/modules/node"
	"github.com/bianjieai/iritamod/modules/perm"
	sidechain "github.com/bianjieai/iritamod/modules/side-chain"
	iritaslashing "github.com/bianjieai/iritamod/modules/slashing"
)

const (
	MsgTypeUnjailValidator = "unjail_validator"

	MsgTypeUpdateIdentity = "update_identity"
	MsgTypeCreateIdentity = "create_identity"

	DocTypeAssignRoles   = "assign_roles"
	DocTypeUnassignRoles = "unassign_roles"
	DocTypeBlockAccount  = "block_account"
	DocTypeUnlockAccount = "unblock_account"

	MsgTypeCreateValidator = "create_validator" // type for MsgCreateValidator
	MsgTypeUpdateValidator = "update_validator" // type for MsgUpdateValidator
	MsgTypeRemoveValidator = "remove_validator" // type for MsgRemoveValidator
	MsgTypeGrantNode       = "grant_node"       // type for MsgGrantNode
	MsgTypeRevokeNode      = "revoke_node"      // type for MsgRevokeNode

	MsgTypeCreateBlockHeader = "create_block_header"
	MsgTypeCreateSpace       = "create_space"
	MsgTypeTransferSpace     = "transfer_space"
)

type (
	MsgCreateIdentity = identity.MsgCreateIdentity
	MsgUpdateIdentity = identity.MsgUpdateIdentity

	MsgNodeCreate = node.MsgCreateValidator
	MsgNodeUpdate = node.MsgUpdateValidator
	MsgNodeRemove = node.MsgRemoveValidator
	MsgNodeGrant  = node.MsgGrantNode
	MsgNodeRevoke = node.MsgRevokeNode

	MsgAssignRoles    = perm.MsgAssignRoles
	MsgUnassignRoles  = perm.MsgUnassignRoles
	MsgBlockAccount   = perm.MsgBlockAccount
	MsgUnblockAccount = perm.MsgUnblockAccount

	MsgUnjailValidator = iritaslashing.MsgUnjailValidator

	MsgCreateBlockHeader = sidechain.MsgCreateBlockHeader
	MsgCreateSpace       = sidechain.MsgCreateSpace
	MsgTransferSpace     = sidechain.MsgTransferSpace
)
