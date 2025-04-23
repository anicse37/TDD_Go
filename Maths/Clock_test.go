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
		{SimpleTime(0, 0, 30), math.Pi},
		{SimpleTime(0, 0, 0), 0},
		{SimpleTime(0, 0, 45), (math.Pi / 2) * 3},
		{SimpleTime(0, 0, 18), (math.Pi / 30) * 18},
	}
	for _, values := range cases {
		t.Run(NameTest(values.time), func(t *testing.T) {
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
		{SimpleTime(0, 0, 30), clockFace.Point{0, -1}},
		{SimpleTime(0, 0, 0), clockFace.Point{0, 1}},
		{SimpleTime(0, 0, 15), clockFace.Point{1, 0}},
		{SimpleTime(0, 0, 20), clockFace.Point{0.86, -0.49}},
		{SimpleTime(0, 0, 49), clockFace.Point{-0.91, 0.4}},
	}
	for _, values := range cases {
		t.Run(NameTest(values.time), func(t *testing.T) {
			got := clockFace.SecondsHandPoint(values.time)
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
func SimpleTime(hour, minute, seconds int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, seconds, 0, time.UTC)
}

/*---------------------------------------------------------------------------*/
func NameTest(t time.Time) string {
	return t.Format("15:04:05")
	/*---------------------------------------------------------------------------*/
}
