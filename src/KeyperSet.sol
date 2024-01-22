// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "openzeppelin/contracts/access/Ownable.sol";

error AlreadyFinalized();
error NotRegistered();

contract KeyperSet is Ownable {
    bool finalized;
    address[] members;
    uint64 threshold;
    address broadcaster;
    address publisher;

    constructor() Ownable(msg.sender) {}

    function isFinalized() external view returns (bool) {
        return finalized;
    }

    function getNumMembers() external view returns (uint64) {
        return uint64(members.length);
    }

    function getMember(uint64 index) external view returns (address) {
        return members[index];
    }

    function getMembers() external view returns (address[] memory) {
        return members;
    }

    function getThreshold() external view returns (uint64) {
        return threshold;
    }

    function getPublisher() external view returns (address) {
        return publisher;
    }

    function addMembers(address[] calldata newMembers) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        for (uint64 j = 0; j < newMembers.length; j++) {
            members.push(newMembers[j]);
        }
    }

    function setThreshold(uint64 _threshold) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        threshold = _threshold;
    }

    function setKeyBroadcaster(address _broadcaster) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        broadcaster = _broadcaster;
    }

    function setPublisher(address _publisher) public onlyOwner {
        if (finalized) {
            revert AlreadyFinalized();
        }
        publisher = _publisher;
    }

    function setFinalized() public onlyOwner {
        finalized = true;
    }

    function keyperIndex(address candidate) public view returns (int64) {
        for (uint64 i = 0; i < members.length; i++) {
            if (members[i] == candidate) {
                return int64(i);
            }
        }
        return -1;
    }

    function isAllowedToBroadcastEonKey(
        address a
    ) external view returns (bool) {
        return a == publisher;
    }
}
