package ptranderror

import (
	"errors"
	"fmt"
)

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

var moreError = errors.New("withdraw more")

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return moreError
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
