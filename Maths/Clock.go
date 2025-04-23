package clockFace

import (
	"math"
	"time"
)

const secondHandLength = 90
const clockCentreX = 150
const clockCentreY = 150

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	P := SecodHandPoint(t)
	P = Point{P.X * secondHandLength, P.Y * secondHandLength} //Scale
	P = Point{P.X, -P.Y}                                      // invert y from top left to bottom left
	P = Point{P.X + clockCentreX, P.Y + clockCentreY}         //transform from (0,0) to (150,150)
	return P
}
func SecodHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}
func SimpleTime(hour, minute, seconds int) time.Time {
	return time.Date(312, time.October, 28, hour, minute, seconds, 0, time.UTC)
}
func TestName(t time.Time) string {
	return t.Format("15:04:05")
}
