package main

import "fmt"

type Account struct {
	AccountNumber int
	HolderName    string
	Balance       float64
}

func (A *Account) Deposit(amount float64) {
	A.Balance = A.Balance + amount
	fmt.Printf("BDT : %f\n is deposited, Your current account balance is : %f\n", amount, A.Balance)
}

func (A *Account) Withdraw(amount float64) {
	if A.Balance < amount {
		fmt.Print("Inefficient balance")
	} else {
		A.Balance = A.Balance - amount
		fmt.Printf("BDT : %f\n is withdrawer, Your current account balance is : %f\n", amount, A.Balance)
	}
}

func main() {
	newA1 := Account{12, "Akash", 13000.4}
	fmt.Println(newA1)
	newA1.Withdraw(12000)
	fmt.Println(newA1)
	newA1.Deposit(10000)
	fmt.Println(newA1)
}
