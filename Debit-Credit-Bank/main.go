package main

import (
	"bank/bank"
	"fmt"
	"os"
	"strconv"
)

func main() {
	account := bank.NewAccount()

	if len(os.Args) < 3 {
		fmt.Println(("usage: go run . [credit|debit] [amount]"))
		return
	}

	operation := os.Args[1]
	amountStr := os.Args[2]

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		fmt.Println("Invalid amount.Please enter a valid number")
		return
	}

	switch operation {
	case "credit":
		account.Credit(amount)
	case "debit":
		account.Debit(amount)
	default:
		fmt.Println("Unkown operation. use credit or debit .")
	}
}