package propertybasedtests_test

import (
	propertybasedtests "aniket/Property_based_tests"

	"fmt"
	"math/rand"
	"testing"
	"time"
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
			got := propertybasedtests.ConvertToRoman(values.Numb)
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
	/*------------------------------------------------------------------*/
	t.Run("Testing both functions", func(t *testing.T) {

		rand.Seed(time.Now().UnixNano())
		numeric := rand.Intn(3999) + 1
		roman := propertybasedtests.ConvertToRoman(numeric)
		newNumeric := propertybasedtests.ConvertToNumeric(roman)
		if newNumeric != numeric {
			t.Error("Something is Wrong")
		}

	})

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
