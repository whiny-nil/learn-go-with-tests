package roman

import "strings"

type RomanNumber struct {
	Value  int
	Symbol string
}

var allRomanNumbers = []RomanNumber{
	{1000, "M"},
	{900, "CM"},
	{500, "D"},
	{400, "CD"},
	{100, "C"},
	{90, "XC"},
	{50, "L"},
	{40, "XL"},
	{10, "X"},
	{9, "IX"},
	{5, "V"},
	{4, "IV"},
	{1, "I"},
}

func ConvertToRoman(arabic int) string {
	var result strings.Builder

	for _, rn := range allRomanNumbers {
		for arabic >= rn.Value {
			result.WriteString(rn.Symbol)
			arabic -= rn.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) int {
	var total int

	for range roman {
		total++
	}

	return total
}
