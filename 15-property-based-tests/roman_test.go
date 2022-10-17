package roman

import "strings"
import "testing"
import "testing/quick"

var tests = []struct {
	arabic uint16
	roman  string
}{
	{1, "I"},
	{2, "II"},
	{3, "III"},
	{4, "IV"},
	{5, "V"},
	{6, "VI"},
	{7, "VII"},
	{8, "VIII"},
	{9, "IX"},
	{10, "X"},
	{14, "XIV"},
	{18, "XVIII"},
	{20, "XX"},
	{39, "XXXIX"},
	{40, "XL"},
	{44, "XLIV"},
	{47, "XLVII"},
	{49, "XLIX"},
	{50, "L"},
	{89, "LXXXIX"},
	{400, "CD"},
	{500, "D"},
	{900, "CM"},
	{1000, "M"},
	{1984, "MCMLXXXIV"},
	{3999, "MMMCMXCIX"},
	{2014, "MMXIV"},
	{1006, "MVI"},
	{798, "DCCXCVIII"},
}

func TestRomanNumbers(t *testing.T) {
	for _, test := range tests {
		t.Run(`$test.arabic => $test.roman`, func(t *testing.T) {
			got := ConvertToRoman(test.arabic)
			expected := test.roman

			if got != expected {
				t.Errorf("expected %s, got %s", expected, got)
			}
		})
	}
}

func TestArabicNumbers(t *testing.T) {
	for _, test := range tests {
		t.Run(`$test.roman => $test.arabic`, func(t *testing.T) {
			got := ConvertToArabic(test.roman)
			expected := test.arabic

			if got != expected {
				t.Errorf("expected %d, got %d", expected, got)
			}
		})
	}
}

// Test that converting arabic to roman and back gives the same number that we started with
func TestConversionProperty(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(arabic)
		fromRoman := ConvertToArabic(roman)
		return fromRoman == arabic
	}

	if err := quick.Check(assertion, nil); err != nil {
		t.Error("failed conversion property check", err)
	}
}

// Test that converting to roman produces a string with no more than 3 consecutive symbols being the same
// Test that only I, X and C can be "subtractors"
func TestFormatProperty(t *testing.T) {
	assertion := func(arabic uint16) bool {
		if arabic > 3999 {
			return true
		}
		roman := ConvertToRoman(arabic)

		invalidStrings := [10]string{"IIII", "VVVV", "XXXX", "LLLL", "CCCC", "DDDD", "MMMM", "VX", "LC", "DM"}
		containsInvalidString := false

		for _, str := range invalidStrings {
			containsInvalidString = containsInvalidString || strings.Contains(roman, str)
		}

		return !containsInvalidString
	}

	err := quick.Check(assertion, nil)
	if err != nil {
		t.Error("failed formatting check", err)
	}
}
