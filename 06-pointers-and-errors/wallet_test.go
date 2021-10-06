package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		wallet.Deposit(Bitcoin(5))

		assertBalance(t, wallet, Bitcoin(15))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(10))

		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(5))
	})

	t.Run("withdraw too much", func(t *testing.T) {
		wallet := Wallet{Bitcoin(15)}
		err := wallet.Withdraw(Bitcoin(20))

		assertError(t, ErrorInsufficientFunds, err)
		assertBalance(t, wallet, Bitcoin(15))
	})
}

func assertBalance(t testing.TB, wallet Wallet, expected Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != expected {
		t.Errorf("expected %s, got %s", expected, got)
	}
}

func assertError(t testing.TB, expected error, got error) {
	t.Helper()
	if got == nil {
		t.Fatal("expected error, got nil")
	}

	if got != expected {
		t.Errorf("expected %g, got %g", expected, got)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("expected no error, got %g", got)
	}
}
