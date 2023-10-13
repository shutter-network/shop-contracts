// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "src/IKeyperSet.sol";
import "src/IKeyperSetManager.sol";
import "./KeyperSet.sol";
import "openzeppelin/contracts/access/AccessControl.sol";
import "openzeppelin/contracts/utils/Pausable.sol";

contract KeyperSetManager is IKeyperSetManager, AccessControl, Pausable {
    struct KeyperSetData {
        uint64 activationSlot;
        address contractAddress;
    }

    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    KeyperSetData[] public keyperSets;

    constructor() {
        _grantRole(DEFAULT_ADMIN_ROLE, msg.sender);
    }

    function addKeyperSet(
        uint64 activationSlot,
        address keyperSetContract
    ) external onlyRole(DEFAULT_ADMIN_ROLE) {
        if (
            keyperSets.length > 0 &&
            activationSlot < keyperSets[keyperSets.length - 1].activationSlot
        ) {
            revert AlreadyHaveKeyperSet();
        }
        if (!IKeyperSet(keyperSetContract).isFinalized()) {
            revert KeyperSetNotFinalized();
        }
        keyperSets.push(KeyperSetData(activationSlot, keyperSetContract));
        emit KeyperSetAdded(activationSlot, keyperSetContract);
    }

    function getNumKeyperSets() external view returns (uint64) {
        return uint64(keyperSets.length);
    }

    function getKeyperSetIndexBySlot(
        uint64 slot
    ) external view returns (uint64) {
        for (uint256 i = keyperSets.length; i > 0; i--) {
            if (keyperSets[i - 1].activationSlot <= slot) {
                return uint64(i - 1);
            }
        }
        revert NoActiveKeyperSet();
    }

    function getKeyperSetActivationSlot(
        uint64 index
    ) external view returns (uint64) {
        return keyperSets[index].activationSlot;
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
