package pointers

import "testing"

func TestBalance(t *testing.T) {
	wallet := Wallet{10}
	expect := 10
	assertEquals(t, wallet.Balance(), expect)
}

func TestDeposit(t *testing.T) {
	wallet := Wallet{10}
	wallet.Deposit(5)
	expect := 15
	assertEquals(t, wallet.Balance(), expect)
}

func assertEquals(t *testing.T, actual, expect int) {
	if actual != expect {
		t.Errorf("actual '%d', but expect '%d'", actual, expect)
	}
}
