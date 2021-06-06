// * added test program
package main

import (
	"fmt"
	"go_core/07_oop/pkg/hw"
)

func main() {
	g := hw.Geom{}
	g.X1, g.Y1, g.X2, g.Y2 = 1, 1, 4, 5

	fmt.Printf("Distance between (%v, %v) and (%v, %v) is %v.\n", g.X1, g.Y1, g.X2, g.Y2, g.Distance())
}
