package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(20))
		want := Bitcoin(20)

		assertBalance(t, wallet, want)
	})

	t.Run("Witdrawl", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Witdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()

	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s and want %s", got, want)

	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Error("wanted an error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Fatal("got an error but didn't want one")
	}
}
