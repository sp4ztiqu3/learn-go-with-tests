package pointers

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(Bitcoin(10))

		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})
	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBal := Bitcoin(5)
		wallet := Wallet{balance: startingBal}
		err := wallet.Withdraw(Bitcoin(10))
		assertError(t, err, ErrInsufficientBalance)
		assertBalance(t, wallet, startingBal)
	})
}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got error but didn't want one")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("wanted error but did not get one")
	}

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
