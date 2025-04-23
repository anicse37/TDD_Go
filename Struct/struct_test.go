package structure_test

import (
	"fmt"
	"testing"

	structure "github.com/anicse37/TDD_Go/Struct"
)

func TestPerimeter(t *testing.T) {
	t.Run("This is perimeter of triangle", func(t *testing.T) {
		rectangle := structure.Rectangle{10, 10}
		got := rectangle.Perimeter()
		want := 40.0
		if got != want {
			t.Errorf("got %.2f | want %.2f", got, want)
		}
	})
	t.Run("This is perimeter  of circle", func(t *testing.T) {
		radius := structure.Circle{5}
		got := structure.RoundOff(radius.Perimeter())
		want := structure.RoundOff(31.415927)
		if got != want {
			t.Errorf("got %f | want %f", got, want)
		}
	})
}
func TestArea(t *testing.T) {
	t.Run("This is area of rectangle", func(t *testing.T) {
		area := structure.Rectangle{10, 5}
		got := area.Area()
		want := 50.0
		if got != want {
			t.Errorf("got %f | want %f", got, want)
		}
	})
	t.Run("This is area of circle", func(t *testing.T) {
		radius := structure.Circle{5}
		got := structure.RoundOff(radius.Area())
		want := structure.RoundOff(78.539816)
		if got != want {
			t.Errorf("got %f | want %f", got, want)
		}
	})
}

func TestAreaShape(t *testing.T) {
	TestAll := []struct {
		shape structure.Shape
		Want  float64
	}{
		{structure.Rectangle{10, 20}, 200.0},
		{structure.Circle{10}, 314.1592},
		{structure.Rectangle{10, 5}, 50},
		{structure.Triangle{10, 5}, 25},
	}
	for index, _ := range TestAll {
		got := structure.RoundOff(TestAll[index].shape.Area())
		want := structure.RoundOff(TestAll[index].Want)
		fmt.Println(want, got)
	}

	/*
	* I prefer the upper method with  ,_
	 */

	// for _, value := range TestAll {
	// 	got := structure.RoundOff(value.shape.Area())
	// 	want := structure.RoundOff(value.Want)
	// 	fmt.Println(want, got)
	// }
}
