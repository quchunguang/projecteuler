package projecteuler

import (
	"fmt"
)

//////
type Point struct {
	X, Y int
}

var Points []Point

func GenPoints(N int) {
	var (
		Sb         int = 290797
		Sn, Tx, Ty int
	)
	for n := 1; n <= N; n++ {
		Sn = (Sb * Sb) % 50515093
		Tx = (Sn % 2000) - 1000
		Sb = Sn
		Sn = (Sb * Sb) % 50515093
		Ty = (Sn % 2000) - 1000
		Sb = Sn
		Points = append(Points, Point{Tx, Ty})
	}
}
func IsConvex(ps []Point, p Point)   {}
func InHole(ps []Point, p Point)     {}
func Area(ps []Point) (area float64) { return 0.0 }
func GenTri() chan [3]Point {
	var c = make(chan [3]Point, 100)
	go func() {
		for i := 0; i < len(Points); i++ {
			for j := i + 1; j < len(Points); j++ {
				for k := j + 1; k < len(Points); k++ {
					c <- [3]Point{Points[i], Points[j], Points[k]}
				}
			}
		}
		close(c)
	}()
	return c
}
func PE252() int {
	var N int = 500
	GenPoints(N)
	s := 0
	for _ = range GenTri() {
		s++
	}
	fmt.Println(s)
	return 0
}
