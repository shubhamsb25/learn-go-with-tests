package pointers

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(wallet Wallet, want Bitcoin, t *testing.T) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(wallet, Bitcoin(10), t)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Withdraw(Bitcoin(5))
		assertBalance(wallet, Bitcoin(5), t)
	})
}
