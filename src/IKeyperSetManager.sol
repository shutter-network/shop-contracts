// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "openzeppelin-contracts/contracts/access/Ownable.sol";

error KeyperSetNotFinalized();
error AlreadyHaveKeyperSet();
error NoActiveKeyperSet();
error AlreadyDeactivated();

interface IKeyperSetManager {
    function addKeyperSet(
        uint64 activationSlot,
        address keyperSetContract
    ) external;

    function getNumKeyperSets() external view returns (uint64);

    function getKeyperSetIndexBySlot(
        uint64 slot
    ) external view returns (uint64);

    function getKeyperSetAddress(uint64 index) external view returns (address);

    function getKeyperSetActivationSlot(
        uint64 index
    ) external view returns (uint64);

    event KeyperSetAdded(uint64 activationSlot, address keyperSetContract);
}
