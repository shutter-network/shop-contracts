// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "./RestrictedPausable.sol";

error InsufficientFunds();
error BlockAlreadyFinalized();
error TransferEtherFailed();

contract Inbox is RestrictedPausable {
    struct Transaction {
        bytes encryptedTransaction;
        uint64 gasLimit;
    }

    event EncryptedTransactionSubmitted(
        uint64 block,
        bytes encryptedTransaction,
        uint64 gasLimit,
        uint256 fee
    );

    event FeesWithdrawn(address recipient);

    bytes32 public constant WITHDRAW_ROLE = keccak256("WITHDRAW_ROLE");
    bytes32 public constant SEQUENCER_ROLE = keccak256("SEQUENCER_ROLE");

    mapping(uint64 blockNumber => Transaction[]) private transactions;

    function submitEncryptedTransaction(
        uint64 blockNumber,
        bytes memory encryptedTransaction,
        uint64 gasLimit,
        address excessFeeRecipient
    ) external payable whenNotPaused {
        if (blockNumber <= uint64(block.number)) {
            revert BlockAlreadyFinalized();
        }
        uint256 fee = gasLimit * block.basefee;
        if (msg.value < fee) {
            revert InsufficientFunds();
        }

        Transaction[] storage transactionsForBlock = transactions[blockNumber];
        transactionsForBlock.push(Transaction(encryptedTransaction, gasLimit));

        emit EncryptedTransactionSubmitted(
            blockNumber,
            encryptedTransaction,
            gasLimit,
            fee
        );

        uint256 excessValue = msg.value - fee;
        if (excessValue > 0) {
            _sendEther(excessFeeRecipient, excessValue);
        }
    }

    function getTransactions(
        uint64 blockNumber
    ) external view returns (Transaction[] memory) {
        return transactions[blockNumber];
    }

    function clear() external onlyRole(SEQUENCER_ROLE) {
        delete transactions[uint64(block.number)];
    }

    function withdraw(address recipient) external onlyRole(WITHDRAW_ROLE) {
        _sendEther(recipient, address(this).balance);
        emit FeesWithdrawn(recipient);
    }

    function _sendEther(address recipient, uint256 value) private {
        (bool sent, ) = recipient.call{value: value}("");
        if (!sent) {
            revert TransferEtherFailed();
        }
    }
}
