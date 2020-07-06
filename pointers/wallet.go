package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin is one of electronic coin based on black-chain
type Bitcoin int

// Wallet is package for electronic coin
type Wallet struct {
	balance Bitcoin
}

// Balance of bitcoin
func (w Wallet) Balance() Bitcoin {
	return w.balance
}

// Deposit money into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("Current balance is %s", w)
	w.balance += amount
}

func (w *Wallet) String() string {
	return fmt.Sprintf("%d BTC", w.balance)
}

// Withdraw is used for withdraw amount of money
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficienctFunds
	}
	w.balance -= amount
	return nil
}

// ErrInsufficienctFunds stands for balance is insufficent
var ErrInsufficienctFunds = errors.New("balance is not sufficient")
