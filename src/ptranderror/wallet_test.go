package ptranderror

import "testing"

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("want an error got nil")
		}
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Errorf("excepted nil got an error")
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin((10)))
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin((10)))
		assertNoError(t, err)
	})

	t.Run("withdraw more", func(t *testing.T) {
		w := Wallet{balance: Bitcoin(20)}
		err := w.Withdraw(Bitcoin(100))

		assertBalance(t, w, Bitcoin(20))
		assertError(t, err, moreError)
	})
}
