package iteration

import "strings"

func For(s string) string {
	var temp string
	for i := 0; i < 5; i++ {
		temp += s
	}
	return temp
}

/*
*In Repeat function, string builder and WriteString which
*concatinates the string and does not waste memory
 */
const repeatCount = 5

func Repeat(s string) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
