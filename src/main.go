package main

import (
	"bank-go/src/accounts"
	"bank-go/src/owners"
	"fmt"
)

func main() {

	fmt.Println("Bem vindo ao Banco Go '\n'")

	checkingAccount := new(accounts.CheckingAccount)

	checkingAccount.Owner.Name = "Nayara Ferreira"
	checkingAccount.AgencyCode = 8515
	checkingAccount.AccountCode = 1919325
	fmt.Println(checkingAccount.DepositAccount(600))

	fmt.Println("Você está logado na conta: ", checkingAccount.Owner)
	fmt.Println("Número de agência: ", checkingAccount.AgencyCode, " Número da conta: ", checkingAccount.AccountCode)
	fmt.Println("Seu saldo é de R$: ", checkingAccount.ReturnBalance(checkingAccount))
	fmt.Println("")
	fmt.Println(checkingAccount.DepositAccount(1000.25))
	fmt.Println("")
	fmt.Println("Seu saldo é de R$: ", checkingAccount.ReturnBalance(checkingAccount))
	fmt.Println("")
	fmt.Println(checkingAccount.WithdrawAccount(500))
	fmt.Println("")
	fmt.Println("Seu saldo é de R$: ", checkingAccount.ReturnBalance(checkingAccount))
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
	fmt.Println(checkingAccount1.DepositAccount(100))

	fmt.Println("Iniciando transferência...")
	fmt.Println("Conta de Origem: ", checkingAccount.Owner.Name)
	fmt.Println("Saldo: ", checkingAccount.ReturnBalance(checkingAccount))
	fmt.Println("Conta de Destino: ", checkingAccount1.Owner.Name)
	fmt.Println("Saldo: ", checkingAccount1.ReturnBalance(checkingAccount1))

	if checkingAccount.TransferAccount(500, checkingAccount1) {
		fmt.Println("Conta de Origem: ", checkingAccount.Owner.Name)
		fmt.Println("Saldo após a transferência: ", checkingAccount.ReturnBalance(checkingAccount1))
		fmt.Println("Conta de Destino: ", checkingAccount1.Owner.Name)
		fmt.Println("Saldo após a transferência: ", checkingAccount1.ReturnBalance(checkingAccount1))
	} else {

		fmt.Println("Transferência não pode ser realizada")
	}

	fmt.Println("Criação de uma nova conta")

	checkingAccount2 := new(accounts.CheckingAccount)
	newOwner := new(owners.Owner)

	newOwner.CPF = 11002362159
	newOwner.Name = "Mariana R C"
	checkingAccount2.Owner = *newOwner
	checkingAccount2.Owner.Occupation.Code = 1
	checkingAccount2.Owner.Occupation.Description = "System Analist"
	checkingAccount2.AccountCode = 2215394
	checkingAccount2.AgencyCode = 8515

	fmt.Println("Conta criada com sucesso!", checkingAccount2)

	checkingAccount2.DepositAccount(100000)

	fmt.Println("Saldo da conta: ", checkingAccount2.ReturnBalance(checkingAccount2))

}
