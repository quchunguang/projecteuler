package projecteuler

import (
	"math"
)

func Factors(n int) (fs []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			fs = append(fs, i)
		}
	}
	return
}

//////
var Pentagons []int

func Pentagonal(n int) int {
	return n * (3*n - 1) / 2
}

func GenPentagons(max int) []int {
	curr := len(Pentagons)
	if curr >= max {
		return Pentagons
	}
	for n := curr + 1; ; n++ {
		Pentagons = append(Pentagons, Pentagonal(n))
		if Pentagonal(n) >= max {
			break
		}
	}
	return Pentagons
}

//////
var Abundants []int

func GenAbundants(N int) {
	for i := 2; i <= N; i++ {
		if SumInts(Factors(i)) > i {
			Abundants = append(Abundants, i)
		}
	}
}

//////
type DivMod struct {
	Div, Mod int
}

func InDivMods(dms []DivMod, dm DivMod) int {
	for i, v := range dms {
		if v.Div == dm.Div && v.Mod == dm.Mod {
			return i
		}
	}
	return -1
}

//////
// If n*n + a*n + b are all primes when n = 0..N-1, return N. Or, return -1
//   n² + n + 41, N=40
//   n² - 79n + 1601, N=80
func QuadraticPrimes(a, b int) (n int) {
	for {
		M := n*n + a*n + b
		GenPrimes(M)
		if !InInts(primes, M) {
			break
		}
		n++
	}
	return
}

//////
func IsOddComposites(n int) bool {
	var i int = 1
	for p := n - 2*i*i; p > 0; p = n - 2*i*i {
		if InInts(primes, p) {
			return true
		}
		i++
	}
	return false
}

//////
var triangles = []int{1, 3}

func IsTriangle(N int) bool {
	if N > triangles[len(triangles)-1] {
		for i := len(triangles) + 1; triangles[len(triangles)-1] < N; i++ {
			triangles = append(triangles, i*(i+1)/2)
		}
	}
	if InInts(triangles, N) {
		return true
	}
	return false
}

//////
// Generate all primes less than max to global primes
var primes = []int{2, 3, 5, 7, 11, 13}

func GenPrimes(max int) {
	if primes[len(primes)-1] >= max {
		return
	}
	for i := primes[len(primes)-1] + 2; i <= max; i += 2 {
		upbound := int(math.Sqrt(float64(i)))
		for j := 0; primes[j] <= upbound; j++ {
			if i%primes[j] == 0 {
				goto NEXT
			}
		}
		primes = append(primes, i)
		// fmt.Println(len(primes), &primes[0])
	NEXT:
	}
}

// Generate prime factors and return as a map
func PrimeFactors(n int) map[int]int {
	if primes[len(primes)-1] < n {
		GenPrimes(n * 2)
	}
	var pfmap = make(map[int]int)
	for j := 0; n != 1; j++ {
		for n%primes[j] == 0 {
			pfmap[primes[j]]++
			n /= primes[j]
		}
	}
	return pfmap
}

//////
var sumprimes = []int{2, 5, 10, 17, 28, 41}

// Generate summery of primes at lest to n
func SumPrimes(n int) {
	if sumprimes[len(sumprimes)-1] >= n {
		return
	}

	for sumprimes[len(sumprimes)-1] < n {
		GenPrimes(primes[len(primes)-1] * 2)
		for i := len(sumprimes); i < len(primes); i++ {
			sumprimes = append(sumprimes, sumprimes[i-1]+primes[i])
		}
	}
}

// Φ(n) = n * (1 - 1/p1) * (1 - 1/p2) * ... * (1 - 1/pk)
func EulerPhi(N int64) (ret int64) {
	ret = N
	pfsmap := PrimeFactors(int(N))
	for p, _ := range pfsmap {
		ret *= int64(p) - 1
		ret /= int64(p)
	}
	return
}
