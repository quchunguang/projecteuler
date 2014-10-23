package projecteuler

import (
	"fmt"
)

// Problem 251 - Cardano Triplets
//
func PE251() (ret int) {
	return
}

// Problem 252 - Convex Holes
// Given a set of points on a plane, we define a convex hole to be a convex polygon having as vertices any of the given points and not containing any of the given points in its interior (in addition to the vertices, other given points may lie on the perimeter of the polygon).
//
// As an example, the image below shows a set of twenty points and a few such convex holes. The convex hole shown as a red heptagon has an area equal to 1049694.5 square units, which is the highest possible area for a convex hole on the given set of points.
//
// For our example, we used the first 20 points (T2k−1, T2k), for k = 1,2,…,20, produced with the pseudo-random number generator:
// S0   =   290797
// Sn+1     =   Sn2 mod 50515093
// Tn   =   ( Sn mod 2000 ) − 1000
//
// i.e. (527, 144), (−488, 732), (−454, −947), …
//
// What is the maximum area for a convex hole on the set containing the first 500 points in the pseudo-random sequence?
// Specify your answer including one digit after the decimal point.
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

// Problem 253 - Tidying up
//
func PE253() (ret int) {
	return
}

// Problem 254 - Sums of Digit Factorials
//
func PE254() (ret int) {
	return
}

// Problem 255 - Rounded Square Roots
//
func PE255() (ret int) {
	return
}

// Problem 256 - Tatami-Free Rooms
//
func PE256() (ret int) {
	return
}

// Problem 257 - Angular Bisectors
//
func PE257() (ret int) {
	return
}

// Problem 258 - A lagged Fibonacci sequence
//
func PE258() (ret int) {
	return
}

// Problem 259 - Reachable Numbers
//
func PE259() (ret int) {
	return
}

// Problem 260 - Stone Game
//
func PE260() (ret int) {
	return
}

// Problem 261 - Pivotal Square Sums
//
func PE261() (ret int) {
	return
}

// Problem 262 - Mountain Range
//
func PE262() (ret int) {
	return
}

// Problem 263 - An engineers' dream come true
//
func PE263() (ret int) {
	return
}

// Problem 264 - Triangle Centres
//
func PE264() (ret int) {
	return
}

// Problem 265 - Binary Circles
//
func PE265() (ret int) {
	return
}

// Problem 266 - Pseudo Square Root
//
func PE266() (ret int) {
	return
}

// Problem 267 - Billionaire
//
func PE267() (ret int) {
	return
}

// Problem 268 - Counting numbers with at least four distinct prime factors less than 100
//
func PE268() (ret int) {
	return
}

// Problem 269 - Polynomials with at least one integer root
//
func PE269() (ret int) {
	return
}

// Problem 270 - Cutting Squares
//
func PE270() (ret int) {
	return
}

// Problem 271 - Modular Cubes, part 1
//
func PE271() (ret int) {
	return
}

// Problem 272 - Modular Cubes, part 2
//
func PE272() (ret int) {
	return
}

// Problem 273 - Sum of Squares
//
func PE273() (ret int) {
	return
}

// Problem 274 - Divisibility Multipliers
//
func PE274() (ret int) {
	return
}

// Problem 275 - Balanced Sculptures
//
func PE275() (ret int) {
	return
}

// Problem 276 - Primitive Triangles
//
func PE276() (ret int) {
	return
}

// Problem 277 - A Modified Collatz sequence
//
func PE277() (ret int) {
	return
}

// Problem 278 - Linear Combinations of Semiprimes
//
func PE278() (ret int) {
	return
}

// Problem 279 - Triangles with integral sides and an integral angle
//
func PE279() (ret int) {
	return
}

// Problem 280 - Ant and seeds
//
func PE280() (ret int) {
	return
}

// Problem 281 - Pizza Toppings
//
func PE281() (ret int) {
	return
}

// Problem 282 - The Ackermann function
//
func PE282() (ret int) {
	return
}

// Problem 283 - Integer sided triangles for which the  area/perimeter ratio is integral
//
func PE283() (ret int) {
	return
}

// Problem 284 - Steady Squares
//
func PE284() (ret int) {
	return
}

// Problem 285 - Pythagorean odds
//
func PE285() (ret int) {
	return
}

// Problem 286 - Scoring probabilities
//
func PE286() (ret int) {
	return
}

// Problem 287 - Quadtree encoding (a simple compression algorithm)
//
func PE287() (ret int) {
	return
}

// Problem 288 - An enormous factorial
//
func PE288() (ret int) {
	return
}

// Problem 289 - Eulerian Cycles
//
func PE289() (ret int) {
	return
}

// Problem 290 - Digital Signature
//
func PE290() (ret int) {
	return
}

// Problem 291 - Panaitopol Primes
//
func PE291() (ret int) {
	return
}

// Problem 292 - Pythagorean Polygons
//
func PE292() (ret int) {
	return
}

// Problem 293 - Pseudo-Fortunate Numbers
//
func PE293() (ret int) {
	return
}

// Problem 294 - Sum of digits - experience #23
//
func PE294() (ret int) {
	return
}

// Problem 295 - Lenticular holes
//
func PE295() (ret int) {
	return
}

// Problem 296 - Angular Bisector and Tangent
//
func PE296() (ret int) {
	return
}

// Problem 297 - Zeckendorf Representation
//
func PE297() (ret int) {
	return
}

// Problem 298 - Selective Amnesia
//
func PE298() (ret int) {
	return
}

// Problem 299 - Three similar triangles
//
func PE299() (ret int) {
	return
}

// Problem 300 - Protein folding
//
func PE300() (ret int) {
	return
}
