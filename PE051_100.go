package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

//////
func PE52() int {
	return 142857
}

//////
func PE53(N int) (ret int) {
	for n := 1; n <= 100; n++ {
		for r := 1; r <= n; r++ {
			if Comb(n, r) > N {
				fmt.Println(n, r, Comb(n, r))
				ret += n - r - r + 1
				break
			}
		}
	}
	return
}

//////
func PE55(N int) (ret int) {
	for n := 1; n < N; n++ {
		if IsLychrel(n) {
			ret++
		}
	}
	return
}

//////
func PE56(N int) (ret int) {
	for a := 1; a < N; a++ {
		for b := 1; b < N; b++ {
			x := BigPow(BigNum(strconv.Itoa(a)), int64(b))
			s := int(BigDigSum(x))
			if s > ret {
				ret = s
			}
		}
	}
	return
}

//////
func PE57S(N int) (ret int) {
	var a, b int = 1, 2
	for i := 2; i <= N; i++ {
		a, b = b, 2*b+a
		// fmt.Println(a+b, "/", b, len(strconv.Itoa(a+b)), len(strconv.Itoa(b)))
		if len(strconv.Itoa(a+b)) > len(strconv.Itoa(b)) {
			ret++
		}
	}
	return
}
func PE57(N int) (ret int) {
	a := BigNum("1")
	b := BigNum("2")
	for i := 2; i <= N; i++ {
		c := b
		b = BigSum(BigMulInt(b, 2), a)
		a = c
		if BigLen(BigSum(a, b)) > BigLen(b) {
			ret++
		}
	}
	// fmt.Println(BigStr(BigSum(a, b)), BigStr(b))
	return
}

//////
// Run time about 1 hour
func PE58() int {
	var n, step int
	var yes, no int
	var tl, tr, bl, br int // number of four angle
	no = 1                 // 1 is not prime
	n = 1
	for i := 3; ; i += 2 {
		maxprimes := primes[len(primes)-1]
		if maxprimes < i*i {
			GenPrimes(2 * maxprimes)
		}
		step = i - 1
		tr = n + step
		tl = tr + step
		bl = tl + step
		br = bl + step
		n = br
		if InInts(primes, tr) {
			yes++
		} else {
			no++
		}
		if InInts(primes, tl) {
			yes++
		} else {
			no++
		}
		if InInts(primes, bl) {
			yes++
		} else {
			no++
		}
		if InInts(primes, br) {
			yes++
		} else {
			no++
		}
		fmt.Println(i, yes, no+yes)
		if 10*yes < yes+no {
			return i
		}
	}
}

//////
func Decrypt(key []byte, cipher []byte) []byte {
	ret := make([]byte, len(cipher))
	for i := 0; i < len(cipher); i += 3 {
		ret[i] = cipher[i] ^ key[0]
		if i+1 < len(cipher) {
			ret[i+1] = cipher[i+1] ^ key[1]
		}
		if i+2 < len(cipher) {
			ret[i+2] = cipher[i+2] ^ key[2]
		}
	}
	return ret
}
func GuessEnglishText(text []byte) (ret int) {
	wordlist := []string{" the ", " be ", " to ", " of ", " and ", " a ", " in ", " that ", " have "}
	for _, word := range wordlist {
		ret += strings.Count(string(text), word)
	}
	return
}
func PE59(cipherfile string) (ret int) {
	cipher := CSV(cipherfile)
	key := []byte("aaa")
	for {
		// Ten more words(in middle of a sentence) found in text.
		if GuessEnglishText(Decrypt(key, cipher)) >= 10 {
			break
		}
		// next key
		key[2]++
		if key[2] > byte('z') {
			key[2] = byte('a')
			key[1]++
		}
		if key[1] > byte('z') {
			key[1] = byte('a')
			key[0]++
		}
		if key[0] > byte('z') {
			// NOT FOUND!!!
			fmt.Println("NOT FOUND!!!")
			return
		}
	}
	// FOUND!!!
	// fmt.Println(string(key), GuessEnglishText(Decrypt(key, cipher)), string(Decrypt(key, cipher)))
	text := Decrypt(key, cipher)
	for _, c := range text {
		ret += int(c)
	}
	return
}

//////
var polygonals []int

func Polygonals(n int) (ret [6]int) {
	ret[0] = n * (n + 1) / 2   // Triangle   P(3,n)
	ret[1] = n * n             // Square     P(4,n)
	ret[2] = n * (3*n - 1) / 2 // Pentagonal P(5,n)
	ret[3] = n * (2*n - 1)     // Hexagonal  P(6,n)
	ret[4] = n * (5*n - 3) / 2 // Heptagonal P(7,n)
	ret[5] = n * (3*n - 2)     // Octagonal  P(8,n)
	return
}

func GenPolygonals4() {
	for i := 1; ; i++ {
		s := Polygonals(i)
		if s[0] > 1e4 {
			return
		}
		for j := 0; j < 6; j++ {
			if s[j] >= 1e3 && s[j] < 1e4 {
				InsertUniq(&polygonals, s[j])
			}
		}
	}
}
func InsertUniq(slice *[]int, value int) {
	for _, v := range *slice {
		if v == value {
			return
		}
	}
	*slice = append(*slice, value)
}
func InsertUniqC6(data *[][6]int, item [6]int) {
	var min int = 1e4
	var index int
	var res [6]int
	for i := 0; i < 6; i++ {
		if item[i] < min {
			min = item[i]
			index = i
		}
	}
	for i := 0; i < 6; i++ {
		res[i] = item[(i+index)%6]
	}
	for _, v := range *data {
		for i := 0; i < 6; i++ {
			if v[i] != res[i] {
				goto NEXT
			}
		}
		return
	NEXT:
	}
	*data = append(*data, res)
}
func Permutation(slice []int, polymap map[int][]int) (res [][6]int) {
	for _, v1 := range slice {
		for _, v2 := range slice {
			if v2 == v1 || !Adj(polymap, v1, v2) {
				continue
			}
			for _, v3 := range slice {
				if v3 == v1 || v3 == v2 || !Adj(polymap, v2, v3) {
					continue
				}
				for _, v4 := range slice {
					if v4 == v1 || v4 == v2 || v4 == v3 || !Adj(polymap, v3, v4) {
						continue
					}
					for _, v5 := range slice {
						if v5 == v1 || v5 == v2 || v5 == v3 || v5 == v4 || !Adj(polymap, v4, v5) {
							continue
						}
						for _, v6 := range slice {
							if v6 == v1 || v6 == v2 || v6 == v3 || v6 == v4 || v6 == v5 || !Adj(polymap, v5, v6) {
								continue
							}
							if Adj(polymap, v6, v1) {
								InsertUniqC6(&res, [6]int{v1, v2, v3, v4, v5, v6})
							}
						}
					}
				}
			}
		}
	}
	for _, v := range res {
		fmt.Println(v)
	}
	return
}
func Adj(polymap map[int][]int, v1, v2 int) bool {
	for _, v := range polymap[v1] {
		if v == v2 {
			return true
		}
	}
	return false
}
func DeepSet(slice map[int][]int, seed int, deep int) []int {
	var from = make([]int, 0)
	var to = make([]int, 0)
	from = append(from, seed)
	for i := 0; i < deep; i++ {
		for _, value := range from {
			for _, v := range slice[value] {
				InsertUniq(&to, v)
			}
		}
		from = to
		to = make([]int, 0)
	}
	return from
}
func PE61_wrong() (ret int) {
	GenPolygonals4()

	// Create adj map from polygonals
	polymap := make(map[int][]int)
	var keys []int
	for _, v := range polygonals {
		polymap[v/100] = append(polymap[v/100], v%100)
		InsertUniq(&keys, v/100)
	}

	// Display by order
	sort.Ints(keys)
	// for _, k := range keys {
	//  fmt.Println(k, polymap[k])
	// }

	// for k, _ := range polymap {
	//  set := DeepSet(polymap, k,6)
	//  for _, v := range set {
	//      if v == k {
	//          fmt.Println("FOUND!!!")
	//      }
	//  }
	// }
	Permutation(keys, polymap)
	return
}

//////
func PE62() (ret int) {
	var i, n int
	digmap := make(map[string][]int)
	for i = 1e3; i <= 1e4-1; i++ {
		n = i * i * i
		dignums := strconv.Itoa(n)
		dignums = SortString(dignums)
		digmap[dignums] = append(digmap[dignums], i)
	}

	ret = 1e5
	for _, v := range digmap {
		if len(v) >= 5 {
			for _, item := range v {
				if item < ret {
					ret = item
					fmt.Println(v)
				}
			}
		}
	}
	ret = ret * ret * ret
	return
}

//////
func PE63() (ret int) {
	for a := 1; a <= 9; a++ {
		for n := 1; ; n++ {
			num := float64(n) * (math.Log10(float64(a)) - 1)
			if num >= 0 || num < -1.0 {
				break
			}
			fmt.Println(num, a, n)
			ret++
		}
	}
	return
}

//////
func PE67(filename string) int {
	return MaxPathSum(filename, 100)
}

//////
func PE69_slow(N int) (ret int) {
	var max float64 = 0
	for n := 2; n < N; n++ {
		ratephi := RatePhi(n)
		if ratephi > max {
			max = ratephi
			ret = n
		}
	}
	fmt.Println(ret, max)
	return
}

//////
func PE97() (ret int) {
	// ret = 2^7830457
	ret = 1
	for i := 0; i < 7830457; i++ {
		ret *= 2
		if ret > 1e10 {
			ret = ret % 1e10
		}
	}
	// ret = ret*28433 + 1
	ret = ret*28433 + 1
	ret = ret % 1e10
	return
}

//////
func PE99(filename string) (ret int) {
	var maxlog float64 = 0.0

	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for i := 1; ; i++ {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		pair := strings.Split(line, ",")
		base, _ := strconv.Atoi(strings.TrimSpace(pair[0]))
		Power, _ := strconv.Atoi(strings.TrimSpace(pair[1]))
		res := float64(Power) * math.Log(float64(base))
		if res > maxlog {
			maxlog = res
			ret = i
		}
	}
	return
}
