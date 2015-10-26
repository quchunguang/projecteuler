package projecteuler

import (
	"sort"
	"strconv"
)

// Return a^b where a,b are int64.
func Power(a, b int64) int64 {
	var i int64
	var s int64 = 1
	for i = 0; i < b; i++ {
		s *= a
	}
	return s
}

// Return a^b where a,b are int.
func PowInt(a, b int) int {
	var ret int = 1
	for i := 0; i < b; i++ {
		ret *= a
	}
	return ret
}

// Return A1^m+A2^m+..., for Ai in nlist.
func PowSum(nlist []int, m int) int {
	var sum int
	for _, v := range nlist {
		sum += PowInt(v, m)
	}
	return sum
}

// Check if two slice of ints are totally equal.
func EqualInts(a, b []int) bool {
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

// Check if v in an asc-sorted int slice
func InInts(slice []int, v int) bool {
	index := sort.SearchInts(slice, v)
	if index == len(slice) || slice[index] != v {
		return false
	}
	return true
}

// Join concatenates the elements of ints to create a single string.
func JoinInts(slice []int, sep string) (ret string) {
	for _, i := range slice {
		ret += strconv.Itoa(i)
		ret += sep
	}
	ret = ret[:len(ret)-len(sep)]
	return
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

// Sum of slice of ints
func SumInts(list []int) (sum int) {
	for _, v := range list {
		sum += v
	}
	return
}

// Get all digital numbers of a int.
func MapNumInts(a int) (ret []int) {
	for a > 0 {
		ret = append(ret, a%10)
		a /= 10
	}
	return
}

// Check if (a,b,c) are right Triangle numbers
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

//////
func Gcd(m, n int) int {
	if m < n {
		m, n = n, m
	}
	for m%n != 0 {
		m, n = n, m%n
	}
	return n
}

// Return Max int from []int
func MaxInts(slice []int) int {
	var ret int = MinInt
	for _, v := range slice {
		if ret < v {
			ret = v
		}
	}
	return ret
}
