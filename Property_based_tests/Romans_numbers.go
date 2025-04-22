package propertybasedtests

import (
	"fmt"
	"strings"
	"time"
)

type RomanNumbers struct {
	Value  int
	Symbol string
}

var allRomanNumbers = []RomanNumbers{
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

func ConvertToRoman(numb int) string {
	var result strings.Builder
	start := time.Now()
	for _, keys := range allRomanNumbers {
		if time.Since(start) > (3 * time.Second) {
			fmt.Println("Too much time")
			break
		}
		for numb >= keys.Value {
			result.WriteString(keys.Symbol)
			numb -= keys.Value
		}

	}
	return result.String()
}

func ConvertToNumeric(str string) int {
	var numeric = 0

	for _, numeral := range allRomanNumbers {
		for strings.HasPrefix(str, numeral.Symbol) {
			numeric += numeral.Value
			str = strings.TrimPrefix(str, numeral.Symbol)
		}
	}

	return numeric
}
