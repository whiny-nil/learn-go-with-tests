package clockface

import (
	"math"
	"time"
)

const secondHandLength = 90
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
	rads := secondsInRadians(t)
	x := math.Sin(rads)
	y := math.Cos(rads)

	return Point{x, y}
}

func secondsInRadians(t time.Time) float64 {
	return math.Pi / (30 / float64(t.Second()))
}
