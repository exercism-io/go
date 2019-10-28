// Package account includes a solution for the "Bank Account" problem in the Go track on https://exercism.io.
package account

import "sync"

// Account is a bank account.
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open opens a new account with an initial deposit.
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	account := Account{
		balance: initialDeposit,
		closed:  false,
	}
	return &account
}

// Close closes an account.
func (account *Account) Close() (payout int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if account.closed {
		return 0, false
	}
	account.closed = true
	return account.balance, true
}

// Balance returns the current balance.
func (account *Account) Balance() (balance int64, ok bool) {
	account.RLock()
	defer account.RUnlock()
	if account.closed {
		return 0, false
	}
	return account.balance, true
}

// Deposit deposits/withdraw money to/from the account.
func (account *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	account.Lock()
	defer account.Unlock()
	if account.closed {
		return 0, false
	}
	newBalance = account.balance + amount
	if newBalance < 0 {
		return account.balance, false
	}
	account.balance = newBalance
	return newBalance, true
}
