package sums

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	expected := 6
	got := Sum(numbers)

	if got != expected {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAll(t *testing.T) {
	expected := []int{3, 9}
	got := SumAll([]int{1, 2}, []int{9, 0})

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected %d, got %d", expected, got)
	}
}

func TestSumAllTails(t *testing.T) {
	t.Run("it adds the tails of given slices", func(t *testing.T) {
		expected := []int{2, 0}
		got := SumAllTails([]int{1, 2}, []int{9, 0})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})

	t.Run("it handles empty slices", func(t *testing.T) {
		expected := []int{2, 0}
		got := SumAllTails([]int{1, 2}, []int{})

		if !reflect.DeepEqual(expected, got) {
			t.Errorf("expected %d, got %d", expected, got)
		}
	})
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenate := func(x, y string) string {
			return x + y
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenate, ""), "abc")
	})
}

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Riya", Balance: 100}
		chris = Account{Name: "Chris", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}

		transactions = []Transaction{
			NewTransaction(chris, riya, 100),
			NewTransaction(adil, chris, 25),
		}
	)

	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}

	AssertEqual(t, newBalanceFor(riya), 200)
	AssertEqual(t, newBalanceFor(chris), 0)
	AssertEqual(t, newBalanceFor(adil), 175)
}

func TestFind(t *testing.T) {
	t.Run("find first even number", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

		firstEvenNumber, found := Find(numbers, func(x int) bool {
			return x%2 == 0
		})

		AssertTrue(t, found)
		AssertEqual(t, firstEvenNumber, 2)
	})

	t.Run("find best programmer", func(t *testing.T) {
		type Programmer struct {
			Name string
		}

		programmers := []Programmer{
			{Name: "Bill"},
			{Name: "Marc"},
		}

		bestProgrammer, found := Find(programmers, func(x Programmer) bool {
			return x.Name == "Marc"
		})

		AssertTrue(t, found)
		AssertEqual(t, bestProgrammer, Programmer{Name: "Marc"})
	})
}

func AssertEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}

func AssertNotEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got == want {
		t.Errorf("got %v, didn't want that", got)
	}
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
