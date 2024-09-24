package bank

import "fmt"

type BankAccount struct {
	name          string
	accountNumber string
	balance       float64
}

func NewAccount() BankAccount {
	return BankAccount{
		name:          "Allan Kamau",
		accountNumber: "1234567890",
		balance:       1000.0,
	}
}

func (b *BankAccount) Credit(amount float64) {
	if amount > b.balance {
		fmt.Println("insufficient balance.")
		return
	}
	b.balance -= amount
	fmt.Printf("New balance after credit: %.2f\n", b.balance)
}

func (b *BankAccount) Debit(amount float64) {
	b.balance += amount
	fmt.Printf("New balnce after debit: %.2f\n", b.balance)
}
