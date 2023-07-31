package pointers

import "testing"

func TestWallet(t *testing.T) {
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

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(5)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(6))

		assertError(err, ErrInsufficientFunds, t)
		assertBalance(wallet, startingBalance, t)
	})
}

func assertBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but didn't get one")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
