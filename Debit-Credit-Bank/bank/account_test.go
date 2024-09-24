package bank

import "testing"

func TestCredit(t *testing.T) {
	account := NewAccount()

	account.Credit(500)
	expectedBalance := 500.0

	if account.balance != expectedBalance {
		t.Errorf("credit(600) failed, expected balance: %.2f, got: %.2f", expectedBalance, account.balance)
	}

	account.Credit(600)
	expectedBalance = 500.0

	if account.balance != expectedBalance {
		t.Errorf("credit(600) failed, expected balance: %.2f, got: %.2f", expectedBalance, account.balance)
	}

}

func TestDebit(t *testing.T) {
	account := NewAccount()

	 account.Debit(500)
	expectedBalance := 1500.0

	if expectedBalance != account.balance {
		t.Errorf("Debit(500) failed, expected balance: %.2f, got: %.2f ", expectedBalance, account.balance)
	}
}