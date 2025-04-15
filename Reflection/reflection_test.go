package reflection_test

import (
	reflection "aniket/Reflection"
	"reflect"
	"testing"
)

func TestReflection(t *testing.T) {
	cases := []struct {
		Name   string
		Input  interface{}
		Result []string
	}{
		{
			"Single string passing",
			"Aniket",
			[]string{"Aniket"},
		},
		{
			"String passing",
			struct {
				Name string
				City string
			}{"Aniket", "Mohali"},
			[]string{"Aniket", "Mohali"},
		},
		{
			"Single int passing",
			23,
			[]string{"23"},
		},
		{
			"String passing",
			struct {
				Name string
				Age  int
			}{"Aniket", 23},
			[]string{"Aniket", "23"},
		},
		{
			"Slice of structure passing",
			[]struct {
				Name string
				City string
			}{
				{"Aniket", "Mohali"},
				{"Gtech", "Mohali"},
			},
			[]string{"Aniket", "Mohali", "Gtech", "Mohali"},
		},
	}
	for i := 0; i < len(cases); i++ {
		t.Run(cases[i].Name, func(t *testing.T) {
			var got []string
			reflection.Walk(cases[i].Input, func(input string) {
				got = append(got, input)
			})
			if !reflect.DeepEqual(got, cases[i].Result) {
				t.Errorf("Got %v || Want %v \n", got, cases[i].Result)
			}
		})
	}
}
