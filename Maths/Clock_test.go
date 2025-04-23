package clockFace_test

import (
	clockFace "aniket/Maths"
	"math"
	"testing"
	"time"
)

func TestSecondsInRadians(t *testing.T) {
	cases := []struct {
		time  time.Time
		angle float64
	}{
		{clockFace.SimpleTime(0, 0, 30), math.Pi},
		{clockFace.SimpleTime(0, 0, 0), 0},
		{clockFace.SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{clockFace.SimpleTime(0, 0, 18), (math.Pi / 30) * 18},
	}
	for _, values := range cases {
		t.Run(clockFace.TestName(values.time), func(t *testing.T) {
			got := clockFace.SecondsInRadians(values.time)
			if !compare(got, values.angle) {
				t.Fatalf("Got %v || Want %v ", got, values.angle)
			}
		})
	}
}

/*---------------------------------------------------------------------------*/
func TestSecodHandPoint(t *testing.T) {
	cases := []struct {
		time  time.Time
		point clockFace.Point
	}{
		{clockFace.SimpleTime(0, 0, 30), clockFace.Point{0, -1}},
		{clockFace.SimpleTime(0, 0, 0), clockFace.Point{0, 1}},
		{clockFace.SimpleTime(0, 0, 15), clockFace.Point{1, 0}},
		{clockFace.SimpleTime(0, 0, 20), clockFace.Point{0.86, -0.49}},
		{clockFace.SimpleTime(0, 0, 49), clockFace.Point{-0.91, 0.4}},
	}
	for _, values := range cases {
		t.Run(clockFace.TestName(values.time), func(t *testing.T) {
			got := clockFace.SecodHandPoint(values.time)
			if !roughlyEqualTo(got, values.point) {
				t.Fatalf("Got %v || Want %v \n", got, values.point)
			}
		})
	}
}

/*---------------------------------------------------------------------------*/
func roughlyEqualTo(a, b clockFace.Point) bool {
	return compare(a.X, b.X) && compare(a.Y, b.Y)
}

/*---------------------------------------------------------------------------*/
func compare(got, want float64) bool {
	got *= 100
	temp := int(got)
	got = float64(temp) / 100
	// fmt.Println(got)
	want *= 100
	temp = int(want)
	want = float64(temp) / 100
	// fmt.Println(want)
	return got == want
}

/*---------------------------------------------------------------------------*/
