// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "openzeppelin/contracts/utils/math/Math.sol";
import "./KeyperSet.sol";
import "./RestrictedPausable.sol";

error KeyperSetNotFinalized();
error AlreadyHaveKeyperSet();
error NoActiveKeyperSet();
error AlreadyDeactivated();

contract KeyperSetManager is RestrictedPausable {
    struct KeyperSetData {
        uint64 activationBlock;
        address contractAddress;
    }

    constructor(
        address _adminRoleAddress
    ) RestrictedPausable(_adminRoleAddress) {}

    KeyperSetData[] private keyperSets;

    event KeyperSetAdded(
        uint64 activationBlock,
        address keyperSetContract,
        address[] members,
        uint64 threshold,
        uint64 eon
    );

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
        KeyperSet keyperSet = KeyperSet(keyperSetContract);
        emit KeyperSetAdded(
            activationBlock,
            keyperSetContract,
            keyperSet.getMembers(),
            keyperSet.getThreshold(),
            uint64(keyperSets.length - 1)
        );
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
}
