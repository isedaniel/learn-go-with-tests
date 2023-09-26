package clock

import (
	"math"
	"time"
)

// Represents a two dimensional coordinate
type Point struct {
	X float64
	Y float64
}

// Returns the second point of the hour hand of an analogue clock
func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func SecToRadian(t time.Time) float64 {
	return math.Pi
}