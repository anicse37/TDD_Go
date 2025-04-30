package clockFace

import (
	"fmt"
	"io"
	"math"
	"time"
)

const (
	secondHandLength = 90
	minuteHandLength = 80
	hourHandLength   = 60
	clockCentreX     = 150
	clockCentreY     = 150
)

type Point struct {
	X float64
	Y float64
}

/*-----------------------------------------------------------------------------------------------------*/
func SVGWritter(w io.Writer, t time.Time) {
	sh := SecondHand(t)
	mh := MinuteHand(t)
	hh := HourHand(t)
	text := GetClockNumbers()
	io.WriteString(w, svgStart)
	io.WriteString(w, bezel)
	io.WriteString(w, secondHandTag(sh))
	io.WriteString(w, minuteHandTag(mh))
	io.WriteString(w, hourHandTag(hh))
	io.WriteString(w, timeLetter(text))

	io.WriteString(w, svgEnd)
}

/*-----------------------------------------------------------------------------------------------------*/
func GetClockNumbers() string {
	var text string
	for i := 1; i <= 12; i++ {
		angle := (float64(i) * math.Pi) / 6
		x := clockCentreY + math.Sin(angle)*90
		y := clockCentreY + 5 - math.Cos(angle)*90
		text += fmt.Sprintf(`<text x="%f" y="%f" font-size="12" text-anchor="middle" fill="black">%d</text>`, x, y, i)
	}
	return text
}

/*-----------------------------------------------------------------------------------------------------*/
func SecondHand(t time.Time) Point {
	P := SecondsHandPoint(t)
	P = Point{P.X * secondHandLength, P.Y * secondHandLength} //Scale
	P = Point{P.X, -P.Y}                                      // invert y from top left to bottom left
	P = Point{P.X + clockCentreX, P.Y + clockCentreY}         //transform from (0,0) to (150,150)
	return P
}
func SecondsHandPoint(t time.Time) Point {
	angle := SecondsInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func SecondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

/*-----------------------------------------------------------------------------------------------------*/

func MinuteHand(t time.Time) Point {
	P := MinuteHandPoint(t)
	P = Point{P.X * minuteHandLength, P.Y * minuteHandLength}
	P = Point{P.X, -P.Y}
	P = Point{P.X + clockCentreX, P.Y + clockCentreY}
	return P
}
func MinuteHandPoint(t time.Time) Point {
	angle := MinuteInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func MinuteInRadians(t time.Time) float64 {
	return ((math.Pi / (30 / (float64(t.Minute())))) + (SecondsInRadians(t) / 60))
}

/*-----------------------------------------------------------------------------------------------------*/
func HourHand(t time.Time) Point {
	P := HourHandPoint(t)
	P = Point{P.X * hourHandLength, P.Y * hourHandLength}
	P = Point{P.X, -P.Y}
	P = Point{P.X + clockCentreX, P.Y + clockCentreY}
	return P
}
func HourHandPoint(t time.Time) Point {
	angle := HourInRadians(t)
	x := math.Sin(angle)
	y := math.Cos(angle)
	return Point{x, y}
}
func HourInRadians(t time.Time) float64 {
	return ((math.Pi / (6 / float64(t.Hour()%12))) + MinuteInRadians(t)/12)
}

/*-----------------------------------------------------------------------------------------------------*/
func secondHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#f00;stroke-width:3px;"/>`, p.X, p.Y)
}
func minuteHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000000;stroke-width:3px;"/>`, p.X, p.Y)
}
func timeLetter(text string) string {
	return fmt.Sprintf("%s", text)
}
func hourHandTag(p Point) string {
	return fmt.Sprintf(`<line x1="150" y1="150" x2="%f" y2="%f" style="fill:none;stroke:#000000;stroke-width:3px;"/>`, p.X, p.Y)
}

const svgStart = `<?xml version="1.0" encoding="UTF-8" standalone="no"?>
<!DOCTYPE svg PUBLIC "-//W3C//DTD SVG 1.1//EN" "http://www.w3.org/Graphics/SVG/1.1/DTD/svg11.dtd">
<svg xmlns="http://www.w3.org/2000/svg"
     width="100%"
     height="100%"
     viewBox="0 0 300 300"
     version="2.0">`

const bezel = `<circle cx="150" cy="150" r="100" style="fill:#fff;stroke:#000;stroke-width:5px;"/>`

const svgEnd = `</svg>`
