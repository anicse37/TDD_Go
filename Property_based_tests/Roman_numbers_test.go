package propertybasedtests_test

import (
	propertybasedtests "aniket/Property_based_tests"
	"testing"
)

func TestRomanNumbers(t *testing.T) {
	cases := []struct {
		Desprition string
		Numb       int
		Want       string
	}{
		{"1 converts to I", 1, "I"},
		{"2 converts to II", 2, "II"},
		{"3 converts to III", 3, "III"},
		{"4 converts to IV", 4, "IV"},
		{"9 converts to IX", 9, "IX"},
		{"10 converts to X", 10, "X"},
		{"49 converts to XLIX", 49, "XLIX"},
		{"400 converts to CD", 400, "CD"},
		{"1984 converts to MCMLXXXIV", 1984, "MCMLXXXIV"},
		{"2014 converts to MCMLXXXIV", 2014, "MMXIV"},
		{"999 converts to MCMLXXXIV", 999, "CMXCIX"},
	}
	for _, values := range cases {
		t.Run(values.Desprition, func(t *testing.T) {
			got := propertybasedtests.ConvertToRoman(values.Numb)
			Compare(got, values.Want, t)
		})
	}
}
func Compare(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
