// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "../src/Inbox.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyBroadcastContract.sol";

contract DeployScript is Script {
    Inbox inbox;
    KeyperSetManager keyperSetManager;
    KeyBroadcastContract keyBroadcastContract;

    KeyperSet keyperSet;

    address inboxAddress;
    address keyBroadcastContractAddress;
    address keyperSetManagerAddress;

    function setUp() public {}

    function run() public {
        uint256 adminPrivateKey = vm.envUint("ADMIN_PRIVATE_KEY");
        uint256 activationDelta = vm.envUint("ACTIVATION_DELTA");
        uint256 threshold = vm.envUint("THRESHOLD");
        address[] memory keypers = vm.envAddress("KEYPER_ADDRESSES", ",");
        if (keypers.length == 0) {
            //TODO: error
        }
        //TODO: sanity check threshold

        inboxAddress = vm.envAddress("INBOX_ADDRESS");
        keyperSetManagerAddress = vm.envAddress("KEYPERSETMANAGER_ADDRESS");
        keyBroadcastContractAddress = vm.envAddress("KEYBROADCAST_ADDRESS");

        inbox = Inbox(inboxAddress);
        keyperSetManager = KeyperSetManager(keyperSetManagerAddress);
        keyBroadcastContract = KeyBroadcastContract(
            keyBroadcastContractAddress
        );

        vm.startBroadcast(adminPrivateKey);

        address adminAddress = vm.addr(adminPrivateKey);
        console.log("admin is:", adminAddress);

        keyperSet = new KeyperSet();
        console.log("KeyperSet deployed at:", address(keyperSet));
        keyperSet.addMembers(keypers);
        keyperSet.setThreshold(uint64(threshold));

        //FIXME:  only one predetermined address can set the key.
        // We want to do this automatically:
        // Ideally all keypers should call to set it, and when the threshold
        // is reached, the key is "activated" (returned by a read, event emitted)
        keyperSet.setKeyBroadcaster(keypers[0]);
        console.log("KeyBroadcaster is:", address(keypers[0]));

        // Finalize the KeyperSet
        keyperSet.setFinalized();
        console.log("KeyperSet finalized");

        uint64 activationBlock = uint64(block.number + activationDelta);
        keyperSetManager.addKeyperSet(activationBlock, address(keyperSet));
        console.log(
            "KeyperSet added to KeyperSetManager with activation block:",
            activationBlock
        );

        vm.stopBroadcast();
    }
}
