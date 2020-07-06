package pointers

// Wallet is a electronic coin based on black-chain
type Wallet struct {
	balance int
}

// Balance of bitcoin
func (w Wallet) Balance() int {
	return w.balance
}

// Deposit money into wallet
func (w *Wallet) Deposit(money int) {
	w.balance += money
}
