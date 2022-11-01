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

func Reduce[T any](collection []T, accumulator func(T, T) T, initialValue T) T {
	value := initialValue
	for _, nextValue := range collection {
		value = accumulator(value, nextValue)
	}

	return value
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func BalanceFor(transactions []Transaction, name string) float64 {
	var balance float64
	for _, t := range transactions {
		if t.From == name {
			balance -= t.Sum
		}
		if t.To == name {
			balance += t.Sum
		}
	}

	return balance
}
