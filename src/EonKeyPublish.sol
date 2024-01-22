// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "src/KeyBroadcastContract.sol";

contract EonKeyPublish is EonKeyPublisher {
    mapping(uint64 => bytes) private keys;
    uint64[] publishers;
    KeyperSet keyperSet;
    KeyBroadcastContract broadcaster;
    uint64 eon;

    constructor(
        KeyperSet _keyperSet,
        KeyBroadcastContract _broadcaster,
        uint64 _eon
    ) {
        keyperSet = KeyperSet(_keyperSet);
        broadcaster = KeyBroadcastContract(_broadcaster);
        eon = _eon;
    }

    function eonKeyConfirmed(bytes memory eonKey) public view returns (bool) {
        uint required = keyperSet.getThreshold();
        if (publishers.length >= keyperSet.getThreshold()) {
            for (uint i = 0; i < publishers.length; i++) {
                if (keccak256(keys[publishers[i]]) == keccak256(eonKey)) {
                    required--;
                    if (required == 0) {
                        break;
                    }
                }
            }
        }
        return required <= 0;
    }

    function publishEonKey(bytes memory eonKey) public {
        if (eonKey.length == 0) {
            revert InvalidKey();
        }
        if (!keyperSet.isFinalized()) {
            revert KeyperSetNotFinalized();
        }
        uint64 keyperId = uint64(keyperSet.keyperIndex(msg.sender));
        if (keyperId < 0) {
            revert NotAllowed();
        }
        publishers.push(keyperId);
        if (eonKeyConfirmed(eonKey)) {
            // broadcast
            broadcaster.broadcastEonKey(eon, eonKey);
            //exit
            return;
        }
        // only store key, if not yet broadcast
        keys[keyperId] = eonKey;
    }
}
