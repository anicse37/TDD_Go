package Pointerserrors_test

import (
	"testing"

	Pointerserrors "github.com/anicse37/TDD_Go/Pointers_Errors"
)

func TestBank(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {

		account := Pointerserrors.Bank{}
		account.Deposit(100)
		got := account.Display()
		var want Pointerserrors.Money = 100
		if got != want {
			t.Errorf("got %s | want %s ", got, want)
		}
	})
	t.Run("Withdraw", func(t *testing.T) {
		account := Pointerserrors.Bank{}
		account.Deposit(100)
		account.Withdraw(50)
		got := account.Display()
		var want Pointerserrors.Money = 50
		if got != want {
			t.Errorf("got %s | want %s ", got, want)
		}
	})
}
