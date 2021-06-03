package hw

import (
	"fmt"
	. "math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	x1, y1, x2, y2 float64
}

func (g Geom) Distance() float64 {
	if g.x1 < 0 || g.x2 < 0 || g.y1 < 0 || g.y2 < 0 {
		fmt.Println("Координаты не могут быть меньше нуля")
		return -1
	}
	// возврат расстояния между точками
	return Sqrt(Pow(g.x2-g.x1, 2) + Pow(g.y2-g.y1, 2))
}
