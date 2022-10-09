package accounts

import (
	"bank-go/src/owners"
)

type CheckingAccount struct {
	Owner       owners.Owner
	AgencyCode  int
	AccountCode int
	balance     float64
}

func (c *CheckingAccount) WithdrawAccount(value float64) string {
	allowsWithdrawal := value > 0 && value <= c.balance

	if allowsWithdrawal {
		c.balance -= value
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *CheckingAccount) DepositAccount(value float64) (string, float64) {

	if value > 0 {
		c.balance += value
		return "Depósito realizado com sucesso", c.balance
	} else {
		return "Valor do depósito não pode ser negativo", c.balance
	}
}

func (c *CheckingAccount) TransferAccount(value float64, accountTarget *CheckingAccount) bool {

	if value < c.balance && value > 0 {
		c.balance -= value
		accountTarget.DepositAccount(value)
		return true
	} else {
		return false
	}
}

func (c *CheckingAccount) ReturnBalance(accountTarget *CheckingAccount) float64 {
	return c.balance
}
