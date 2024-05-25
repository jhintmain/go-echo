package account

import "errors"

type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("Can't withraw money")

// Deposit : 입급
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// Balance : 잔액
func (account Account) Balance() int {
	return account.balance
}

// Withraw : 출금
func (account *Account) Withraw(amount int) error {
	if account.balance-amount < 0 {
		return errNoMoney
	}
	account.balance -= amount
	return nil
}

// NewAccount : 새로운 계좌를 생성하는 func

func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}
