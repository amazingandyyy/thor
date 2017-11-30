package account

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/state"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/vechain/vecore/acc"
)

type Account struct {
	address      acc.Address
	addrHash     common.Hash // hash of ethereum address of the account
	account      acc.Account
	storage      state.Storage
	dirtyStorage state.Storage
	deleted      bool
}

func newAccount(addr acc.Address, acc acc.Account) *Account {
	return &Account{
		address:      addr,
		addrHash:     crypto.Keccak256Hash(addr[:]),
		account:      acc,
		storage:      make(state.Storage),
		dirtyStorage: make(state.Storage),
	}
}

func (c *Account) setBalance(amount *big.Int) {
	c.account.Balance = amount
}

func (c *Account) getBalance() *big.Int {
	return c.account.Balance
}

// // addBalance add amount to c's balance.
// // It is used to add funds to the destination account of a transfer.
// // :rtype: bool, if changed return true, or return false
// func (c *Account) addBalance(amount *big.Int) bool {
// 	if amount.Sign() == 0 {
// 		return false
// 	}
// 	newBalance := new(big.Int).Add(c.account.Balance, amount)
// 	c.setBalance(newBalance)
// 	return true
// }
