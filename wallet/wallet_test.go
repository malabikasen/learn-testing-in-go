package wallet

import "testing"

func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(Bitcoin(10.0))
	got := wallet.Balance()
	var want Bitcoin
	want = 10.0

	if got != want {
		t.Errorf("got %.2f, wanted %.2f", got, want)
	}
}
