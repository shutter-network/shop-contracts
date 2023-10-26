// SPDX-License-Identifier: MIT
pragma solidity ^0.8.21;

import "forge-std/Test.sol";
import "../src/Inbox.sol";

contract NonReceivable {}

contract InboxTest is Test {
    Inbox public inbox;
    Inbox.Transaction public transaction;

    function setUp() public {
        inbox = new Inbox(30e6);
        transaction = Inbox.Transaction(hex"12345678", 1e5, 0);
        vm.fee(1e9);
    }

    receive() external payable {}

    function testSubmitEncryptedTransaction() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        _submitTx(blockNumber + 1, fee, bytes4(0));
        assertEq(inbox.getTransactions(blockNumber + 1).length, 1);

        _submitTx(blockNumber + 1, fee, bytes4(0));
        assertEq(inbox.getTransactions(blockNumber + 1).length, 2);
    }

    function testSubmitEncryptedTransactionPastBlock() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        _submitTx(blockNumber, fee, BlockAlreadyFinalized.selector);
        assertEq(inbox.getTransactions(blockNumber).length, 0);
        _submitTx(blockNumber - 1, fee, BlockAlreadyFinalized.selector);
        assertEq(inbox.getTransactions(blockNumber - 1).length, 0);
    }

    function testSubmitEncryptedTransactionVariableFunds() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        _submitTx(blockNumber + 1, fee - 1, InsufficientFunds.selector);
        _submitTx(blockNumber + 1, fee, bytes4(0));
        assertEq(address(inbox).balance, fee);
        address excessFeeRecipient = address(1);
        inbox.submitEncryptedTransaction{value: fee + 1}(
            blockNumber + 1,
            transaction.encryptedTransaction,
            transaction.gasLimit,
            excessFeeRecipient
        );
        assertEq(address(inbox).balance, 2 * fee);
        assertEq(excessFeeRecipient.balance, 1);
    }

    function testClearTransactions() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        _submitTx(blockNumber + 1, fee, bytes4(0));
        _submitTx(blockNumber + 1, fee, bytes4(0));
        vm.roll(blockNumber + 1);
        assertEq(inbox.getTransactions(uint64(block.number)).length, 2);

        address sequencer = address(1);
        vm.startPrank(sequencer);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                sequencer,
                inbox.SEQUENCER_ROLE()
            )
        );
        inbox.clear();
        vm.stopPrank();

        inbox.grantRole(inbox.SEQUENCER_ROLE(), sequencer);
        vm.prank(sequencer);
        inbox.clear();
        assertEq(inbox.getTransactions(uint64(block.number)).length, 0);
    }

    function testWithdraw() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        _submitTx(blockNumber + 1, fee, bytes4(0));

        address withdrawAddress = address(1);
        vm.startPrank(withdrawAddress);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                withdrawAddress,
                inbox.WITHDRAW_ROLE()
            )
        );
        inbox.withdraw(withdrawAddress);
        vm.stopPrank();

        inbox.grantRole(inbox.WITHDRAW_ROLE(), withdrawAddress);

        uint256 balanceBefore = withdrawAddress.balance;
        uint256 inboxBalance = address(inbox).balance;
        vm.prank(withdrawAddress);
        inbox.withdraw(withdrawAddress);
        uint256 balanceAfter = withdrawAddress.balance;

        assertEq(balanceAfter, balanceBefore + inboxBalance);
        assertEq(address(inbox).balance, 0);
    }

    function testTransferFailed() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockNumber = uint64(block.number);

        NonReceivable nonReceivableContract = new NonReceivable();

        vm.startPrank(address(nonReceivableContract));
        vm.deal(address(nonReceivableContract), 1 ether);
        vm.expectRevert(TransferEtherFailed.selector);
        inbox.submitEncryptedTransaction{value: fee + 1}(
            blockNumber + 1,
            transaction.encryptedTransaction,
            transaction.gasLimit,
            address(nonReceivableContract)
        );
        vm.stopPrank();
    }

    function testBlockGasLimitExceeded() public {
        uint256 fee = block.basefee * transaction.gasLimit;
        uint64 blockToInclude = uint64(block.number) + 1;

        _submitTx(blockToInclude, fee, bytes4(0));

        // Submit a transaction to fill up to the limit exactly
        uint64 remainingGasLimit = 30e6 - transaction.gasLimit;
        inbox.submitEncryptedTransaction{
            value: remainingGasLimit * block.basefee
        }(
            blockToInclude,
            transaction.encryptedTransaction,
            remainingGasLimit,
            msg.sender
        );

        // Already 1 gas will exceed the total block limit
        vm.expectRevert(BlockGasLimitExceeded.selector);
        inbox.submitEncryptedTransaction{value: 1 * block.basefee}(
            blockToInclude,
            transaction.encryptedTransaction,
            1,
            msg.sender
        );
    }

    function testSetBlockGasLimit() public {
        uint64 oldBlockGasLimit = inbox.getBlockGasLimit();
        uint64 newBlockGasLimit = oldBlockGasLimit + 1;
        address blockGasLimitSetter = address(1);

        vm.startPrank(blockGasLimitSetter);
        vm.expectRevert(
            abi.encodeWithSelector(
                IAccessControl.AccessControlUnauthorizedAccount.selector,
                blockGasLimitSetter,
                inbox.BLOCK_GAS_LIMIT_SETTER_ROLE()
            )
        );
        inbox.setBlockGasLimit(newBlockGasLimit);
        vm.stopPrank();
        assertEq(oldBlockGasLimit, inbox.getBlockGasLimit());

        inbox.grantRole(
            inbox.BLOCK_GAS_LIMIT_SETTER_ROLE(),
            blockGasLimitSetter
        );

        vm.prank(blockGasLimitSetter);
        inbox.setBlockGasLimit(newBlockGasLimit);
        assertEq(inbox.getBlockGasLimit(), newBlockGasLimit);
    }

    function _submitTx(
        uint64 blockNumber,
        uint256 value,
        bytes4 revertSelector
    ) private {
        uint256 expectedLength = inbox.getTransactions(blockNumber).length + 1;
        bytes memory data = bytes.concat(
            abi.encode(expectedLength - 1),
            transaction.encryptedTransaction
        );

        if (revertSelector != bytes4(0)) {
            vm.expectRevert(revertSelector);
            expectedLength -= 1;
        }
        inbox.submitEncryptedTransaction{value: value}(
            blockNumber,
            data,
            transaction.gasLimit,
            msg.sender
        );
        assertEq(inbox.getTransactions(blockNumber).length, expectedLength);
        if (revertSelector != bytes4(0)) {
            return;
        }
        assertEq(
            inbox
            .getTransactions(blockNumber)[expectedLength - 1]
                .encryptedTransaction,
            data
        );
    }
}
