package maps_test

import (
	maps "aniket/Maps"
	"testing"
)

/*---------------------------------------------------*/

func TestMap(t *testing.T) {
	M1 := maps.MyMaps{1: "One", 2: "Two"}
	got := M1.TestMap()
	want := maps.MyMaps{1: "One", 2: "Two"}
	if !(maps.AreMapsSame(got, want)) {
	}
}

/*---------------------------------------------------*/
func TestSearch(t *testing.T) {
	M2 := maps.MyMaps{1: "One", 2: "Two", 3: "Three"}
	got, err := M2.Search("Two")
	want := true
	if got != want {
		t.Fatal(err)
	}
}

/*---------------------------------------------------*/

func TestAdd(t *testing.T) {
	t.Run("New Place", func(t *testing.T) {

		M3 := maps.MyMaps{1: "One", 2: "Two", 3: "Three"}
		got, err := M3.Add(4, "Four")
		if err != nil {
			t.Fatal(err)
		}
		want := maps.MyMaps{1: "One", 2: "Two", 3: "Three", 4: "Four"}
		if !maps.AreMapsSame(got, want) {
			t.Errorf("Got %v || Want %v ", got, want)
		}
	})
	// t.Run("Old Place", func(t *testing.T) {
	// 	M4 := maps.MyMaps{1: "One", 2: "Two", 3: "Three"}
	// 	got, err := M4.Add(2, "Two")
	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	want := maps.MyMaps{1: "One", 2: "Two", 3: "Three"}

	// 	if !maps.AreMapsSame(got, want) {
	// 		t.Errorf("Got %v || Want %v ", got, want)
	// 	}

	// })
}

/*---------------------------------------------------*/
func TestDelete(t *testing.T) {
	M5 := maps.MyMaps{1: "One", 2: "Two", 3: "Three"}
	got, err := M5.Delete(2, "Two")
	if err != nil {
		t.Fatal(err)
	}
	want := maps.MyMaps{1: "One", 3: "Three"}
	if !maps.AreMapsSame(got, want) {
		t.Errorf("Got %v || Want %v ", got, want)
	}

}

/*---------------------------------------------------*/
