package projecteuler

import (
	"fmt"
	"sort"
)

//////
// Generate perm by recurse
func Perm(list []int, i int, n int) {
	var j = 0
	if i == n {
		fmt.Println(list)
	} else {
		for j = i; j <= n; j++ {
			list[i], list[j] = list[j], list[i]
			Perm(list, i+1, n)
			list[i], list[j] = list[j], list[i]
		}
	}
}
func Comb(n, m int) int {
	var p int = 1
	var a, b int = m, n - m
	if n-m < m {
		a, b = b, a
	}
	for i := 1; i <= a; i++ {
		p += p * b / i
	}
	return p
}

//////
func IsPermutations(a, b int) bool {
	ma := MapNumInts(a)
	mb := MapNumInts(b)
	sort.Ints(ma)
	sort.Ints(mb)
	return IntsEquals(ma, mb)
}

// Permutation n, m of charset string.
// n == len(charset) >= m
// Every result (as string) will processed by PermStrCallback.
func PermStr(charset string, m int, prefix string) {
	if m == 0 {
		PermStrCallback(prefix)
		return
	}
	for i, c := range charset {
		PermStr(charset[:i]+charset[i+1:], m-1, prefix+string(c))
	}
}

var PermStrCallback func(string)

// Combination n, m of charset string.
// n == len(charset) >= m
// Every result (as string) will processed by CombStrCallback.
func CombStr(charset string, m int, prefix string) {
	if m == 0 {
		CombStrCallback(prefix)
		return
	}
	// Select first
	CombStr(charset[1:], m-1, prefix+string(charset[0]))
	// Not select first
	if len(charset) > m {
		CombStr(charset[1:], m, prefix)
	}
}

var CombStrCallback func(string)
