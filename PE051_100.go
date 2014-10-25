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

// Problem 51 - Prime digit replacements
//
// By replacing the 1st digit of the 2-digit number *3, it turns out that six of the nine possible values: 13, 23, 43, 53, 73, and 83, are all prime.
//
// By replacing the 3rd and 4th digits of 56**3 with the same digit, this 5-digit number is the first example having seven primes among the ten generated numbers, yielding the family: 56003, 56113, 56333, 56443, 56663, 56773, and 56993. Consequently 56003, being the first member of this family, is the smallest prime with this property.
//
// Find the smallest prime which, by replacing part of the number (not necessarily adjacent digits) with the same digit, is part of an eight prime value family.
func PE51() (ret int) {
	// No result  found in 5-digit numbers

	// Finding in 6-digit numbers
	GenPrimes(1e6)
	CombStrCallback = checkPrimeSelect
	for i := 1; i <= 6; i++ {
		CombStr("123456", i, "")
	}
	ret = pe51ret
	return
}

var pe51ret int

func checkPrimeSelect(s string) {
	selects := TrueDigits(s, 6)
	strmap := make(map[string]int)
	for _, p := range primes {
		if p < 100000 {
			continue
		}
		sp := strconv.Itoa(p)
		check, key := SplitDigits(sp, selects)
		if IsSameStr(check) {
			strmap[key]++
		}
	}
	for k, v := range strmap {
		if v >= 8 {
			// Found!!! Generate result.
			for i := 0; i <= 9; i++ {
				res := ""
				u := 0
				for j := 0; j < 6; j++ {
					if selects[j] {
						res += strconv.Itoa(i)
					} else {
						res += string(k[u])
						u++
					}
				}
				ires, _ := strconv.Atoi(res)
				if InInts(primes, ires) {
					// Find first (smallest) prime and return.
					pe51ret = ires
					return
				}
			}
		}
	}
}

// Problem 52 - Permuted multiples
//
// It can be seen that the number, 125874, and its double, 251748, contain exactly the same digits, but in a different order.
//
// Find the smallest positive integer, x, such that 2x, 3x, 4x, 5x, and 6x, contain the same digits.
func PE52() int {
	for i := 1; ; i++ {
		dns := DigNums(i)
		sort.Ints(dns)
		for j := 2; j <= 6; j++ {
			dns2 := DigNums(i * j)
			sort.Ints(dns2)
			if !EqualInts(dns, dns2) {
				goto NEXT
			}
		}
		// Found!!!
		return i
	NEXT:
	}
	return 0
}

// Problem 53 - Combinatoric selections
//
// There are exactly ten ways of selecting three from five, 12345:
//
// 123, 124, 125, 134, 135, 145, 234, 235, 245, and 345
//
// In combinatorics, we use the notation, 5C3 = 10.
//
// In general,
// nCr =
// n!
// r!(n−r)!
// 	,where r ≤ n, n! = n×(n−1)×...×3×2×1, and 0! = 1.
//
// It is not until n = 23, that a value exceeds one-million: 23C10 = 1144066.
//
// How many, not necessarily distinct, values of  nCr, for 1 ≤ n ≤ 100, are greater than one-million?
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

// Problem 54 - Poker hands
//
// In the card game poker, a hand consists of five cards and are ranked, from lowest to highest, in the following way:
//     Rank
//     0 High Card: Highest value card.
//     1 One Pair: Two cards of the same value.
//     2 Two Pairs: Two different pairs.
//     3 Three of a Kind: Three cards of the same value.
//     4 Straight: All cards are consecutive values.
//     5 Flush: All cards of the same suit.
//     6 Full House: Three of a kind and a pair.
//     7 Four of a Kind: Four cards of the same value.
//     8 Straight Flush: All cards are consecutive values of same suit.
//     9 Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
//
// The cards are valued in the order:
// 2, 3, 4, 5, 6, 7, 8, 9, 10, Jack, Queen, King, Ace.
//
// If two players have the same ranked hands then the rank made up of the highest value wins; for example, a pair of eights beats a pair of fives (see example 1 below). But if two ranks tie, for example, both players have a pair of queens, then highest cards in each hand are compared (see example 4 below); if the highest cards tie then the next highest cards are compared, and so on.
//
// Consider the following five hands dealt to two players:
// Hand	 	Player 1	 		Player 2	 			Winner
// 1	 	5H 5C 6S 7S KD		2C 3S 8S 8D TD			Player 2
// 			Pair of Fives		Pair of Eights
// 2	 	5D 8C 9S JS AC		2C 5C 7D 8S QH			Player 1
// 			Highest card Ace	Highest card Queen
// 3	 	2D 9C AS AH AC		3D 6D 7D TD QD			Player 2
// 			Three Aces			Flush with Diamonds
// 4	 	4D 6S 9H QH QC		3D 6D 7H QD QS			Player 1
// 			Pair of Queens		Pair of Queens
// 			Highest card Nine	Highest card Seven
// 5	 	2H 2D 4C 4D 4S		3C 3D 3S 9S 9D			Player 1
// 			Full House			Full House
// 			With Three Fours	with Three Threes
//
// The file, poker.txt, contains one-thousand random hands dealt to two players. Each line of the file contains ten cards (separated by a single space): the first five are Player 1's cards and the last five are Player 2's cards. You can assume that all hands are valid (no invalid characters or repeated cards), each player's hand is in no specific order, and in each hand there is a clear winner.
//
// How many hands does Player 1 win?
func PE54(filename string) (ret int) {
	table := CSWs(filename, " ")
	for _, hand := range table {
		p1 := hand[:5]
		p2 := hand[5:]
		if p1Win5Card(p1, p2) {
			ret++
		}
	}
	return
}

func p1Win5Card(p1, p2 []string) bool {
	r1, v1, h1 := rank5Card(p1)
	r2, v2, h2 := rank5Card(p2)
	if r1 > r2 {
		return true
	} else if r1 < r2 {
		return false
	} else if v1 > v2 {
		return true
	} else if v1 < v2 {
		return false
	} else if h1 > h2 {
		return true
	} else if h1 < h2 {
		return false
	}
	fmt.Println("Cannot judge!!!")
	fmt.Println(p1, " <=> ", p2)

	return true
}

var cardValue = map[byte]int{'2': 2, '3': 3, '4': 4, '5': 5, '6': 6, '7': 7, '8': 8, '9': 9, 'T': 10, 'J': 11, 'Q': 12, 'K': 13, 'A': 14}

func rank5Card(card []string) (rank, value, high int) {
	var straight, flush, four, three, pair, pairthree, pair2 int
	sortCard(card)

	flush = testCardFlush(card)
	straight = testCardStraight(card)

	if flush > 0 && straight > 0 {
		if straight == 14 {
			rank = 9 // Royal Flush: Ten, Jack, Queen, King, Ace, in same suit.
		} else {
			rank = 8 // Straight Flush: All cards are consecutive values of same suit.
		}
		value = flush
		high = flush
		return
	}

	four, high = testCardFour(card)
	if four > 0 {
		rank = 7 // Four of a Kind: Four cards of the same value.
		value = four
		return
	}

	three, pairthree, high = testCardThree(card)
	if pairthree > 0 {
		rank = 6 // Full House: Three of a kind and a pair.
		value = three
		high = pair
		return
	}

	if flush > 0 {
		rank = 5 // Flush: All cards of the same suit.
		value = flush
		high = flush
		return
	}

	if straight > 0 {
		rank = 4 // Straight: All cards are consecutive values.
		value = straight
		high = straight
		return
	}

	if three > 0 {
		rank = 3 // Three of a Kind: Three cards of the same value.
		value = three
		return
	}

	pair, pair2, high = testCardPair(card)
	if pair2 > 0 {
		rank = 2 // Two Pairs: Two different pairs.
		value = pair2
		high = pair // BAD: pair2, pair all same, no chance compare the last one
		return
	}

	if pair > 0 {
		rank = 1 // One Pair: Two cards of the same value.
		value = pair
		return
	}

	rank = 0 // High Card: Highest value card.
	value = cardValue[card[4][0]]
	high = cardValue[card[4][0]]
	return
}

// Test four
func testCardFour(card []string) (four, high int) {
	for i := 0; i < 2; i++ {
		if cardValue[card[i][0]] == cardValue[card[i+3][0]] {
			four = cardValue[card[i+3][0]]
			if i == 0 {
				high = cardValue[card[4][0]]
			} else {
				high = cardValue[card[0][0]]
			}
		}
	}
	return
}

// Test three/pairthree
func testCardThree(card []string) (three, pairthree, high int) {
	for i := 0; i < 3; i++ {
		if cardValue[card[i][0]] == cardValue[card[i+2][0]] {
			three = cardValue[card[i+2][0]]
			if i == 0 && cardValue[card[3][0]] == cardValue[card[4][0]] {
				pairthree = cardValue[card[4][0]]
			} else if i == 2 && cardValue[card[0][0]] == cardValue[card[1][0]] {
				pairthree = cardValue[card[1][0]]
			}
		}
	}
	for i := 4; i >= 0; i-- {
		if cardValue[card[i][0]] != three {
			high = cardValue[card[i][0]]
		}
	}
	return
}

// Test pair/pair2
func testCardPair(card []string) (pair, pair2, high int) {
	for i := 0; i < 4; i++ {
		if cardValue[card[i][0]] == cardValue[card[i+1][0]] {
			if pair == 0 {
				pair = cardValue[card[i+1][0]]
			} else {
				pair2 = cardValue[card[i+1][0]]
			}
		}
	}
	for i := 4; i >= 0; i-- {
		if cardValue[card[i][0]] != pair && cardValue[card[i][0]] != pair2 {
			high = cardValue[card[i][0]]
			break
		}
	}
	return
}

// Test straight
func testCardStraight(card []string) (straight int) {
	var i int
	for i = 1; i < 5 && cardValue[card[i][0]] == cardValue[card[0][0]]+i; i++ {
	}
	if i == 5 {
		straight = cardValue[card[4][0]]
	}
	return
}

// Test flush
func testCardFlush(card []string) (flush int) {
	var i int
	for i = 1; i < 5 && card[i][1] == card[0][1]; i++ {
	}
	if i == 5 {
		flush = cardValue[card[4][0]]
	}
	return
}

func sortCard(card []string) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if cardValue[card[i][0]] > cardValue[card[j][0]] {
				card[i], card[j] = card[j], card[i]
			}
		}
	}
}

// Problem 55 - Lychrel numbers
//
// If we take 47, reverse and add, 47 + 74 = 121, which is palindromic.
//
// Not all numbers produce palindromes so quickly. For example,
//
// 349 + 943 = 1292,
// 1292 + 2921 = 4213
// 4213 + 3124 = 7337
//
// That is, 349 took three iterations to arrive at a palindrome.
//
// Although no one has proved it yet, it is thought that some numbers, like 196, never produce a palindrome. A number that never forms a palindrome through the reverse and add process is called a Lychrel number. Due to the theoretical nature of these numbers, and for the purpose of this problem, we shall assume that a number is Lychrel until proven otherwise. In addition you are given that for every number below ten-thousand, it will either (i) become a palindrome in less than fifty iterations, or, (ii) no one, with all the computing power that exists, has managed so far to map it to a palindrome. In fact, 10677 is the first number to be shown to require over fifty iterations before producing a palindrome: 4668731596684224866951378664 (53 iterations, 28-digits).
//
// Surprisingly, there are palindromic numbers that are themselves Lychrel numbers; the first example is 4994.
//
// How many Lychrel numbers are there below ten-thousand?
//
// NOTE: Wording was modified slightly on 24 April 2007 to emphasise the theoretical nature of Lychrel numbers.
func PE55(N int) (ret int) {
	for n := 1; n < N; n++ {
		if IsLychrel(n) {
			ret++
		}
	}
	return
}

// Problem 56 - Powerful digit sum
//
// A googol (10^100) is a massive number: one followed by one-hundred zeros; 100^100 is almost unimaginably large: one followed by two-hundred zeros. Despite their size, the sum of the digits in each number is only 1.
//
// Considering natural numbers of the form, ab, where a, b < 100, what is the maximum digital sum?
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

// Problem 57 - Square root convergents
//
// It is possible to show that the square root of two can be expressed as an infinite continued fraction.
//
// √ 2 = 1 + 1/(2 + 1/(2 + 1/(2 + ... ))) = 1.414213...
//
// By expanding this for the first four iterations, we get:
//
// 1 + 1/2 = 3/2 = 1.5
// 1 + 1/(2 + 1/2) = 7/5 = 1.4
// 1 + 1/(2 + 1/(2 + 1/2)) = 17/12 = 1.41666...
// 1 + 1/(2 + 1/(2 + 1/(2 + 1/2))) = 41/29 = 1.41379...
//
// The next three expansions are 99/70, 239/169, and 577/408, but the eighth expansion, 1393/985, is the first example where the number of digits in the numerator exceeds the number of digits in the denominator.
//
// In the first one-thousand expansions, how many fractions contain a numerator with more digits than denominator?
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

// Problem 58 - Spiral primes
//
// Starting with 1 and spiralling anticlockwise in the following way, a square spiral with side length 7 is formed.
//
// 37 36 35 34 33 32 31
// 38 17 16 15 14 13 30
// 39 18  5  4  3 12 29
// 40 19  6  1  2 11 28
// 41 20  7  8  9 10 27
// 42 21 22 23 24 25 26
// 43 44 45 46 47 48 49
//
// It is interesting to note that the odd squares lie along the bottom right diagonal, but what is more interesting is that 8 out of the 13 numbers lying along both diagonals are prime; that is, a ratio of 8/13 ≈ 62%.
//
// If one complete new layer is wrapped around the spiral above, a square spiral with side length 9 will be formed. If this process is continued, what is the side length of the square spiral for which the ratio of primes along both diagonals first falls below 10%?
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

// Problem 59 - XOR decryption
//
// Each character on a computer is assigned a unique code and the preferred standard is ASCII (American Standard Code for Information Interchange). For example, uppercase A = 65, asterisk (*) = 42, and lowercase k = 107.
//
// A modern encryption method is to take a text file, convert the bytes to ASCII, then XOR each byte with a given value, taken from a secret key. The advantage with the XOR function is that using the same encryption key on the cipher text, restores the plain text; for example, 65 XOR 42 = 107, then 107 XOR 42 = 65.
//
// For unbreakable encryption, the key is the same length as the plain text message, and the key is made up of random bytes. The user would keep the encrypted message and the encryption key in different locations, and without both "halves", it is impossible to decrypt the message.
//
// Unfortunately, this method is impractical for most users, so the modified method is to use a password as a key. If the password is shorter than the message, which is likely, the key is repeated cyclically throughout the message. The balance for this method is using a sufficiently long password key for security, but short enough to be memorable.
//
// Your task has been made easy, as the encryption key consists of three lower case characters. Using cipher.txt (right click and 'Save Link/Target As...'), a file containing the encrypted ASCII codes, and the knowledge that the plain text must contain common English words, decrypt the message and find the sum of the ASCII values in the original text.
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

// Problem 60 - Prime pair sets
//
// The primes 3, 7, 109, and 673, are quite remarkable. By taking any two primes and concatenating them in any order the result will always be prime. For example, taking 7 and 109, both 7109 and 1097 are prime. The sum of these four primes, 792, represents the lowest sum for a set of four primes with this property.
//
// Find the lowest sum for a set of five primes for which any two primes concatenate to produce another prime.
// Note:
// * This program can NOT ENSURE result is right.
// * This program will never stop. result 26033 (first length 5) appeared at about 60 seconds.
//   5 8389 5 [13 5197 5701 6733 8389] 26033 26033
func PE60() (ret int) {
	var data [][]string
	max := 0
	GenPrimes(1e5)
	for _, n := range primes {
		for _, set := range data {
			if canJoinSet(set, n) {
				// Append new set with n to data
				newone := make([]string, len(set)+1)
				copy(newone, set)
				newone[len(newone)-1] = strconv.Itoa(n)
				data = append(data, newone)

				// Check if result is better
				sumnewone := sumIntStrings(newone)
				lennewone := len(newone)
				if lennewone == max && sumnewone < ret || lennewone > max {
					ret = sumnewone
					fmt.Println("Length:", lennewone, "Sum:", sumnewone, "Set:", newone)
				}
				if lennewone > max {
					max = lennewone
				}
			}
		}
		data = append(data, []string{strconv.Itoa(n)})
	}
	return
}

func isPrime(n int) bool {
	sqrtn := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtn; i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func sumIntStrings(ss []string) (ret int) {
	for _, v := range ss {
		a, _ := strconv.Atoi(v)
		ret += a
	}
	return
}

func canJoinSet(set []string, value int) bool {
	sv := strconv.Itoa(value)
	for _, v := range set {
		a, _ := strconv.Atoi(v + sv)
		b, _ := strconv.Atoi(sv + v)
		if !isPrime(a) || !isPrime(b) {
			return false
		}
	}
	return true
}

// Problem 61 - Cyclical figurate numbers
//
// Triangle, square, pentagonal, hexagonal, heptagonal, and octagonal numbers are all figurate (polygonal) numbers and are generated by the following formulae:
// Triangle 	  	P3,n=n(n+1)/2 	  	1, 3, 6, 10, 15, ...
// Square		    P4,n=n2 		  	1, 4, 9, 16, 25, ...
// Pentagonal 	  	P5,n=n(3n−1)/2 	  	1, 5, 12, 22, 35, ...
// Hexagonal 	  	P6,n=n(2n−1) 	  	1, 6, 15, 28, 45, ...
// Heptagonal 	  	P7,n=n(5n−3)/2 	  	1, 7, 18, 34, 55, ...
// Octagonal 	  	P8,n=n(3n−2) 	  	1, 8, 21, 40, 65, ...
//
// The ordered set of three 4-digit numbers: 8128, 2882, 8281, has three interesting properties.
//
//     The set is cyclic, in that the last two digits of each number is the first two digits of the next number (including the last number with the first).
//     Each polygonal type: triangle (P3,127=8128), square (P4,91=8281), and pentagonal (P5,44=2882), is represented by a different number in the set.
//     This is the only set of 4-digit numbers with this property.
//
// Find the sum of the only ordered set of six cyclic 4-digit numbers for which each polygonal type: triangle, square, pentagonal, hexagonal, heptagonal, and octagonal, is represented by a different number in the set.

var permPE61 []string

func PE61() (ret int) {
	// Generate layer based data
	polygonali := GenPolygonals4i()

	// RoundPermStr("012345") == "0" + PermStr("12345",5)
	PermStrCallback = callbackPE61
	PermStr("12345", 5, "")

	var polygonalio [6][]int
	polygonalio[0] = polygonali[0]
	for _, perm := range permPE61 {
		for i, c := range perm {
			polygonalio[i+1] = polygonali[c-0x30]
		}
		// This will cause stack overflow, WHY?
		TracePath(polygonalio, []int{})
	}
	return
}

func callbackPE61(ret string) {
	permPE61 = append(permPE61, ret)
}

func TracePath(polygonali [6][]int, prefix []int) {
	round := len(prefix)
	if round == 0 {
		for _, v := range polygonali[round] {
			TracePath(polygonali, []int{v})
		}
		return
	}
	if round == 7 {
		if prefix[0] == prefix[6] {
			// Found!!!
			fmt.Println(prefix[:6], SumInts(prefix[:6]))
		}
		return
	}
	for _, v := range polygonali[round%6] {
		if connected(prefix[round-1], v) {
			prefix = append(prefix, v)
			TracePath(polygonali, prefix)
		}
	}
	return
}

func connected(a, b int) bool {
	if a%100 == b/100 {
		return true
	}
	return false
}

func GenPolygonals4i() (polygonals [6][]int) {
	for i := 1; ; i++ {
		s := Polygonals(i)
		if s[0] > 1e4 {
			return
		}
		for j := 0; j < 6; j++ {
			if s[j] >= 1e3 && s[j] < 1e4 {
				InsertUniq(&polygonals[j], s[j])
			}
		}
	}
	return
}

//////
func PE61_missunderstanding() (ret int) {
	polygonals := GenPolygonals4()

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

func Polygonals(n int) (ret [6]int) {
	ret[0] = n * (n + 1) / 2   // Triangle   P(3,n)
	ret[1] = n * n             // Square     P(4,n)
	ret[2] = n * (3*n - 1) / 2 // Pentagonal P(5,n)
	ret[3] = n * (2*n - 1)     // Hexagonal  P(6,n)
	ret[4] = n * (5*n - 3) / 2 // Heptagonal P(7,n)
	ret[5] = n * (3*n - 2)     // Octagonal  P(8,n)
	return
}

func GenPolygonals4() (polygonals []int) {
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
	return
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

// Problem 62 - Cubic permutations
//
// The cube, 41063625 (3453), can be permuted to produce two other cubes: 56623104 (3843) and 66430125 (4053). In fact, 41063625 is the smallest cube which has exactly three permutations of its digits which are also cube.
//
// Find the smallest cube for which exactly five permutations of its digits are cube.
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

// Problem 63 - Powerful digit counts
//
// The 5-digit number, 16807=7^5, is also a fifth power. Similarly, the 9-digit number, 134217728=8^9, is a ninth power.
//
// How many n-digit positive integers exist which are also an nth power?
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

// Problem 64 - Odd period square roots
//
// All square roots are periodic when written as continued fractions and can be written in the form:
// √N = [a0; (a1,a2,a3, ...)]
//
// For example, let us consider √23.
// It can be seen that the sequence is repeating. For conciseness, we use the notation √23 = [4;(1,3,1,8)], to indicate that the block (1,3,1,8) repeats indefinitely.
//
// The first ten continued fraction representations of (irrational) square roots are:
//
// √2=[1;(2)], period=1
// √3=[1;(1,2)], period=2
// √5=[2;(4)], period=1
// √6=[2;(2,4)], period=2
// √7=[2;(1,1,1,4)], period=4
// √8=[2;(1,4)], period=2
// √10=[3;(6)], period=1
// √11=[3;(3,6)], period=2
// √12= [3;(2,6)], period=2
// √13=[3;(1,1,1,1,6)], period=5
//
// Exactly four continued fractions, for N ≤ 13, have an odd period.
//
// How many continued fractions for N ≤ 10000 have an odd period?
func PE64() (ret int) {
	return
}

// Problem 65 - Convergents of e
//
func PE65() (ret int) {
	return
}

// Problem 66 - Diophantine equation
//
func PE66() (ret int) {
	return
}

// Problem 67 - Maximum path sum II
//
// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
//
// 3
// 7 4
// 2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.
//
// NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
func PE67(filename string) int {
	return MaxPathSum(filename, 100)
}

// Problem 68 - Magic 5-gon ring
//
func PE68() (ret int) {
	return
}

// Problem 69 - Totient maximum
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number of numbers less than n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.
// n 	Relatively Prime 	φ(n) 	n/φ(n)
// 2 	1 	1 	2
// 3 	1,2 	2 	1.5
// 4 	1,3 	2 	2
// 5 	1,2,3,4 	4 	1.25
// 6 	1,5 	2 	3
// 7 	1,2,3,4,5,6 	6 	1.1666...
// 8 	1,3,5,7 	4 	2
// 9 	1,2,4,5,7,8 	6 	1.5
// 10 	1,3,7,9 	4 	2.5
//
// It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.
//
// Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
func PE69() (ret int) {
	return
}

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

// Problem 70 - Totient permutation
//
func PE70() (ret int) {
	return
}

// Problem 71 - Ordered fractions
//
func PE71() (ret int) {
	return
}

// Problem 72 - Counting fractions
//
func PE72() (ret int) {
	return
}

// Problem 73 - Counting fractions in a range
//
func PE73() (ret int) {
	return
}

// Problem 74 - Digit factorial chains
//
func PE74() (ret int) {
	return
}

// Problem 75 - Singular integer right triangles
//
func PE75() (ret int) {
	return
}

// Problem 76 - Counting summations
//
func PE76() (ret int) {
	return
}

// Problem 77 - Prime summations
//
func PE77() (ret int) {
	return
}

// Problem 78 - Coin partitions
//
func PE78() (ret int) {
	return
}

// Problem 79 - Passcode derivation
//
func PE79() (ret int) {
	return
}

// Problem 80 - Square root digital expansion
//
func PE80() (ret int) {
	return
}

// Problem 81 - Path sum: two ways
//
func PE81() (ret int) {
	return
}

// Problem 82 - Path sum: three ways
//
func PE82() (ret int) {
	return
}

// Problem 83 - Path sum: four ways
//
func PE83() (ret int) {
	return
}

// Problem 84 - Monopoly odds
//
func PE84() (ret int) {
	return
}

// Problem 85 - Counting rectangles
//
func PE85() (ret int) {
	return
}

// Problem 86 - Cuboid route
//
func PE86() (ret int) {
	return
}

// Problem 87 - Prime power triples
//
func PE87() (ret int) {
	return
}

// Problem 88 - Product-sum numbers
//
func PE88() (ret int) {
	return
}

// Problem 89 - Roman numerals
//
func PE89() (ret int) {
	return
}

// Problem 90 - Cube digit pairs
//
func PE90() (ret int) {
	return
}

// Problem 91 - Right triangles with integer coordinates
//
func PE91() (ret int) {
	return
}

// Problem 92 - Square digit chains
//
func PE92() (ret int) {
	return
}

// Problem 93 - Arithmetic expressions
//
func PE93() (ret int) {
	return
}

// Problem 94 - Almost equilateral triangles
//
func PE94() (ret int) {
	return
}

// Problem 95 - Amicable chains
//
func PE95() (ret int) {
	return
}

// Problem 96 - Su Doku
//
func PE96() (ret int) {
	return
}

// Problem 97 - Large non-Mersenne prime
//
// The first known prime found to exceed one million digits was discovered in 1999, and is a Mersenne prime of the form 26972593−1; it contains exactly 2,098,960 digits. Subsequently other Mersenne primes, of the form 2p−1, have been found which contain more digits.
//
// However, in 2004 there was found a massive non-Mersenne prime which contains 2,357,207 digits: 28433×27830457+1.
//
// Find the last ten digits of this prime number.
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

// Problem 98 - Anagramic squares
//
func PE98() (ret int) {
	return
}

// Problem 99 - Largest exponential
//
// Comparing two numbers written in index form like 2^11 and 3^7 is not difficult, as any calculator would confirm that 2^11 = 2048 < 3^7 = 2187.
//
// However, confirming that 632382^518061 > 519432^525806 would be much more difficult, as both numbers contain over three million digits.
//
// Using base_exp.txt (right click and 'Save Link/Target As...'), a 22K text file containing one thousand lines with a base/exponent pair on each line, determine which line number has the greatest numerical value.
//
// NOTE: The first two lines in the file represent the numbers in the example given above.
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

// Problem 100 - Arranged probability
//
func PE100() (ret int) {
	return
}
