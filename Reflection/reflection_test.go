package reflection_test

import (
	reflection "aniket/Reflection"
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Aniket"},
			[]string{"Aniket"},
		},
		{
			"struct with one string field",
			struct {
				Name string
				City string
			}{"Aniket", "Mohali"},
			[]string{"Aniket", "Mohali"},
		},
		{
			"struct with slice field",
			[]struct { //struct
				Name string
				Age  int
			}{
				{"Aniket", 23}, //slice and int
				{"Gtech", 25},
				{"QWE", 24},
			},
			[]string{"Aniket", "23", "Gtech", "25", "QWE", "24"},
		},
		{
			"struct with one string field",
			2,
			[]string{"2"},
		},
	}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			reflection.Walk(test.Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
