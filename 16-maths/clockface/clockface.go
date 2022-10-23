package clockface

import (
	"math"
	"time"
)

const clockCenterX = 150
const clockCenterY = 150

// A Point represents a 2D Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time 't',
// represented as a point
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * secondHandLength, p.Y * secondHandLength} // scale
	p = Point{p.X, -p.Y}                                      // flip over X-axis
	p = Point{p.X + clockCenterX, p.Y + clockCenterY}         // translate

	return p
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	return (secondsInRadians(t) / 60) + (math.Pi / (30 / float64(t.Minute())))
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func hoursInRadians(t time.Time) float64 {
	return (minutesInRadians(t) / 12) + (math.Pi / (6 / float64(t.Hour()%12)))
}

func angleToPoint(rads float64) Point {
	x := math.Sin(rads)
	y := math.Cos(rads)

	return Point{x, y}
}
