package roman

import "strings"

type RomanNumber struct {
	Value  uint16
	Symbol string
}

type RomanNumbers []RomanNumber

func (r RomanNumbers) ValueOf(symbols ...byte) uint16 {
	symbol := string(symbols)
	for _, s := range r {
		if s.Symbol == symbol {
			return s.Value
		}
	}

	return 0
}

var allRomanNumbers = RomanNumbers{
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

func ConvertToRoman(arabic uint16) string {
	var result strings.Builder

	for _, rn := range allRomanNumbers {
		for arabic >= rn.Value {
			result.WriteString(rn.Symbol)
			arabic -= rn.Value
		}
	}

	return result.String()
}

func ConvertToArabic(roman string) uint16 {
	var total uint16 = 0

	for i := 0; i < len(roman); i++ {
		symbol := roman[i]

		if couldBeSubtractive(i, symbol, roman) {
			if value := allRomanNumbers.ValueOf(symbol, roman[i+1]); value != 0 {
				total += value
				i++ // move past this character for the next loop
			} else {
				total += allRomanNumbers.ValueOf(symbol)
			}
		} else {
			total += allRomanNumbers.ValueOf(symbol)
		}
	}

	return total
}

func couldBeSubtractive(index int, currentSymbol uint8, roman string) bool {
	return index + 1 < len(roman) && (currentSymbol == 'I' || currentSymbol == 'X' || currentSymbol == 'C')
}