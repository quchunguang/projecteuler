package projecteuler

import (
	"fmt"
	"sort"
	"strconv"
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

//////
var CoinSetCents = map[rune]int{'a': 200, 'b': 100, 'c': 50, 'd': 20, 'e': 10, 'f': 5, 'g': 2, 'h': 1}
var CountPlanCoin int = 0

func PlanCoin(coinset string, remain int, preset string) {
	if remain == 0 {
		// Found one way
		// fmt.Println(preset)
		CountPlanCoin++
		return
	}
	if len(coinset) == 0 {
		return
	}
	c := rune(coinset[0])
	r := remain
	p := preset
	for j := 0; r >= 0; j++ {
		PlanCoin(coinset[1:], r, p)
		r = r - CoinSetCents[c]
		p = p + string(c)
	}
}

//////
var PandigitalProducts []int

func PandigitalProduct(ret string) {
	var sa, sb, sc string
	var a, b, c int

	// There are only 3 possible divides of a given permutation
	//1,3,5
	sa, sb, sc = ret[0:1], ret[1:4], ret[4:]
	a, _ = strconv.Atoi(sa)
	b, _ = strconv.Atoi(sb)
	c, _ = strconv.Atoi(sc)
	if a*b == c {
		fmt.Println(a, ":", b, ":", c)
		InsertUniq(&PandigitalProducts, c)
	}

	//1,4,4
	sa, sb, sc = ret[0:1], ret[1:5], ret[5:]
	a, _ = strconv.Atoi(sa)
	b, _ = strconv.Atoi(sb)
	c, _ = strconv.Atoi(sc)
	if a*b == c {
		fmt.Println(a, ":", b, ":", c)
		InsertUniq(&PandigitalProducts, c)
	}

	//2,3,4
	sa, sb, sc = ret[0:2], ret[2:5], ret[5:]
	a, _ = strconv.Atoi(sa)
	b, _ = strconv.Atoi(sb)
	c, _ = strconv.Atoi(sc)
	if a*b == c {
		fmt.Println(a, ":", b, ":", c)
		InsertUniq(&PandigitalProducts, c)
	}
}
func PermutationPandigital(pandigital string, n int, prefix string) {
	if n == 1 {
		for _, c := range pandigital {
			ret := prefix + string(c)
			PandigitalProduct(ret)
		}
		return
	}
	for i, c := range pandigital {
		PermutationPandigital(pandigital[:i]+pandigital[i+1:], n-1, prefix+string(c))
	}
}

// len(alphabet) >= n
func PermutationStr(alphabet string, n int, prefix string) {
	if n == 1 {
		for _, c := range alphabet {
			ret := prefix + string(c)
			// Do something with ret
			if IsSubStrDiv(ret) {
				subStrDiv, _ := strconv.Atoi(ret)
				SumSubStrDiv += subStrDiv
			}
		}
		return
	}
	for i, c := range alphabet {
		PermutationStr(alphabet[:i]+alphabet[i+1:], n-1, prefix+string(c))
	}
}
