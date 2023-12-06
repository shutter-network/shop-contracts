// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package predeploy

import "github.com/ethereum/go-ethereum/common"

var (
	inbox                    = "0x4200420000000000000000000000000000000000"
	InboxAddr                = common.HexToAddress(inbox)
	keyperSetManager         = "0x4200420000000000000000000000000000000001"
	KeyperSetManagerAddr     = common.HexToAddress(keyperSetManager)
	keyBroadcastContract     = "0x4200420000000000000000000000000000000002"
	KeyBroadcastContractAddr = common.HexToAddress(keyBroadcastContract)

	Predeploys = make(map[string]*common.Address)
)

func init() {

	Predeploys["Inbox"] = &InboxAddr
	Predeploys["KeyperSetManager"] = &KeyperSetManagerAddr
	Predeploys["KeyBroadcastContract"] = &KeyBroadcastContractAddr
}
