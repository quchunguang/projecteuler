package projecteuler

import (
	"sort"
)

//////
func Power(a, b int64) int64 {
	var i int64
	var s int64 = 1
	for i = 0; i < b; i++ {
		s *= a
	}
	return s
}
func PowInt(a, b int) int {
	var ret int = 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}
func PowSum(nlist []int, m int) int {
	var sum int
	for _, v := range nlist {
		sum += PowInt(v, m)
	}
	return sum
}

func EqualInts(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Check if v in an asc-sorted int slice
func InInts(slice []int, v int) bool {
	index := sort.SearchInts(slice, v)
	if index == len(slice) || slice[index] != v {
		return false
	}
	return true
}

// Check if v in an unsorted int slice
func InIntsUnsort(slice []int, v int) bool {
	for _, value := range slice {
		if value == v {
			return true
		}
	}
	return false
}

func SumInts(list []int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return
}
func IntsEquals(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
func MapNumInts(a int) (ret []int) {
	for a > 0 {
		ret = append(ret, a%10)
		a /= 10
	}
	return
}

//////
func RightTri(a, b, c int) bool {
	if a*a+b*b == c*c {
		return true
	}
	return false
}

//////
func Champernowne(N int) (ret int) {
	a := 1
	r := 1
	for {
		if N <= 9*a*r {
			ret = DigNums(a + (N-1)/r)[(N-1)%r]
			// fmt.Println(a+(N-1)/r, (N-1)%r, ret)
			break
		}
		N -= 9 * a * r
		a *= 10
		r++
	}
	return
}

//////
func Triangle(n int) int {
	return n * (n + 1) / 2
}
func Hexagonal(n int) int {
	return n * (2*n - 1)
}

//////
var base int = 1e10

func MulTail(a int, b int) (ret int) {
	a = a % base
	b = b % base
	ret = a * b
	ret = ret % base
	return
}
func PowerTail(a int, b int) (ret int) {
	ret = 1
	for i := 0; i < b; i++ {
		ret = MulTail(ret, a)
	}
	return
}

func DeleteIndex(f []int, i int) []int {
	return append(f[:i], f[i+1:]...)
}

//////
func DigNums(a int) (ret []int) {
	for a > 0 {
		ret = append(ret, a%10)
		a /= 10
	}
	// Reverse ret
	last := len(ret) - 1
	for i := 0; i < last-i; i++ {
		ret[i], ret[last-i] = ret[last-i], ret[i]
	}
	return
}
