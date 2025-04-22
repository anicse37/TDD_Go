package clockFace

import (
	"math"
	"time"
)

type Point struct {
	X float64
	Y float64
}

func SecodHand(t time.Time) Point {
	return Point{150, 60}
}
func SecondsInRadians(t time.Time) float64 {
	return float64(t.Second()) * (math.Pi / 30)
}
func SimpleTime(hour, minute, seconds int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, seconds, 0, time.UTC)
}
func TestName(t time.Time) string {
	return t.Format("15:04:05")
}
func SecodHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
