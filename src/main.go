package main

import "fmt"

type CheckingAccount struct {
	owner       string
	agencyCode  int
	accountCode int
	balance     float64
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
	fmt.Println("Seu saldo hoje é de R$: ", checkingAccount.balance)
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("Fim do Extrato")
	fmt.Println("*******************************")

	checkingAccount = new(CheckingAccount)

	checkingAccount.owner = "Brownie"
	checkingAccount.agencyCode = 8515
	checkingAccount.accountCode = 12369852
	checkingAccount.balance = -29.37

	fmt.Println("Você está logado na conta: ", checkingAccount.owner)
	fmt.Println("Número de agência: ", checkingAccount.agencyCode, " Número da conta: ", checkingAccount.accountCode)
	fmt.Println("Seu saldo hoje é de R$: ", checkingAccount.balance)
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("Fim do Extrato")
	fmt.Println("*******************************")

	checkingAccount1 := CheckingAccount{
		owner:       "Mariana",
		agencyCode:  8515,
		accountCode: 19859698,
		balance:     5000.47,
	}

	fmt.Println("Você está logado na conta: ", checkingAccount1.owner)
	fmt.Println("Número de agência: ", checkingAccount1.agencyCode, " Número da conta: ", checkingAccount1.accountCode)
	fmt.Println("Seu saldo hoje é de R$: ", checkingAccount1.balance)
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("Fim do Extrato")
	fmt.Println("*******************************")

}
