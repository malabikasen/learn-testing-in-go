package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10.0)
	got := wallet.Balance()
	want := 10.0

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got, want)
	}
}