package pointers

import "testing"

func TestBalance(t *testing.T) {
	wallet := Wallet{Bitcoin(10)}
	expect := Bitcoin(10)
	assertBalance(t, wallet.Balance(), expect)
}

func TestDeposit(t *testing.T) {
	wallet := Wallet{Bitcoin(10)}
	wallet.Deposit(Bitcoin(5))
	expect := Bitcoin(15)
	assertBalance(t, wallet.Balance(), expect)
}

func TestWithdraw(t *testing.T) {
	t.Run("Withdraw with sufficient balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		wallet.Withdraw(Bitcoin(4))
		expect := Bitcoin(6)
		assertBalance(t, wallet.Balance(), expect)
	})

	t.Run("Withdraw with insufficient balance", func(t *testing.T) {
		wallet := Wallet{Bitcoin(10)}
		err := wallet.Withdraw(Bitcoin(12))
		assertError(t, err, ErrInsufficienctFunds)
	})
}

func assertBalance(t *testing.T, actual, expect Bitcoin) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%d', but expect '%d'", actual, expect)
	}
}

func assertError(t *testing.T, actual, expect error) {
	t.Helper()
	if actual != expect {
		t.Errorf("actual '%s', but expect '%s'", actual, expect)
	}
}
