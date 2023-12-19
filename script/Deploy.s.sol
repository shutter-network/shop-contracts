// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "forge-std/Script.sol";
import "../src/Inbox.sol";
import "../src/KeyperSet.sol";
import "../src/KeyperSetManager.sol";
import "../src/KeyBroadcastContract.sol";

contract DeployScript is Script {
    Inbox inbox;
    KeyperSet keyperSet;
    KeyperSetManager keyperSetManager;
    KeyBroadcastContract keyBroadcastContract;

    function setUp() public {}

    function run() public {
        uint256 deployerPrivateKey = vm.envUint("PRIVATE_KEY");
        vm.startBroadcast(deployerPrivateKey);

        address deployerAddress = vm.addr(deployerPrivateKey);
        address sequencerAddress = vm.envAddress("SEQUENCER_ADDRESS");
        console.log("deployer is:", deployerAddress);

        // Step 0: Deploy Inbox
        inbox = new Inbox(
            uint64(vm.envUint("BLOCK_GAS_LIMIT")),
            deployerAddress
        );
        inbox.initialize(deployerAddress, sequencerAddress);
        console.log("Inbox deployed at:", address(inbox));

        // Step 1: Deploy KeyperSet
        keyperSet = new KeyperSet();
        console.log("KeyperSet deployed at:", address(keyperSet));

        // Step 2: Call setKeyBroadcaster with deployer address (the sender)
        keyperSet.setKeyBroadcaster(deployerAddress);
        console.log("KeyBroadcaster set to deployer's address");

        // Finalize the KeyperSet
        keyperSet.setFinalized();
        console.log("KeyperSet finalized");

        // Step 3: Deploy KeyperSetManager and call addKeyperSet
        keyperSetManager = new KeyperSetManager(deployerAddress);
        keyperSetManager.initialize(deployerAddress, sequencerAddress);
        console.log("KeyperSetManager deployed at:", address(keyperSetManager));

        uint64 activationBlock = uint64(block.number + 10); // Assuming "near future" is 10 blocks ahead
        keyperSetManager.addKeyperSet(activationBlock, address(keyperSet));
        console.log(
            "KeyperSet added to KeyperSetManager with activation block:",
            activationBlock
        );

        // Step 4: Deploy KeyBroadcastContract and call broadcastEonKey
        keyBroadcastContract = new KeyBroadcastContract(
            address(keyperSetManager)
        );
        console.log(
            "KeyBroadcastContract deployed at:",
            address(keyBroadcastContract)
        );

        // Broadcast the eonKey
        bytes memory eonKey = vm.envBytes("EON_KEY"); // Replace with actual eon key
        keyBroadcastContract.broadcastEonKey(0, eonKey);
        console.log("Eon key broadcasted for eon 0");

        vm.stopBroadcast();
    }
}
