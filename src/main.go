package main

import (
	"fmt"
)

type CheckingAccount struct {
	owner       string
	agencyCode  int
	accountCode int
	balance     float64
}

func (c *CheckingAccount) withdrawAccount(value float64) string {
	allowsWithdrawal := value > 0 && value <= c.balance

	if allowsWithdrawal {
		c.balance -= value
		return "Saque realizado com sucesso"
	} else {
		return "Saldo Insuficiente"
	}
}

func (c *CheckingAccount) depositAccount(value float64) string {
	DepositsWithdrawal := value > 0 && c.accountCode != 0

	if DepositsWithdrawal {
		c.balance += value
		return "Depósito realizado com sucesso"
	} else {
		return "Valor do depósito não pode ser negativo "
	}
}

func main() {

	fmt.Println("Bem vindo ao Banco Go '\n'")

	var checkingAccount *CheckingAccount

	checkingAccount = new(CheckingAccount)

	checkingAccount.owner = "Nayara"
	checkingAccount.agencyCode = 8515
	checkingAccount.accountCode = 1919325
	checkingAccount.balance = 125.50

	fmt.Println("Você está logado na conta: ", checkingAccount.owner)
	fmt.Println("Número de agência: ", checkingAccount.agencyCode, " Número da conta: ", checkingAccount.accountCode)
	fmt.Println("Seu saldo é de R$: ", checkingAccount.balance)
	fmt.Println("")
	fmt.Println(checkingAccount.depositAccount(1000.25))
	fmt.Println("")
	fmt.Println(checkingAccount.withdrawAccount(500))
	fmt.Println("*******************************")
	fmt.Println("")
	fmt.Println("       Fim do Extrato")
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("")

	fmt.Println("Seu saldo é de R$: ", checkingAccount.balance)

}
