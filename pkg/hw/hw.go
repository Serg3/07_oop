package hw

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

// * renamed a function and a variable with less title
func (g Geom) Distance() float64 {
	// * only if condition for simple reading
	if g.X1 < 0 || g.X2 < 0 || g.Y1 < 0 || g.Y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1
	}
	// возврат расстояния между точками
	// * removed returned variable
	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
}
