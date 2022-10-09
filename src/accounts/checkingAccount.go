package accounts

type CheckingAccount struct {
	Owner       string
	AgencyCode  int
	AccountCode int
	Balance     float64
}

func (c *CheckingAccount) WithdrawAccount(value float64) string {
	allowsWithdrawal := value > 0 && value <= c.Balance

	if allowsWithdrawal {
		c.Balance -= value
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *CheckingAccount) DepositAccount(value float64) (string, float64) {

	if value > 0 {
		c.Balance += value
		return "Depósito realizado com sucesso", c.Balance
	} else {
		return "Valor do depósito não pode ser negativo", c.Balance
	}
}

func (c *CheckingAccount) TransferAccount(value float64, accountTarget *CheckingAccount) bool {

	if value < c.Balance && value > 0 {
		c.Balance -= value
		accountTarget.DepositAccount(value)
		return true
	} else {
		return false
	}
}
