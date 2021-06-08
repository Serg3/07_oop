package hw

import (
	"math"
)

// Geom struct describes coordinates for 2 points
// * added comments to teh export objects
type Geom struct {
	X1, Y1, X2, Y2 float64
}

// Distance() calculates and returns a distance
// between between two 2D points
// * renamed a function and a variable with less title
func (g Geom) Distance() float64 {
	// * removed returned variable
	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
}
