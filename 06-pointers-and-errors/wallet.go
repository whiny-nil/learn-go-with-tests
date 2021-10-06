package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amt Bitcoin) {
	w.balance += amt
}

var ErrorInsufficientFunds = errors.New("withdrawal amount exceeds balance")

func (w *Wallet) Withdraw(amt Bitcoin) error {
	if amt > w.balance {
		return ErrorInsufficientFunds
	}

	w.balance -= amt
	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
