package structure

import "math"

type Shape interface {
	Area() float64
	// Perimeter() float64
}
type Rectangle struct {
	Height float64
	Width  float64
}
type Circle struct {
	Radius float64
}
type Triangle struct {
	Height float64
	Base   float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
func (t Triangle) Area() float64 {
	return (t.Height * t.Base) * 0.5
}

func (c Circle) Perimeter() float64 {
	return (math.Pi * 2 * c.Radius)
}
func (c Circle) Area() float64 {
	return (math.Pi * c.Radius * c.Radius)
}
func RoundOff(num float64) float64 {
	return float64(int(num*100)) / 100
}
