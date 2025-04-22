package reflection_test

import (
	reflection "aniket/Reflection"
	"reflect"
	"testing"
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
			myChannel <- channelStruct{"Aniket", 23}
			myChannel <- channelStruct{"Golang", 16}
			close(myChannel)
		}()
		want := []string{"Aniket", "23", "Golang", "16"}
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
			"Aniket": "Mohali",
			"Gtech":  "Chandigarh",
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

/*----------------------------------------------------------------------------------*/
//To check if 2 maps are equal
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
