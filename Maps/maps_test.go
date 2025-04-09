package maps_test

import (
	maps "aniket/Maps"
	"fmt"
	"testing"
)

func TestMaps(t *testing.T) {
	map1 := map[int]string{1: "My", 2: "name", 3: "is", 4: "Aniket"}
	got := maps.Mapfunction("Aniket", map1)
	want := true
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
func TestSreachInMaps(t *testing.T) {
	maps3 := maps.MyMaps{1: "one", 2: "two", 3: "three", 4: "four"}
	t.Run("Found", func(t *testing.T) {
		got := maps3.Found("two")
		want := true
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Not Found", func(t *testing.T) {
		got := maps3.Found("Five")
		want := false
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
func TestAddInMaps(t *testing.T) {
	map2 := maps.MyMaps{1: "one"}
	t.Run("Adding", func(t *testing.T) {
		got, _ := map2.AddValue(2, "two")
		want := maps.MyMaps{1: "one", 2: "two"}
		if !maps.AreMapsEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
	t.Run("Already Exists", func(t *testing.T) {
		got, err := map2.AddValue(2, "one")
		if err != nil {
			fmt.Println("This function passes")
			t.Errorf("%v", maps.ErrorAlreadyExist)
		}
		want := maps.MyMaps{1: "one", 2: "one"}
		if !maps.AreMapsEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestDelete(t *testing.T) {
	maps4 := maps.MyMaps{1: "one", 2: "two", 3: "three"}
	got, err := maps4.Delete(3)
	if err != nil {
		t.Fatalf(maps.ErrorNotFound.Error())
	}
	want := maps.MyMaps{1: "one", 2: "two"}
	if !maps.AreMapsEqual(got, want) {
		t.Errorf("got %v want %v", got, want)

	}
}
