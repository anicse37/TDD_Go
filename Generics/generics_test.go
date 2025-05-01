package generics_test

import (
	"testing"

	generics "github.com/anicse37/TDD_Go/Generics"
)

/*-------------------------------------------------------*/
func TestAssertFunctions(t *testing.T) {
	t.Run("asserting on integers", func(t *testing.T) {
		AssertEqual(t, 1, 1)
		AssertNotEqual(t, 1, 2)
	})
	t.Run("asserting on strings", func(t *testing.T) {
		AssertEqual(t, "Hello", "Hello")
		AssertNotEqual(t, "Hello", "World")
	})
}

/*-------------------------------------------------------*/
func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		myStackOfInts := new(generics.StackOfInts)

		AssertTrue(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(19)
		AssertFalse(t, myStackOfInts.IsEmpty())

		myStackOfInts.Push(84)
		value, _ := myStackOfInts.Pop()
		AssertEqual(t, value, 84)
		value, _ = myStackOfInts.Pop()
		AssertEqual(t, value, 19)
		AssertTrue(t, myStackOfInts.IsEmpty())
	})
	/*---------------------------------------------------*/

	t.Run("Stack of Strings", func(t *testing.T) {
		myStackOfStrings := new(generics.StackOfStrings)

		AssertTrue(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("19")
		AssertFalse(t, myStackOfStrings.IsEmpty())

		myStackOfStrings.Push("84")
		value, _ := myStackOfStrings.Pop()
		AssertEqual(t, value, "84")
		value, _ = myStackOfStrings.Pop()
		AssertEqual(t, value, "19")
		AssertTrue(t, myStackOfStrings.IsEmpty())
	})
	/*---------------------------------------------------*/
	t.Run("Using Stack[int]", func(t *testing.T) {
		myStackOfInts := generics.NewStack[int]()

		myStackOfInts.Push(19)
		myStackOfInts.Push(84)
		numb1, _ := myStackOfInts.Pop()
		numb2, _ := myStackOfInts.Pop()

		AssertEqual(t, numb1+numb2, 103)
	})
	/*---------------------------------------------------*/
	t.Run("Using Stack[string]", func(t *testing.T) {
		myStackOfInts := generics.NewStack[string]()

		myStackOfInts.Push("19")
		myStackOfInts.Push("84")
		numb1, _ := myStackOfInts.Pop()
		numb2, _ := myStackOfInts.Pop()

		AssertEqual(t, numb1+numb2, "8419")
	})
	/*---------------------------------------------------*/
}

/*-------------------------------------------------------*/
func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v", got, want)
	}
}
func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Error("Do not want: ", got)
	}
}

/*-------------------------------------------------------*/
func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("g=Got %v || Want True", got)
	}
}
func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("g=Got %v || Want True", got)
	}
}

/*-------------------------------------------------------*/
