package sums

func Sum(numbers []int) int {
	adder := func(a, b int) int {
		return a + b
	}
	return Reduce(numbers, adder, 0)
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int

	adder := func(a, b []int) []int {
		return append(a, Sum(b))
	}
	return Reduce(numbersToSum, adder, sums)
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int

	adder := func(a, b []int) []int {
		if len(b) == 0 {
			return append(a, 0)
		} else {
			tail := b[1:]
			return append(a, Sum(tail))
		}
	}
	return Reduce(numbersToSum, adder, sums)
}

func Reduce[A, B any](collection []A, accumulator func(B, A) B, initialValue B) B {
	value := initialValue
	for _, nextValue := range collection {
		value = accumulator(value, nextValue)
	}

	return value
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, item := range items {
		if predicate(item) {
			return item, true
		}
	}

	return
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

type Account struct {
	Name    string
	Balance float64
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}
