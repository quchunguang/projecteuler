// from: https://projecteuler.net
package projecteuler

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"math/big"
	"os"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"time"
)

func PE1(N int64) int64 {
	var i int64
	var s int64 = 0
	for i = 3; i < N; i++ {
		if i%3 == 0 || i%5 == 0 {
			s += i
		}
	}
	return s
}

//////
func PE2(N int64) int64 {
	var a, b, s int64 = 1, 2, 0
	for b <= N {
		if b%2 == 0 {
			s += b
		}
		a, b = b, a+b
	}
	return s
}

//////
func reduceN(N *int64, p int64) {
	for (*N)%p == 0 {
		(*N) /= p
		// fmt.Printf("%d*", p)
	}
}
func PE3(N int64) int64 {
	// fmt.Printf("%d = ", N)
	var i int64
	var primes []int64
	primes = append(primes, 2)
	reduceN(&N, 2)

	for i = 3; i <= N; i += 2 {
		for _, p := range primes {
			if i%p == 0 {
				goto Out
			}
		}
		primes = append(primes, i)
		reduceN(&N, i)
	Out:
	}
	return primes[len(primes)-1]
}

//////
func PE4() int64 {
	var a, b, N int64
	var s int64 = 0
	for a = 999; a >= 900; a-- {
		for b = 999; b >= 900; b-- {
			N = a * b
			if Palindrome6(N) {
				if N > s {
					s = N
				}
			}
		}
	}
	return s
}

//////
func PE5(N int64) int64 {
	pruducts := make(map[int64]int64)
	pruducts[2] = 1
	var i, j, s, p int64
	for i = 3; i <= N; i++ {
		s = i
		for p = range pruducts {
			for j = 0; s%p == 0; j++ {
				s /= p
			}
			if pruducts[p] < j {
				pruducts[p] = j
			}
		}
		if s > 1 {
			pruducts[s] = 1
		}
	}
	// fmt.Println(pruducts)
	s = 1
	for i = range pruducts {
		s *= Power(i, pruducts[i])
	}
	return s
}

//////
func PE6(N int64) int64 {
	var i, s1, s2 int64
	for i = 1; i <= N; i++ {
		s1 += i
		s2 += i * i
	}
	return s1*s1 - s2
}
func PE6b(N int64) int64 {
	var i, j, s int64
	for i = 1; i <= N; i++ {
		for j = 1; j <= N; j++ {
			if i != j {
				s += i * j
			}
		}
	}
	return s
}
func PE6c(N int64) int64 {
	return N * (N + 1) * (3*N*N - N - 2) / 12
}

//////
func PE7(N int64) int64 {
	var i int64 = 3
	var count int64 = 1
	var primes []int64
	primes = append(primes, 2)
	for i = 3; count < N; i++ {
		for _, p := range primes {
			if i%p == 0 {
				goto Out
			}
		}
		primes = append(primes, i)
		count++
	Out:
	}
	return primes[len(primes)-1]
}

//////
var data string = `73167176531330624919225119674426574742355349194934
96983520312774506326239578318016984801869478851843
85861560789112949495459501737958331952853208805511
12540698747158523863050715693290963295227443043557
66896648950445244523161731856403098711121722383113
62229893423380308135336276614282806444486645238749
30358907296290491560440772390713810515859307960866
70172427121883998797908792274921901699720888093776
65727333001053367881220235421809751254540594752243
52584907711670556013604839586446706324415722155397
53697817977846174064955149290862569321978468622482
83972241375657056057490261407972968652414535100474
82166370484403199890008895243450658541227588666881
16427171479924442928230863465674813919123162824586
17866458359124566529476545682848912883142607690042
24219022671055626321111109370544217506941658960408
07198403850962455444362981230987879927244284909188
84580156166097919133875499200524063689912560717606
05886116467109405077541002256983155200055935729725
71636269561882670428252483600823257530420752963450`

func PE8(N int) int {
	var i, j, s, val, max int
	data = strings.Replace(data, "\n", "", -1)
	for i = 0; i < len(data)-N+1; i++ {
		s = 1
		for j = 0; j < N; j++ {
			val = int(data[i+j]) - 48
			s *= val
		}
		if s > max {
			max = s
		}
	}
	return max
}

//////
func PE9(N int) int {
	var i, j, k int
	for i = 1; i < N/3; i++ {
		for j = i + 1; j < N/2; j++ {
			k = N - i - j
			if i*i+j*j == k*k {
				fmt.Println(i, j, k)
				return i * j * k
			}
		}
	}
	return 0
}

//////
const NMAX = 2e6

func PE10() int64 {
	var i, j, length, upbound, s int64
	var primes [NMAX / 10]int64
	primes[0] = 2
	primes[1] = 3
	length = 2
	for i = 5; i <= NMAX; i += 2 {
		upbound = int64(math.Sqrt(float64(i)))
		for j = 0; primes[j] <= upbound; j++ {
			// for j = 0; primes[j]*primes[j] <= i; j++ {
			if i%primes[j] == 0 {
				goto Out
			}
		}
		primes[length] = i
		length++
	Out:
	}
	s = 0
	for _, i := range primes {
		s += i
	}
	return s
}

////// Sieve of Eratosthenes
func PE10a() int64 {
	var i, j, total, s int64
	var flags [NMAX]bool
	total = int64(math.Sqrt(NMAX)) // put outside, 22ms->14ms !!!
	for i = 2; i < total; i++ {
		if flags[i] {
			continue
		}
		for j = 2; i*j < NMAX; j++ {
			flags[i*j] = true
		}
	}
	for i = 2; i < NMAX; i++ {
		if flags[i] == false {
			s += i
		}
	}
	return s
}

var flags [NMAX]bool

const CORES = 4

func worker(total int64, coreid int64) {
	var i, j int64
	for i = 2 + coreid; i < total; i += CORES {
		if flags[i] {
			continue
		}
		for j = 2; i*j < NMAX; j++ {
			flags[i*j] = true
		}
	}
}
func PE10b() int64 {
	var i, total, s int64
	var done = make(chan bool)

	runtime.GOMAXPROCS(CORES)

	total = int64(math.Sqrt(NMAX))
	go func() {
		worker(total, 0)
		done <- true
	}()
	go func() {
		worker(total, 1)
		done <- true
	}()
	go func() {
		worker(total, 2)
		done <- true
	}()
	go func() {
		worker(total, 3)
		done <- true
	}()
	<-done
	<-done
	<-done
	<-done

	for i = 2; i < NMAX; i++ {
		if flags[i] == false {
			s += i
		}
	}
	return s
}

var data11 = [23][23]int{
	{8, 02, 22, 97, 38, 15, 0, 40, 0, 75, 04, 05, 07, 78, 52, 12, 50, 77, 91, 8, 0, 0, 0},
	{49, 49, 99, 40, 17, 81, 18, 57, 60, 87, 17, 40, 98, 43, 69, 48, 04, 56, 62, 0, 0, 0, 0},
	{81, 49, 31, 73, 55, 79, 14, 29, 93, 71, 40, 67, 53, 88, 30, 03, 49, 13, 36, 65, 0, 0, 0},
	{52, 70, 95, 23, 04, 60, 11, 42, 69, 24, 68, 56, 01, 32, 56, 71, 37, 02, 36, 91, 0, 0, 0},
	{22, 31, 16, 71, 51, 67, 63, 89, 41, 92, 36, 54, 22, 40, 40, 28, 66, 33, 13, 80, 0, 0, 0},
	{24, 47, 32, 60, 99, 03, 45, 02, 44, 75, 33, 53, 78, 36, 84, 20, 35, 17, 12, 50, 0, 0, 0},
	{32, 98, 81, 28, 64, 23, 67, 10, 26, 38, 40, 67, 59, 54, 70, 66, 18, 38, 64, 70, 0, 0, 0},
	{67, 26, 20, 68, 02, 62, 12, 20, 95, 63, 94, 39, 63, 8, 40, 91, 66, 49, 94, 21, 0, 0, 0},
	{24, 55, 58, 05, 66, 73, 99, 26, 97, 17, 78, 78, 96, 83, 14, 88, 34, 89, 63, 72, 0, 0, 0},
	{21, 36, 23, 9, 75, 0, 76, 44, 20, 45, 35, 14, 0, 61, 33, 97, 34, 31, 33, 95, 0, 0, 0},
	{78, 17, 53, 28, 22, 75, 31, 67, 15, 94, 03, 80, 04, 62, 16, 14, 9, 53, 56, 92, 0, 0, 0},
	{16, 39, 05, 42, 96, 35, 31, 47, 55, 58, 88, 24, 0, 17, 54, 24, 36, 29, 85, 57, 0, 0, 0},
	{86, 56, 0, 48, 35, 71, 89, 07, 05, 44, 44, 37, 44, 60, 21, 58, 51, 54, 17, 58, 0, 0, 0},
	{19, 80, 81, 68, 05, 94, 47, 69, 28, 73, 92, 13, 86, 52, 17, 77, 04, 89, 55, 40, 0, 0, 0},
	{04, 52, 8, 83, 97, 35, 99, 16, 07, 97, 57, 32, 16, 26, 26, 79, 33, 27, 98, 66, 0, 0, 0},
	{88, 36, 68, 87, 57, 62, 20, 72, 03, 46, 33, 67, 46, 55, 12, 32, 63, 93, 53, 69, 0, 0, 0},
	{04, 42, 16, 73, 38, 25, 39, 11, 24, 94, 72, 18, 8, 46, 29, 32, 40, 62, 76, 36, 0, 0, 0},
	{20, 69, 36, 41, 72, 30, 23, 88, 34, 62, 99, 69, 82, 67, 59, 85, 74, 04, 36, 16, 0, 0, 0},
	{20, 73, 35, 29, 78, 31, 90, 01, 74, 31, 49, 71, 48, 86, 81, 16, 23, 57, 05, 54, 0, 0, 0},
	{01, 70, 54, 71, 83, 51, 54, 69, 16, 92, 33, 48, 61, 43, 52, 01, 89, 19, 67, 48, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
}

func PE11() int {
	var i, j, max, horizontal, vertical, diagonal, rdiagonal int
	for i = 0; i < 20; i++ {
		for j = 0; j < 20; j++ {
			horizontal = data11[i][j] * data11[i][j+1] * data11[i][j+2] * data11[i][j+3]
			vertical = data11[i][j] * data11[i+1][j] * data11[i+2][j] * data11[i+3][j]
			diagonal = data11[i][j] * data11[i+1][j+1] * data11[i+2][j+2] * data11[i+3][j+3]
			rdiagonal = data11[i][j+3] * data11[i+1][j+2] * data11[i+2][j+1] * data11[i+3][j]
			if horizontal > max {
				max = horizontal
			}
			if vertical > max {
				max = vertical
			}
			if diagonal > max {
				max = diagonal
			}
			if rdiagonal > max {
				max = rdiagonal
			}
		}
	}
	return max
}

//////
func PE12(N int) int {
	for i := 1; ; i++ {
		s := 1
		n := i * (i + 1) / 2
		pfs := PrimeFactors(n)
		for _, v := range pfs {
			s *= v + 1
		}
		if s >= N {
			// fmt.Println(n, s, pfs)
			return n
		}
	}
}

//////
func PE13(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	read := reader.ReadString
	line, err := read('\n')
	if err != nil {
		return ""
	}
	a := BigNum(line)
	for line, err = read('\n'); err == nil; line, err = read('\n') {
		b := BigNum(line)
		a = BigSum(a, b)
	}
	return strconv.FormatInt(a[0], 10)[0:10]
}

//////
func genIterLen(n int64) int64 {
	var length int64 = 1
	for ; n != 1; length++ {
		if n%2 == 0 { //even
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}
	return length
}
func PE14(limit int64) int64 {
	var i, maxlength, longest int64
	for i = limit; i >= 2; i-- {
		length := genIterLen(i)
		if maxlength < length {
			maxlength = length
			longest = i
		}
	}
	return longest
}

//////
func PE15(N int) int {
	return Comb(2*N, N)
}

//////
func PE16(N int) int64 {
	bignum := BigPow(BigNum("2"), int64(N))
	return BigDigSum(bignum)
}

//////
// N must less than (or equal) 1000
func PE17(N int) (ret int) {
	onelen := map[int]int{
		0:  0,
		1:  3,
		2:  3,
		3:  5,
		4:  4,
		5:  4,
		6:  3,
		7:  5,
		8:  5,
		9:  4,
		10: 3,
		11: 6,
		12: 6,
		13: 8,
		14: 8,
		15: 7,
		16: 7,
		17: 9,
		18: 8,
		19: 8,
	}
	tenlen := map[int]int{
		0: 0,
		1: 0,
		2: 6,
		3: 6,
		4: 5,
		5: 5,
		6: 5,
		7: 7,
		8: 6,
		9: 6,
	}
	onename := map[int]string{
		0:  "",
		1:  "one",
		2:  "two",
		3:  "three",
		4:  "four",
		5:  "five",
		6:  "six",
		7:  "seven",
		8:  "eight",
		9:  "nine",
		10: "ten",
		11: "eleven",
		12: "twelve",
		13: "thirteen",
		14: "fourteen",
		15: "fifteen",
		16: "sixteen",
		17: "seventeen",
		18: "eighteen",
		19: "nineteen",
	}
	tenname := map[int]string{
		0: "",
		1: "",
		2: "twenty",
		3: "thirty",
		4: "forty",
		5: "fifty",
		6: "sixty",
		7: "seventy",
		8: "eighty",
		9: "ninety",
	}
	for n := 1; n <= N; n++ {
		m := n
		name := ""
		if m >= 1000 {
			name += "one thousand"
			ret += 3 + 8 // one thousand
			m %= 1000
		}
		if m >= 100 {
			name += onename[m/100]
			ret += onelen[m/100]

			name += " hundred"
			ret += 7 // hundred
			m %= 100
			if m != 0 {
				name += " and "
				ret += 3
			}
		}
		if m >= 20 {
			name += tenname[m/10]
			ret += tenlen[m/10]
			m %= 10
			if m > 0 {
				name += "-"
			}
		}
		if m > 0 {
			name += onename[m]
			ret += onelen[m]
		}
		fmt.Println(name)
	}
	return
}

//////
func PE18(filename string) int {
	return MaxPathSum(filename, 15)
}

//////
func PE19() (ret int) {
	l, _ := time.LoadLocation("Asia/Shanghai")
	for y := 1901; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			t := time.Date(y, time.Month(m), 1, 1, 1, 39, 108924743, l)
			if t.Weekday() == time.Sunday {
				// fmt.Println(t)
				ret++
			}
		}
	}
	return
}
func PE19b() (ret int) {
	week := 1 // 1900.1.1 is Monday
	for y := 1900; y <= 2000; y++ {
		for m := 1; m <= 12; m++ {
			// get days of this month
			days := 31
			if m == 4 || m == 6 || m == 9 || m == 11 {
				days = 30
			}
			if m == 2 {
				days = 28
				if y%4 == 0 && y%100 != 0 || y%400 == 0 {
					days = 29
				}
			}

			// iterator days of month
			for d := 1; d <= days; d++ {
				if y == 1901 && m == 1 && d == 1 {
					ret = 0
				}
				if d == 1 && week == 0 {
					ret++
				}
				week = (week + 1) % 7
			}
		}
	}
	return
}

//////
func PE20(N int) (ret int) {
	s := BigStr(BigFact(N))
	for _, c := range s {
		ret += int(c - 0x30)
	}
	return ret
}

//////
func PE21(N int) (sum int) {
	for a := 1; a <= N; a++ {
		b := SumInts(Factors(a))
		if SumInts(Factors(b)) == a && a != b {
			// fmt.Println(a, b)
			sum += a
		}
	}
	return
}

//////
func PE22(filename string) (ret int) {
	words := CSW(filename)
	sort.Strings(words)

	for i, w := range words {
		ret += (i + 1) * ScoreWord(w)
	}
	return
}

//////
func PE23() (ret int) {
	GenAbundants(28123)
	// All integers > 28123 can be written as the sum of two abundant numbers.
	for i := 1; i <= 28123; i++ {
		for _, ab := range Abundants {
			if ab >= i {
				// Non-abundant sum found
				ret += i
				break
			}
			if InInts(Abundants, i-ab) {
				break
			}
		}
	}
	return
}

//////
func PE24(n int) int {
	const length = 10 // 0 .. (length-1)
	n = n - 1         // nth choices will pass over n-1 choices
	var s, k [length]int
	var f []int
	var ret string

	for i := 0; i < length; i++ {
		f = append(f, i)
	}

	s[length-1] = 1 // 0! = 1, s[9-i]=i!
	for i := 1; i < length; i++ {
		s[length-i-1] = s[length-i] * i
	}
	for i := 0; i < length; i++ {
		if n < s[i] {
			continue
		}
		k[i] = n / s[i]
		n = n % s[i]
	}

	for i := 0; i < length; i++ {
		ret += strconv.Itoa(f[k[i]])
		f = DeleteIndex(f, k[i])
	}
	reti, _ := strconv.Atoi(ret)
	return reti
}

//////
func PE25(n int) (ret int) {
	fnp := BigNum("1")  //f1
	fn := BigNum("1")   //f2
	for i := 3; ; i++ { //fn
		fnp, fn = fn, BigSum(fnp, fn)
		if BigLen(fn) == n {
			ret = i
			return
		}
	}
}

//////
func PE26(N int) (ret int) {
	var fullprimes []int

	var dms []DivMod
	var maxlen int
	for n := 2; n < N; n++ {
		m := 1
		dms = dms[:0]
		for {
			m *= 10
			if m%n == 0 {
				break
			}
			dm := DivMod{m / n, m % n}
			if index := InDivMods(dms, dm); index >= 0 {
				length := len(dms) - index
				if length == n-1 {
					fullprimes = append(fullprimes, n)
				}
				if length > maxlen {
					maxlen = length
					ret = n
					// fmt.Println(n, maxlen, dms)
				}
				break
			}
			dms = append(dms, dm)
			m %= n
		}
	}

	// Question
	// * For 1/n, max recurring cycle length is n-1.
	// * All max recurring cycle number, are primes.
	// * Not all primes are recurring cycle numbers.
	// color reference ~/bin/colorcat.sh
	GenPrimes(N)
	var r, yy int
	for _, v := range primes {
		if InInts(fullprimes, v) {
			fmt.Printf("%s%4d%s ", CRR, v, CRD)
			yy++
		} else {
			fmt.Printf("%4d ", v)
		}

		r++
		if r%32 == 0 {
			fmt.Printf(" ρ = %4.2f\n", float64(yy)/float64(r))
		}
	}
	fmt.Println()

	return
}

//////
func PE27() (ret int) {
	max := 0
	// ret = QuadraticPrimes(1, 41)
	for a := -999; a < 1000; a++ {
		for b := -999; b < 1000; b++ {
			s := QuadraticPrimes(a, b)
			if s > max {
				max = s
				ret = a * b
				fmt.Println(a, b, s)
			}
		}
	}
	return
}

//////
func PE28(N int) (ret int) {
	ret = 1
	for n := 3; n <= N; n += 2 {
		ret += 4*n*n - 6*n + 6
	}
	return
}

//////
func PE29(N int) (ret int) {
	var data []string
	for i := 2; i <= N; i++ {
		for j := 2; j <= N; j++ {
			data = append(data, BigStr(BigPow(BigNum(strconv.Itoa(i)), int64(j))))
		}
	}
	sort.Sort(SortIntStr(data))
	before := ""
	for _, v := range data {
		if before != v {
			ret++
			before = v
		}
	}
	return
}

//////
func genlimit(n int) int {
	for i := 1; ; i++ {
		if PowInt(9, n)*i < PowInt(10, i-1) {
			return PowInt(10, i-1)
		}
	}
}
func PE30(m int) int {
	var ret []int
	for i := 10; i < genlimit(m); i++ {
		nlist := MapNumInts(i)
		if i == PowSum(nlist, m) {
			ret = append(ret, i)
		}
	}
	return SumInts(ret)
}

//////
func PE31(N int) (ret int) {
	coinset := "abcdefgh" // This will coast 2.5s
	// coinset := "hgfedcba" // This will coast 4.2s
	PlanCoin(coinset, N, "")
	ret = CountPlanCoin
	return
}

//////
func PE32() (ret int) {
	// Permutation (9,9) of "123456789"
	PermutationPandigital("123456789", 9, "")
	ret = SumInts(PandigitalProducts)

	return
}

//////
func PE33() (ret int) {
	var sa, sb string
	var e bool
	mul := big.NewRat(1, 1)
	for a := 10; a < 100; a++ {
		for b := a + 1; b < 100; b++ {
			sa = strconv.Itoa(a)
			sb = strconv.Itoa(b)
			sa, sb, e = CancelStr(sa, sb)
			if e {
				continue
			}
			ca, _ := strconv.Atoi(sa)
			cb, _ := strconv.Atoi(sb)
			rat1 := big.NewRat(int64(a), int64(b))
			rat2 := big.NewRat(int64(ca), int64(cb))
			if rat1.Cmp(rat2) == 0 {
				// fmt.Println(a, b)
				mul.Mul(mul, rat1)
			}
		}
	}
	ret, _ = strconv.Atoi(mul.Denom().String())
	return
}

////// notice: this function will never stop
func PE34() (ret int) {
	var prod [10]int
	prod[0] = 1
	mul := 1
	for i := 1; i <= 9; i++ {
		mul *= i
		prod[i] = mul
	}
	fmt.Println(big.NewInt(1).MulRange(1, 9))
	fmt.Println(prod)

	sum := 0
	for i := 10; ; i++ {
		nums := DigNums(i)
		s := 0
		for _, c := range nums {
			s += prod[c]
		}
		if s == i {
			sum += i
			fmt.Println(i, sum)
		}
	}
	return
}

//////
func PE35(N int) (ret int) {
	GenPrimes(N)
	for _, p := range primes {
		sp := strconv.Itoa(p)
		for i := 1; i < len(sp); i++ {
			rsp := sp[i:] + sp[:i] // Circular numbers at index i
			rp, _ := strconv.Atoi(rsp)
			if !InInts(primes, rp) {
				goto NEXT
			}
		}
		// Found!!!
		fmt.Println(p)
		ret++
	NEXT:
	}
	return
}

//////
func PE36(n int) int {
	var ret []int
	for i := 1; i < n; i++ {
		if IsPalindromeBytes([]byte(strconv.Itoa(i))) &&
			IsPalindromeBytes([]byte(strconv.FormatInt(int64(i), 2))) {
			ret = append(ret, i)
		}
	}
	return SumInts(ret)
}

//////
func PE37() (ret int) {
	var count int = 0
	var N int = 1e3
	GenPrimes(N)
	for i := 0; ; i++ {
		// If not enough primes, double search range
		if i == len(primes) {
			N *= 2
			GenPrimes(N)
		}

		// Check if all cuts of primes[i] are also primes
		sp := strconv.Itoa(primes[i])
		for i := 1; i < len(sp); i++ {
			rp1, _ := strconv.Atoi(sp[i:])
			rp2, _ := strconv.Atoi(sp[:i])
			if !InInts(primes, rp1) || !InInts(primes, rp2) {
				goto NEXT
			}
		}

		// 2, 3, 5, and 7 are not considered to be truncatable primes.
		if primes[i] <= 7 {
			goto NEXT
		}
		// Found!!!
		fmt.Println(primes[i])
		count++
		ret += primes[i]
		if count == 11 { // Found all
			break
		}
	NEXT:
	}
	return
}

//////
func PE38() (ret int) {
	for a := 1; a < 1e4; a++ {
		s := ""
		for n := 1; ; n++ {
			s += strconv.Itoa(a * n)
			if len(s) < 9 {
				continue
			} else if len(s) > 9 {
				break
			} else {
				// len(s)==9
				ss := strings.Split(s, "")
				sort.Strings(ss)
				if strings.Join(ss, "") == "123456789" {
					fmt.Println(a, n, s)
					num, _ := strconv.Atoi(s)
					if num > ret {
						ret = num
					}
				}
			}
		}
	}
	return
}

//////
func PE39(n int) int {
	var a, b, c, p int
	var sum, max, maxp int
	for p = 3; p <= n; p++ {
		sum = 0
		for c = p / 3; c <= p/2; c++ {
			for a = 1; a < p/3; a++ {
				b = p - c - a
				if b <= 0 {
					continue
				}
				if RightTri(a, b, c) {
					sum++
				}
			}
		}
		if sum > max {
			max = sum
			maxp = p
		}
	}
	fmt.Println("summax", max)
	return maxp
}

//////
func PE40(N int) (ret int) {
	ret = 1
	for i := 1; i <= N; i *= 10 {
		ret *= Champernowne(i)
	}
	return
}

//////
// NOTE: This program will never stop!!!
func PE41() (ret int) {
	var N int = 1e3
	numlist := "123456789"
	GenPrimes(N)
	for i := 0; ; i++ {
		// If not enough primes, double search range
		if i == len(primes) {
			N *= 2
			GenPrimes(N)
		}

		// Check if the prime is pandigital
		sp := strconv.Itoa(primes[i])
		ss := strings.Split(sp, "")
		sort.Strings(ss)
		if strings.Join(ss, "") == numlist[:len(sp)] {
			// Found!!!
			fmt.Println(primes[i])
		}
	}
	return
}

//////
func PE42(filename string) (ret int) {
	words := CSW(filename)
	for _, w := range words {
		if IsTriangle(SumAscii(w)) {
			ret++
		}
	}
	return
}

//////
func PE43() (ret int) {
	PermutationStr("0123456789", 10, "")
	ret = SumSubStrDiv
	return
}

//////
func PE44() (ret int) {
	ret = 1 << 30
	for i := 2; ; i++ {
		GenPentagons(Pentagonal(i))
		for j := 1; j < i; j++ {
			pi := Pentagons[i-1]
			pj := Pentagons[j-1]
			if InInts(Pentagons, pi-pj) {
				GenPentagons(pi + pj)
				if InInts(Pentagons, pi+pj) {
					if pi-pj < ret {
						ret = pi - pj
						return
						// fmt.Println(ret)
					}
				}
			}
		}
	}
}

//////
func PE45() int {
	var t, p, h int = 1, 1, 1
	var T, P, H int
	for {
		T = Triangle(t)
		P = Pentagonal(p)
		H = Hexagonal(h)
		if T == P && T == H {
			fmt.Println(t, p, h, T)
		}
		switch {
		case T <= P && T <= H:
			t++
		case P <= T && P <= H:
			p++
		default:
			h++
		}
	}
	return 0
}

//////
var oddcomposites = []int{9, 15}

func PE46() int {
	for i := 17; ; i += 2 {
		upbound := int(math.Sqrt(float64(i)))
		for j := 0; primes[j] <= upbound; j++ {
			if i%primes[j] == 0 {
				oddcomposites = append(oddcomposites, i)
				if !IsOddComposites(i) {
					return i
				}
				goto NEXT
			}
		}
		primes = append(primes, i)
	NEXT:
	}
}

//////
func PE47(n int) int {
	ok := 0
	for i := 4; ; i++ {
		if len(PrimeFactors(i)) != n {
			ok = 0
			continue
		}
		ok++
		if ok == n {
			return i - n + 1
		}
	}
}

//////
func PE48(n int) int {
	var sum int
	for i := 1; i <= n; i++ {
		sum += PowerTail(i, i)
	}
	return sum % base
}

//////
func PE49() (ret []string) {
	GenPrimes(10000)
	for i := 1001; i < 10000; i += 2 {
		j := i + 3330
		k := j + 3330
		if !InInts(primes, i) || !InInts(primes, j) || !InInts(primes, k) {
			continue
		}
		if IsPermutations(i, j) && IsPermutations(i, k) {
			ret = append(ret, strconv.Itoa(i)+strconv.Itoa(j)+strconv.Itoa(k))
		}
	}
	return
}

//////
func PE50(n int) (ret int) {
	GenPrimes(n)
	SumPrimes(n)
	max := sort.SearchInts(sumprimes, n)
	for i := max; ; i-- {
		for j := 0; j <= max-i; j++ {
			if j == 0 {
				if ret = sumprimes[i-1]; InInts(primes, ret) {
					return
				}
			} else {
				if ret = sumprimes[i+j-1] - sumprimes[j-1]; InInts(primes, ret) {
					return
				}
			}
		}
	}
}
