package roman

import "testing"

var tests = []struct {
	arabic int
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
