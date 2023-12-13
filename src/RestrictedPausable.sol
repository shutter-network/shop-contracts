// SPDX-License-Identifier: MIT
pragma solidity ^0.8.22;

import "openzeppelin/contracts/access/AccessControl.sol";
import "openzeppelin/contracts/utils/Pausable.sol";

abstract contract RestrictedPausable is AccessControl, Pausable {
    bytes32 public constant PAUSER_ROLE = keccak256("PAUSER_ROLE");

    constructor(address admin, address pauser) {
        _grantRole(DEFAULT_ADMIN_ROLE, admin);
        _grantRole(PAUSER_ROLE, pauser);
    }

    function pause() external onlyRole(PAUSER_ROLE) {
        _pause();
    }

    function unpause() external onlyRole(DEFAULT_ADMIN_ROLE) {
        _unpause();
    }
}
