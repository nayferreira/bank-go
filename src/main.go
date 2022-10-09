package main

import (
	"bank-go/src/accounts"
	"fmt"
)

func main() {

	fmt.Println("Bem vindo ao Banco Go '\n'")

	checkingAccount := new(accounts.CheckingAccount)

	checkingAccount.Owner.Name = "Nayara Ferreira"
	checkingAccount.AgencyCode = 8515
	checkingAccount.AccountCode = 1919325
	checkingAccount.Balance = 125.50

	fmt.Println("Você está logado na conta: ", checkingAccount.Owner)
	fmt.Println("Número de agência: ", checkingAccount.AgencyCode, " Número da conta: ", checkingAccount.AccountCode)
	fmt.Println("Seu saldo é de R$: ", checkingAccount.Balance)
	fmt.Println("")
	fmt.Println(checkingAccount.DepositAccount(1000.25))
	fmt.Println("")
	fmt.Println("Seu saldo é de R$: ", checkingAccount.Balance)
	fmt.Println("")
	fmt.Println(checkingAccount.WithdrawAccount(500))
	fmt.Println("")
	fmt.Println("Seu saldo é de R$: ", checkingAccount.Balance)
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("")
	fmt.Println("       Fim do Extrato")
	fmt.Println("")
	fmt.Println("*******************************")
	fmt.Println("")

	checkingAccount1 := new(accounts.CheckingAccount)

	checkingAccount1.Owner.Name = "Brownie"
	checkingAccount1.AgencyCode = 8515
	checkingAccount1.AccountCode = 125963
	checkingAccount1.Balance = 0.01

	fmt.Println("Iniciando transferência...")
	fmt.Println("Conta de Origem: ", checkingAccount.Owner.Name)
	fmt.Println("Saldo: ", checkingAccount.Balance)
	fmt.Println("Conta de Destino: ", checkingAccount1.Owner.Name)
	fmt.Println("Saldo: ", checkingAccount1.Balance)

	if checkingAccount.TransferAccount(500, checkingAccount1) {
		fmt.Println("Conta de Origem: ", checkingAccount.Owner.Name)
		fmt.Println("Saldo após a transferência: ", checkingAccount.Balance)
		fmt.Println("Conta de Destino: ", checkingAccount1.Owner)
		fmt.Println("Saldo após a transferência: ", checkingAccount1.Balance)
	} else {

		fmt.Println("Transferência não pode ser realizada")
	}

}
