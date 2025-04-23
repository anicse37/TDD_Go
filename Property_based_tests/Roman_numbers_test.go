package propertybasedtests_test

import (
	"testing/quick"

	propertybasedtests "github.com/anicse37/TDD_Go/Property_based_tests"

	"fmt"
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
		{"2014 converts to MMXIV", 2014, "MMXIV"},
		{"999 converts to CMXCIX", 999, "CMXCIX"},
		{"3999 converts to CMXCIX", 3999, "MMMCMXCIX"},
	}
	/*------------------------------------------------------------------*/
	for _, values := range cases {
		t.Run(values.Desprition, func(t *testing.T) {
			got := propertybasedtests.ConvertToRoman(uint16(values.Numb))
			Compare(got, values.Want, t)
		})
	}
	/*------------------------------------------------------------------*/
	for _, test := range cases[:1] {
		t.Run(fmt.Sprintf("%q gets converted to %d", test.Want, test.Numb), func(t *testing.T) {
			got := propertybasedtests.ConvertToNumeric(test.Want)
			Compare2(got, test.Numb, t)

		})
	}
}

/*------------------------------------------------------------------*/
func TestPropertiesOfFinctions(t *testing.T) {
	assertion := func(numeric uint16) bool {
		if numeric > 3999 {
			return true
		}
		t.Log("testing", numeric)
		roman := propertybasedtests.ConvertToRoman(numeric)
		fromRoman := propertybasedtests.ConvertToNumeric(roman)
		return fromRoman == int(numeric)
	}

	if err := quick.Check(assertion, &quick.Config{
		MaxCount: 1000,
	}); err != nil {
		t.Error("failed checks", err)
	}

}

/*------------------------------------------------------------------*/
func Compare(got, want string, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}

/*------------------------------------------------------------------*/
func Compare2(got, want int, t *testing.T) {
	t.Helper()
	if got != want {
		t.Errorf("Got %v || Want %v\n", got, want)
	}
}
