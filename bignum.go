package projecteuler

import (
	"fmt"
	"strconv"
	"strings"
)

// Create Big int to [lint every 18-numbers in lower first order
func BigNum(data string) []int64 {
	var length, zeros int

	// Add leading zeros
	data = strings.TrimSpace(data)
	if len(data)%18 != 0 {
		zeros = 18 - len(data)%18
		data = strings.Repeat("0", zeros) + data
	}
	length = len(data) / 18

	// Fill int slice in lower first order
	ret := make([]int64, length)
	for i := 0; i < length; i++ {
		ret[length-i-1], _ = strconv.ParseInt(data[i*18:(i+1)*18], 10, 64)
	}
	return ret
}

// Sum two BigNums
func BigSum(a, b []int64) []int64 {
	if len(a) < len(b) {
		a, b = b, a
	}
	lena := len(a)
	lenb := len(b)
	var ret = make([]int64, lena)
	copy(ret, a)
	for i := 0; i < lenb; i++ {
		ret[i] = a[i] + b[i]
	}
	for i := 0; i < lena; i++ {
		if ret[i] >= 1e18 {
			ret[i] -= 1e18
			if i == lena-1 {
				ret = append(ret, 1)
			} else {
				ret[i+1]++
			}
		}
	}
	return ret
}

// Multiply BigNum with int
func BigMulInt(a []int64, b int64) []int64 {
	ret := BigNum("0")
	s := BigSum(ret, a)
	for b > 0 {
		if b%2 == 1 {
			ret = BigSum(ret, s)
		}
		s = BigSum(s, s)
		b /= 2
	}
	return ret
}

// Multiply two BigNum
func BigMul(a, b []int64) []int64 {
	tmp := BigSum(a, BigNum("0"))
	ret := BigNum("0")
	for i := 0; i < len(b); i++ {
		ret = BigSum(ret, BigMulInt(tmp, b[i]))
		tmp = BigMulInt(tmp, 1e18)
	}
	return ret
}

// Convert BigNum to string
func BigStr(a []int64) (ret string) {
	for i := 0; i < len(a); i++ {
		ret = fmt.Sprintf("%018d", a[i]) + ret
	}
	ret = strings.TrimLeft(ret, "0")
	return
}

// Calculate N! and return BigNum
func BigFact(N int) []int64 {
	ret := BigNum("1")
	for i := 2; i <= N; i++ {
		ret = BigMulInt(ret, int64(i))
	}
	return ret
}

// Power n of BigNum a
func BigPow(a []int64, n int64) []int64 {
	ret := BigNum("1")
	s := BigMul(ret, a)
	for n > 0 {
		if n%2 == 1 {
			ret = BigMul(ret, s)
		}
		s = BigMul(s, s)
		n /= 2
	}
	return ret
}

// Length of a BigNum created by BigNum
func BigLen(a []int64) (ret int) {
	var i int
	ret = 18 * (len(a) - 1)
	last := a[len(a)-1]
	for i = 0; last > 0; i++ {
		last /= 10
	}
	ret += i
	return
}

// Sum of all digits of a BigNum
func BigDigSum(a []int64) (ret int64) {
	for _, v := range a {
		ret += DigSum(v)
	}
	return
}
func DigSum(a int64) (ret int64) {
	for a > 0 {
		ret += a % 10
		a /= 10
	}
	return
}

// Compare two BigNum
// 1 if a<b; 0 if a==b; -1 if a>b
func BigLess(a, b []int64) int {
	lena := len(a)
	lenb := len(b)
	if lena < lenb {
		return 1
	} else if lena > lenb {
		return -1
	}
	// lena==lenb
	for i := lena - 1; i >= 0; i-- {
		if a[i] < b[i] {
			return 1
		} else if a[i] > b[i] {
			return -1
		}
	}
	// equal
	return 0
}
