package reflection_test

import (
	"reflect"
	"testing"

	reflection "github.com/anicse37/TDD_Go/Reflection"
)

func TestReflection(t *testing.T) {
	type channelStruct struct {
		Name string
		Age  int
	}
	function := func() (channelStruct, channelStruct) {
		return channelStruct{"one", 1}, channelStruct{"three", 3}
	}
	cases := []struct {
		Name   string
		Input  interface{}
		Result []string
	}{
		{
			"Single string passing",
			"github.com/anicse37/TDD_Go",
			[]string{"github.com/anicse37/TDD_Go"},
		},
		{
			"String passing",
			struct {
				Name string
				City string
			}{"github.com/anicse37/TDD_Go", "Mohali"},
			[]string{"github.com/anicse37/TDD_Go", "Mohali"},
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
			}{"github.com/anicse37/TDD_Go", 23},
			[]string{"github.com/anicse37/TDD_Go", "23"},
		},
		{
			"Slice of structure passing",
			[]struct {
				Name string
				City string
			}{
				{"github.com/anicse37/TDD_Go", "Mohali"},
				{"Gtech", "Mohali"},
			},
			[]string{"github.com/anicse37/TDD_Go", "Mohali", "Gtech", "Mohali"},
		},
		{
			"Using Functions",
			function,
			[]string{"one", "1", "three", "3"},
		},
	}
	t.Run("Test for Channles", func(t *testing.T) {
		var got []string
		myChannel := make(chan channelStruct)
		go func() {
			myChannel <- channelStruct{"github.com/anicse37/TDD_Go", 23}
			myChannel <- channelStruct{"Golang", 16}
			close(myChannel)
		}()
		want := []string{"github.com/anicse37/TDD_Go", "23", "Golang", "16"}
		reflection.Walk(myChannel, func(input string) {
			got = append(got, input)
		})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("Got %v || Want %v \n", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (channelStruct, channelStruct) {
			return channelStruct{"ten", 10}, channelStruct{"six", 6}
		}

		var got []string
		want := []string{"ten", "10", "six", "6"}

		reflection.Walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"github.com/anicse37/TDD_Go": "Mohali",
			"Gtech":                      "Chandigarh",
		}

		var got []string
		reflection.Walk(aMap, func(input string) {
			got = append(got, input)
		})

		AreMapsEqual(t, got, "Mohali")
		AreMapsEqual(t, got, "Chandigarh")
	})
	/*--------------------------------------------------------------------------*/
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
func AreMapsEqual(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", haystack, needle)
	}
}
