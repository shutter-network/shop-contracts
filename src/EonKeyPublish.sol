// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "src/KeyBroadcastContract.sol";

contract EonKeyPublish is EonKeyPublisher {
    mapping(bytes32 => uint64) private numVotesForKey;
    mapping(address => bool) private hasVoted;
    KeyperSet keyperSet;
    KeyBroadcastContract broadcaster;
    uint64 eon;

    event EonVoteRegistered(uint64 eon, bytes key);

    constructor(address _keyperSet, address _broadcaster, uint64 _eon) {
        keyperSet = KeyperSet(_keyperSet);
        broadcaster = KeyBroadcastContract(_broadcaster);
        eon = _eon;
    }

    function eonKeyConfirmed(bytes memory eonKey) public view returns (bool) {
        uint required = keyperSet.getThreshold();
        uint64 votes = numVotesForKey[keccak256(eonKey)];
        if (votes >= required) {
            return true;
        }
        return false;
    }

    function publishEonKey(bytes memory eonKey, uint64 keyperId) public {
        if (eonKey.length == 0) {
            revert InvalidKey();
        }
        if (!keyperSet.isFinalized()) {
            revert KeyperSetNotFinalized();
        }
        if (!hasVoted[msg.sender]) {
            hasVoted[msg.sender] = true;
            numVotesForKey[keccak256(eonKey)]++;
            emit EonVoteRegistered(eon, eonKey);
        }
        // check msg.sender is keyper
        if (keyperSet.getMember(keyperId) != msg.sender) {
            revert NotAllowed();
        }
        if (eonKeyConfirmed(eonKey)) {
            // broadcast
            broadcaster.broadcastEonKey(eon, eonKey);
            //exit
            return;
        }
    }
}
