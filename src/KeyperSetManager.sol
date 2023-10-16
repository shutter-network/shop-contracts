// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "./KeyperSet.sol";
import "openzeppelin/contracts/access/AccessControl.sol";
import "openzeppelin/contracts/utils/Pausable.sol";
import "openzeppelin/contracts/utils/math/Math.sol";

error KeyperSetNotFinalized();
error AlreadyHaveKeyperSet();
error NoActiveKeyperSet();
error AlreadyDeactivated();

contract KeyperSetManager is AccessControl, Pausable {
    struct KeyperSetData {
        uint64 activationBlock;
        address contractAddress;
    }

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    KeyperSetData[] private keyperSets;

    event KeyperSetAdded(uint64 activationBlock, address keyperSetContract);

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function addKeyperSet(
        uint64 activationBlock,
        address keyperSetContract
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (
            keyperSets.length > 0 &&
            activationBlock <
            Math.max(
                keyperSets[keyperSets.length - 1].activationBlock,
                block.number + 1
            )
        ) {
            revert AlreadyHaveKeyperSet();
        }
        if (!KeyperSet(keyperSetContract).isFinalized()) {
            revert KeyperSetNotFinalized();
        }
        keyperSets.push(KeyperSetData(activationBlock, keyperSetContract));
        emit KeyperSetAdded(activationBlock, keyperSetContract);
    }

    function getNumKeyperSets() external view returns (uint64) {
        return uint64(keyperSets.length);
    }

    function getKeyperSetIndexByBlock(
        uint64 blockNumber
    ) external view returns (uint64) {
        for (uint256 i = keyperSets.length; i > 0; i--) {
            if (keyperSets[i - 1].activationBlock <= blockNumber) {
                return uint64(i - 1);
            }
        }
        revert NoActiveKeyperSet();
    }

    function getKeyperSetActivationBlock(
        uint64 index
    ) external view returns (uint64) {
        return keyperSets[index].activationBlock;
    }

    function getKeyperSetAddress(uint64 index) external view returns (address) {
        return keyperSets[index].contractAddress;
    }

    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }
}
