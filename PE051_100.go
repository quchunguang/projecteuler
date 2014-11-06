package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"math/big"
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
		if !IsPrime(a) || !IsPrime(b) {
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
func PE61() (ret int) {
	// Generate layer based data with origin sort
	GenPolygonals4i()

	PermStrCallback = callbackPE61
	RoundPermStr("012345")

	return
}

var polygonali [6][]int  // origin sort of data
var polygonalio [6][]int // working sort of data

func callbackPE61(ret string) {
	// Use ret set as index to sort data to polygonalio
	for i, c := range ret {
		polygonalio[i] = polygonali[c-0x30]
	}
	TracePath([]int{})
}

func TracePath(prefix []int) {
	round := len(prefix)
	if round == 0 {
		for _, v := range polygonalio[round] {
			TracePath([]int{v})
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
	for _, v := range polygonalio[round%6] {
		if connected(prefix[round-1], v) {
			TracePath(append(prefix, v))
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

func GenPolygonals4i() {
	for i := 1; ; i++ {
		s := Polygonals(i)
		if s[0] > 1e4 {
			return
		}
		for j := 0; j < 6; j++ {
			if s[j] >= 1e3 && s[j] < 1e4 {
				InsertUniq(&polygonali[j], s[j])
			}
		}
	}
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
func PE64(N int64) (ret int) {
	var n int64
	for n = 2; n <= N; n++ {
		_, period := ContinueFraction(n)
		if len(period)%2 != 0 {
			ret++
		}
	}

	return
}

// √13=[3;(1,1,1,1,6)], period=5
// root=13 => first=3, period=[1,1,1,1,6]
func ContinueFraction(root int64) (first int64, period []int64) {
	var b, b0 int64 = 1, 1
	first = int64(math.Sqrt(float64(root)))
	if root == first*first { // square number
		return
	}
	var c, c0 int64 = first, first
	for {
		na, nb, nc := formatFraction(root, b, c)
		period = append(period, na)
		b = nb
		c = nc
		if b0 == b && c0 == c {
			return
		}
	}
	return
}

// b / (√r-c) => na + (√r - nc)/nb
func formatFraction(r int64, b, c int64) (na, nb, nc int64) {
	nb = (r - c*c) / b
	for na = 1; (na*nb-c)*(na*nb-c) < r; na++ {
	}
	na--
	nc = na*nb - c
	return
}

// Problem 65 - Convergents of e
//
// The square root of 2 can be written as an infinite continued fraction.
// The infinite continued fraction can be written, √2 = [1;(2)], (2) indicates that 2 repeats ad infinitum. In a similar way, √23 = [4;(1,3,1,8)].
// It turns out that the sequence of partial values of continued fractions for square roots provide the best rational approximations. Let us consider the convergents for √2.
// Hence the sequence of the first ten convergents for √2 are:
//   1, 3/2, 7/5, 17/12, 41/29, 99/70, 239/169, 577/408, 1393/985, 3363/2378, ...
//
// What is most surprising is that the important mathematical constant,
//   e = [2; 1,2,1, 1,4,1, 1,6,1 , ... , 1,2k,1, ...].
//
// The first ten terms in the sequence of convergents for e are:
//   2, 3, 8/3, 11/4, 19/7, 87/32, 106/39, 193/71, 1264/465, 1457/536, ...
//
// The sum of digits in the numerator of the 10th convergent is 1+4+5+7=17.
//
// Find the sum of digits in the numerator of the 100th convergent of the continued fraction for e.
func PE65() int {
	var list []int64
	var i int64
	for i = 1; i <= 33; i++ { // 100 == 1 + 3*33
		list = append(list, 1, 2*i, 1)
	}
	res := ContinueFractionSimplify(2, list)
	snum := res.Num().String()
	return SumInts(DigitsInts(snum))
}

// √2 = [1;(2)]
// first=1, adder=[2,2,2] => 1+1/(2+1/(2+1/2))
func ContinueFractionSimplify(first int64, adder []int64) *big.Rat {
	ret := big.NewRat(0, 1)
	for i := len(adder) - 1; i >= 0; i-- {
		ret.Add(ret, big.NewRat(adder[i], 1))
		ret.Inv(ret)
	}
	ret.Add(ret, big.NewRat(first, 1))
	return ret
}

// Problem 66 - Diophantine equation
//
// Consider quadratic Diophantine equations of the form:
//
//   x^2 - Dy^2 = 1
//
// For example, when D=13, the minimal solution in x is 6492 - 13×1802 = 1.
// It can be assumed that there are no solutions in positive integers when D is square.
// By finding minimal solutions in x for D = {2, 3, 5, 6, 7}, we obtain the following:
//
//   3^2 - 2×2^2 = 1
//   2^2 - 3×1^2 = 1
//   9^2 - 5×4^2 = 1
//   5^2 - 6×2^2 = 1
//   8^2 - 7×3^2 = 1
//
// Hence, by considering minimal solutions in x for D ≤ 7, the largest x is obtained when D=5.
// Find the value of D ≤ 1000 in minimal solutions of x for which the largest value of x is obtained.
func PE66(D int64) (ret int64) {
	var i, d int64
	xmax := big.NewInt(0)

	for i, d = 2, 2; d <= D; d++ {
		// Pell equation has no natural number solution where D>0 is a square number.
		if d == i*i {
			i++
			continue
		}

		x, y := SolveDiophantine(d)
		// x, y := SolveDiophantine1(d)
		// x,y := SolveDiophantine2(61))
		// x, y := SolveDiophantine3(int64(d))

		if x.Cmp(xmax) == 1 {
			xmax = x
			ret = d
		}
		fmt.Println(d, x, y)
	}

	return
}

// Solved by finding the continued fraction [a0; (a1,a2, ... an-1, 2*a0)] of sqrt(D).
//   p/q = [a0; a1,a2, ... an-1]
// http://blog.csdn.net/wh2124335/article/details/8871535
// http://mathworld.wolfram.com/PellEquation.html
func SolveDiophantine(D int64) (x, y *big.Int) {
	first, period := ContinueFraction(D)
	res := ContinueFractionSimplify(int64(first), period[:len(period)-1])
	p := res.Num()
	q := res.Denom()
	if len(period)%2 == 0 {
		x = p
		y = q
	} else {
		x = new(big.Int)
		y = new(big.Int)
		x.Mul(p, p).Mul(x, big.NewInt(2)).Add(x, big.NewInt(1))
		y.Mul(p, q).Mul(y, big.NewInt(2))
	}
	return
}

// overflow
func SolveDiophantine1(D int) (x, y int) {
	x, y = 1, 1
	for {
		less := x*x - D*y*y - 1
		if less < 0 {
			x++
		} else if less > 0 {
			y++
		} else {
			return
		}
	}
	return
}

// too slow
func SolveDiophantine2(D int) (x, y []int64) {
	x = BigNum("1")
	y = BigNum("1")
	one := BigNum("1")
	for {
		l := BigMul(x, x)
		r := BigSum(BigMulInt(BigMul(y, y), int64(D)), one)
		less := BigLess(l, r)
		if less == 1 {
			x = BigSum(x, one)
		} else if less == -1 {
			y = BigSum(y, one)
		} else {
			return
		}
	}
	return
}

// too slow
func SolveDiophantine3(D int64) (x, y *big.Int) {
	x = big.NewInt(1)
	y = big.NewInt(1)
	d := big.NewInt(D)
	xx := new(big.Int)
	dyy := new(big.Int)
	one := big.NewInt(1)
	for {
		xx.Mul(x, x)
		dyy.Mul(y, y)
		dyy.Mul(dyy, d)
		xx.Sub(xx, dyy)
		cmp := xx.Cmp(one)
		if cmp == -1 {
			x.Add(x, one)
		} else if cmp == 1 {
			y.Add(y, one)
		} else {
			return
		}
	}
	return
}

// Problem 67 - Maximum path sum II
//
// By starting at the top of the triangle below and moving to adjacent numbers on the row below, the maximum total from top to bottom is 23.
//
//    3
//   7 4
//  2 4 6
// 8 5 9 3
//
// That is, 3 + 7 + 4 + 9 = 23.
//
// Find the maximum total from top to bottom in triangle.txt (right click and 'Save Link/Target As...'), a 15K text file containing a triangle with one-hundred rows.
//
// NOTE: This is a much more difficult version of Problem 18. It is not possible to try every route to solve this problem, as there are 299 altogether! If you could check one trillion (1012) routes every second it would take over twenty billion years to check them all. There is an efficient algorithm to solve it. ;o)
func PE67(filename string) int {
	data := SST(filename)
	// Find biggest path
	return findPathMax2(data)
}

type datai struct {
	sm    int  // Biggest sum
	right bool // Previous item go to here from right path?(or left)
}

func findPathMax2(data [][]int) (ret int) {
	// Create temp data structure
	N := len(data)
	sd := make([][]datai, N)
	for i := 0; i < N; i++ {
		sd[i] = make([]datai, i+1)
	}

	// Calculate sd[i][j], the biggest sum data[0][0] -> data[i][j]
	sd[0][0].sm = data[0][0]
	for i := 1; i < N; i++ {
		//j==0
		sd[i][0].sm = sd[i-1][0].sm + data[i][0]
		sd[i][0].right = true
		//j==1..i-1
		for j := 1; j < i; j++ {
			if sd[i-1][j-1].sm > sd[i-1][j].sm {
				sd[i][j].sm = sd[i-1][j-1].sm + data[i][j]
				sd[i][j].right = false
			} else {
				sd[i][j].sm = sd[i-1][j].sm + data[i][j]
				sd[i][j].right = true
			}
		}
		//j==i
		sd[i][i].sm = sd[i-1][i-1].sm + data[i][i]
		sd[i][i].right = false
	}

	// Get result
	rets := make([]int, N)
	for j := 0; j < N; j++ {
		if sd[N-1][j].sm > ret {
			ret = sd[N-1][j].sm
			rets[N-1] = j
		}
	}
	for i := N - 2; i >= 0; i-- {
		if sd[i+1][rets[i+1]].right {
			rets[i] = rets[i+1]
		} else {
			rets[i] = rets[i+1] - 1
		}
	}

	// Print result
	for i := 0; i < N; i++ {
		fmt.Printf("%d ", data[i][rets[i]])
	}
	fmt.Println()
	return
}

// Problem 68 - Magic 5-gon ring
//
// Consider the following "magic" 3-gon ring, filled with the numbers 1 to 6, and each line adding to nine.
// Working clockwise, and starting from the group of three with the numerically lowest external node (4,3,2 in this example), each solution can be described uniquely. For example, the above solution can be described by the set: 4,3,2; 6,2,1; 5,1,3.
//        4
//           3
//         1   2   6
//      5
// It is possible to complete the ring with four different totals: 9, 10, 11, and 12. There are eight solutions in total.
// Total	Solution Set
// 9		4,2,3; 5,3,1; 6,1,2
// 9		4,3,2; 6,2,1; 5,1,3
// 10		2,3,5; 4,5,1; 6,1,3
// 10		2,5,3; 6,3,1; 4,1,5
// 11		1,4,6; 3,6,2; 5,2,4
// 11		1,6,4; 5,4,2; 3,2,6
// 12		1,5,6; 2,6,4; 3,4,5
// 12		1,6,5; 3,5,4; 2,4,6
//
// By concatenating each group it is possible to form 9-digit strings; the maximum string for a 3-gon ring is 432621513.
// Using the numbers 1 to 10, and depending on arrangements, it is possible to form 16- and 17-digit strings. What is the maximum 16-digit string for a "magic" 5-gon ring?
//       0
//           5     1
//        9      6
//     4    8  7  2
//            3
func PE68() (ret int) {
	PermStrCallback = callbackPE68
	PermStr("0123456789", 10, "")
	return
}

func callbackPE68(ret string) {
	// Convert string to ints (0..9)
	ints := DigitsInts(ret)

	// first external node (ints[0]) must be the smallest of all external nodes (ints[0:5])
	min := ints[0]
	for i := 1; i < 5; i++ {
		if ints[i] < min {
			return
		}
	}

	// Convert 0..9 -> 1..10
	for i := 0; i < len(ints); i++ {
		ints[i]++
	}

	s167 := ints[0] + ints[5] + ints[6]
	var s278 int
	for i := 1; i < 5; i++ {
		s278 = ints[i] + ints[i+5] + ints[(i+1)%5+5]
		if s167 != s278 {
			return
		}
	}

	// Found!!!
	// Build string
	s := ""
	for i := 0; i < 5; i++ {
		s += strconv.Itoa(ints[i])
		s += strconv.Itoa(ints[i+5])
		s += strconv.Itoa(ints[(i+1)%5+5])
	}
	// Remove 17-digit strings
	if len(s) == 17 {
		return
	}

	// Output result, the last one printed is the answer (6531031914842725)
	fmt.Println(s167, ints, s)
}

// Problem 69 - Totient maximum
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number of numbers less than n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.
//
// n 	Relatively Prime 	φ(n) 	n/φ(n)
// 2 	1 					1 		2
// 3 	1,2 				2 		1.5
// 4 	1,3 				2 		2
// 5 	1,2,3,4 			4 		1.25
// 6 	1,5 				2 		3
// 7 	1,2,3,4,5,6 		6 		1.1666...
// 8 	1,3,5,7 			4 		2
// 9 	1,2,4,5,7,8 		6 		1.5
// 10 	1,3,7,9 			4 		2.5
//
// It can be seen that n=6 produces a maximum n/φ(n) for n ≤ 10.
// Find the value of n ≤ 1,000,000 for which n/φ(n) is a maximum.
// It will run about 70s.
func PE69(N int64) (ret int64) {
	var ratephi, max float64
	var n int64
	for n = 2; n < N; n++ {
		ratephi = float64(n) / float64(EulerPhi(n))
		if ratephi > max {
			max = ratephi
			ret = n
		}
	}
	return
}

// Problem 70 - Totient permutation
//
// Euler's Totient function, φ(n) [sometimes called the phi function], is used to determine the number of positive numbers less than or equal to n which are relatively prime to n. For example, as 1, 2, 4, 5, 7, and 8, are all less than nine and relatively prime to nine, φ(9)=6.
// The number 1 is considered to be relatively prime to every positive number, so φ(1)=1.
// Interestingly, φ(87109)=79180, and it can be seen that 87109 is a permutation of 79180.
// Find the value of n, 1 < n < 107, for which φ(n) is a permutation of n and the ratio n/φ(n) produces a minimum.
// This program will cost 79m33.594s, and result will get at 83%.
func PE70(N int64) (ret int64) {
	var ratephi, min float64 = 0, 1e10
	var n int64
	for n = 2; n < N; n++ {
		// Process bar
		if n%100000 == 0 {
			fmt.Print(EL, n/100000, "%")
		}
		phi := EulerPhi(n)
		if IsPermutations(int(n), int(phi)) {
			ratephi = float64(n) / float64(phi)
			if ratephi < min {
				min = ratephi
				ret = n
				fmt.Println(EL, n, phi, ratephi)
			}
		}
	}
	fmt.Print(EL)
	return
}

// Problem 71 - Ordered fractions
//
// Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:
//
//   1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that 2/5 is the fraction immediately to the left of 3/7.
// By listing the set of reduced proper fractions for d ≤ 1,000,000 in ascending order of size, find the numerator of the fraction immediately to the left of 3/7.
func PE71(D int64) (ret int64) {
	var max, nd float64
	s37 := big.NewRat(3, 7)
	r := float64(3) / float64(7)
	for d := D; d > 1; d-- {
		n := int(float64(d) * r)
		s := big.NewRat(int64(n), int64(d))
		if s.Cmp(s37) == 0 {
			n--
			s = big.NewRat(int64(n), int64(d))
		}
		nd = float64(n) / float64(d)
		if nd > max {
			max = nd
			ret = s.Num().Int64()
			// fmt.Println(s)
		}
	}
	return
}

// Problem 72 - Counting fractions
//
// Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 21 elements in this set.
// How many elements would be contained in the set of reduced proper fractions for d ≤ 1,000,000?
func PE72(D int64) (ret int64) {
	var d int64
	for d = 2; d <= D; d++ {
		// Process bar
		if d%10000 == 0 {
			fmt.Print(EL, d/10000, "%")
		}

		ret += EulerPhi(d)
	}
	fmt.Print(EL)
	return
}

// Problem 73 - Counting fractions in a range
//
// Consider the fraction, n/d, where n and d are positive integers. If n<d and HCF(n,d)=1, it is called a reduced proper fraction.
// If we list the set of reduced proper fractions for d ≤ 8 in ascending order of size, we get:
//
// 1/8, 1/7, 1/6, 1/5, 1/4, 2/7, 1/3, 3/8, 2/5, 3/7, 1/2, 4/7, 3/5, 5/8, 2/3, 5/7, 3/4, 4/5, 5/6, 6/7, 7/8
//
// It can be seen that there are 3 fractions between 1/3 and 1/2.
// How many fractions lie between 1/3 and 1/2 in the sorted set of reduced proper fractions for d ≤ 12,000?
// slow: about 5 minutes
func PE73(D int64) (ret int64) {
	v13 := big.NewRat(1, 3)
	v12 := big.NewRat(1, 2)
	var d, n int64
	for d = 2; d <= D; d++ {
		if d%1000 == 0 {
			fmt.Print(EL, "d = ", d)
		}
		for n = 1; n < d; n++ {
			v := big.NewRat(n, d)
			if v13.Cmp(v) == -1 && v12.Cmp(v) == 1 && v.Num().Int64() == n {
				ret++
			}
		}
	}
	fmt.Print(EL)
	return
}

// Problem 74 - Digit factorial chains
//
// The number 145 is well known for the property that the sum of the factorial of its digits is equal to 145:
//
// 1! + 4! + 5! = 1 + 24 + 120 = 145
//
// Perhaps less well known is 169, in that it produces the longest chain of numbers that link back to 169; it turns out that there are only three such loops that exist:
//
// 169 → 363601 → 1454 → 169
// 871 → 45361 → 871
// 872 → 45362 → 872
//
// It is not difficult to prove that EVERY starting number will eventually get stuck in a loop. For example,
//
// 69 → 363600 → 1454 → 169 → 363601 (→ 1454)
// 78 → 45360 → 871 → 45361 (→ 871)
// 540 → 145 (→ 145)
//
// Starting with 69 produces a chain of five non-repeating terms, but the longest non-repeating chain with a starting number below one million is sixty terms.
// How many chains, with a starting number below one million, contain exactly sixty non-repeating terms?
func PE74(N int) (ret int) {
	Genfacts()
	var ichain []int
	for n := 0; n < N; n++ {
		item := n
		ichain = append(ichain, item)
		for {
			item = SumDigFact(item)
			if InIntsUnsort(ichain, item) {
				if len(ichain) == 60 {
					ret++
				}
				break
			} else {
				ichain = append(ichain, item)
			}
		}
		ichain = nil
	}
	return
}

var facts [10]int

func Genfacts() {
	fact := 1
	facts[0] = 1
	for i := 1; i < 10; i++ {
		fact *= i
		facts[i] = fact
	}
}

// Sum of factor of each of digits of given int
func SumDigFact(N int) (ret int) {
	digits := DigNums(N)
	for _, v := range digits {
		ret += facts[v]
	}
	return
}

// Problem 75 - Singular integer right triangles
//
// It turns out that 12 cm is the smallest length of wire that can be bent to form an integer sided right angle triangle in exactly one way, but there are many more examples.
//
// 12 cm: (3,4,5)
// 24 cm: (6,8,10)
// 30 cm: (5,12,13)
// 36 cm: (9,12,15)
// 40 cm: (8,15,17)
// 48 cm: (12,16,20)
//
// In contrast, some lengths of wire, like 20 cm, cannot be bent to form an integer sided right angle triangle, and other lengths allow more than one solution to be found; for example, using 120 cm it is possible to form exactly three different integer sided right angle triangles.
//
// 120 cm: (30,40,50), (20,48,52), (24,45,51)
//
// Given that L is the length of the wire, for how many values of L ≤ 1,500,000 can exactly one integer sided right angle triangle be formed?
func PE75(L int) (ret int) {
	var triangles = make([]int, L+1)
	var mL = int(math.Sqrt(float64(L) / 2))

	for m := 2; m < mL; m++ {
		for n := 1; n < m; n++ {
			if (n+m)%2 == 1 && Gcd(n, m) == 1 {
				a := m*m - n*n
				b := 2 * m * n
				c := m*m + n*n
				p := a + b + c
				for p <= L {
					triangles[p]++
					if triangles[p] == 1 {
						ret++
					} else if triangles[p] == 2 {
						ret--
					}
					p += a + b + c
				}
			}
		}
	}
	return
}

// Problem 76 - Counting summations
//
// It is possible to write five as a sum in exactly six different ways:
//   4 + 1
//   3 + 2
//   3 + 1 + 1
//   2 + 2 + 1
//   2 + 1 + 1 + 1
//   1 + 1 + 1 + 1 + 1
// How many different ways can one hundred be written as a sum of at least two positive integers?
func PE76(N int) (ret int) {
	// Create 2D slice, P[m][n] are methods of sum of numbers with 1..m-1.
	P := make([][]int, N+1, N+1)
	for n := 1; n <= N; n++ {
		P[n] = make([]int, N+1, N+1)
	}

	// P[n][0] = 1
	for n := 2; n <= N; n++ {
		P[n][0] = 1
	}
	// P[2][i] = 1
	for i := 1; i <= N; i++ {
		P[2][i] = 1
	}

	// P[n][i]
	for n := 3; n <= N; n++ {
		for i := 1; i <= N; i++ {
			for k := i; k >= 0; k -= n - 1 {
				P[n][i] += P[n-1][k]
			}
		}
	}
	// for i := 1; i <= N; i++ {
	// 	fmt.Println(P[i][i])
	// }
	ret = P[N][N]
	return
}

// Too slow to finish
func PE76_slow(N int) (ret int) {
	planNum(N-1, N, nil)
	return
}

func planNum(max int, remain int, preset []int) {
	// fmt.Printf("planNum(%d,%d,%v)\n", max, remain, preset)
	if remain == 0 && max == 0 {
		fmt.Println(preset)
	}
	if max == 0 { // remain!=0
		return
	}
	// not use max
	planNum(max-1, remain, preset)
	// use max
	for remain >= max {
		preset = append(preset, max)
		remain -= max
		planNum(max-1, remain, preset)
	}
}

// Problem 77 - Prime summations
//
// It is possible to write ten as the sum of primes in exactly five different ways:
//   7 + 3
//   5 + 5
//   5 + 3 + 2
//   3 + 3 + 2 + 2
//   2 + 2 + 2 + 2 + 2
// What is the first value which can be written as the sum of primes in over five thousand different ways?
func PE77(W int) (ret int) {
	for n := 4; ; n++ {
		ways = 0
		GenPrimes(n)
		index := sort.SearchInts(primes, n)
		planPrime(index-1, n, nil)
		if ways >= W {
			ret = n
			return
		}
	}
}

var ways int

func planPrime(maxindex int, remain int, preset []int) {
	if remain == 0 && maxindex == -1 {
		// fmt.Println(SumInts(preset), "=", JoinInts(preset, " + "))
		ways++
	}
	if maxindex == -1 { // remain != 0
		return
	}
	for remain >= 0 {
		planPrime(maxindex-1, remain, preset)
		preset = append(preset, primes[maxindex])
		remain -= primes[maxindex]
	}
}

// Problem 78 - Coin partitions
//
// Let p(n) represent the number of different ways in which n coins can be separated into piles. For example, five coins can separated into piles in exactly seven different ways, so p(5)=7.
//   OOOOO
//   OOOO   O
//   OOO   OO
//   OOO   O   O
//   OO   OO   O
//   OO   O   O   O
//   O   O   O   O   O
// Find the least value of n for which p(n) is divisible by one million.
// Cost about 7 minutes
func PE78(L int) (ret int) {
	const N int = 100000 // Add if find no answer

	P0 := make([]int, N+1, N+1)
	P2 := make([]int, N+1, N+1)
	Pn := make([]int, N+1, N+1)

	// P2[i] = 1
	for i := 0; i <= N; i++ {
		P2[i] = 1
	}

	// Pn[i]
	for n := 3; n <= N; n++ {
		if n%100 == 0 {
			fmt.Println(GreenStr("n = " + strconv.Itoa(n)))
		}
		Pn[0] = 1
		for i := 1; i <= N; i++ {
			for k := i; k >= 0; k -= n - 1 {
				Pn[i] += P2[k]
				if Pn[i] > 1e8 {
					Pn[i] %= 1e8
				}
			}
		}
		if (Pn[n]+1)%L == 0 {
			ret = n
			return
		}
		copy(P2, Pn)
		copy(Pn, P0)
	}
	fmt.Println(RedStr("Find no answer, 10 times N and retry"))
	return
}

// Problem 79 - Passcode derivation
//
// A common security method used for online banking is to ask the user for three random characters from a passcode. For example, if the passcode was 531278, they may ask for the 2nd, 3rd, and 5th characters; the expected reply would be: 317.
// The text file, keylog.txt, contains fifty successful login attempts.
// Given that the three characters are always asked for in order, analyse the file so as to determine the shortest possible secret passcode of unknown length.
func PE79() (ret int) {
	return
}

// Problem 80 - Square root digital expansion
//
// It is well known that if the square root of a natural number is not an integer, then it is irrational. The decimal expansion of such square roots is infinite without any repeating pattern at all.
// The square root of two is 1.41421356237309504880..., and the digital sum of the first one hundred decimal digits is 475.
// For the first one hundred natural numbers, find the total of the digital sums of the first one hundred decimal digits for all the irrational square roots.
func PE80() (ret int) {
	return
}

// Problem 81 - Path sum: two ways
//
// In the 5 by 5 matrix below, the minimal path sum from the top left to the bottom right, by only moving to the right and down, is indicated in bold red and is equal to 2427.
//   \begin{pmatrix} \color{red}{131} & 673 & 234 & 103 & 18\\ \color{red}{201} & \color{red}{96} & \color{red}{342} & 965 & 150\\ 630 & 803 & \color{red}{746} & \color{red}{422} & 111\\ 537 & 699 & 497 & \color{red}{121} & 956\\ 805 & 732 & 524 & \color{red}{37} & \color{red}{331} \end{pmatrix}
// Find the minimal path sum, in matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing a 80 by 80 matrix, from the top left to the bottom right by only moving right and down.
func PE81() (ret int) {
	return
}

// Problem 82 - Path sum: three ways
//
//The minimal path sum in the 5 by 5 matrix below, by starting in any cell in the left column and finishing in any cell in the right column, and only moving up, down, and right, is indicated in red and bold; the sum is equal to 994.
//   \begin{pmatrix} \color{red}{131} & 673 & 234 & 103 & 18\\ \color{red}{201} & \color{red}{96} & \color{red}{342} & 965 & 150\\ 630 & 803 & \color{red}{746} & \color{red}{422} & 111\\ 537 & 699 & 497 & \color{red}{121} & 956\\ 805 & 732 & 524 & \color{red}{37} & \color{red}{331} \end{pmatrix}
//Find the minimal path sum, in matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing a 80 by 80 matrix, from the left column to the right column.
func PE82() (ret int) {
	return
}

// Problem 83 - Path sum: four ways
//
// In the 5 by 5 matrix below, the minimal path sum from the top left to the bottom right, by moving left, right, up, and down, is indicated in bold red and is equal to 2297.
//   \begin{pmatrix} \color{red}{131} & 673 & 234 & 103 & 18\\ \color{red}{201} & \color{red}{96} & \color{red}{342} & 965 & 150\\ 630 & 803 & \color{red}{746} & \color{red}{422} & 111\\ 537 & 699 & 497 & \color{red}{121} & 956\\ 805 & 732 & 524 & \color{red}{37} & \color{red}{331} \end{pmatrix}
// Find the minimal path sum, in matrix.txt (right click and "Save Link/Target As..."), a 31K text file containing a 80 by 80 matrix, from the top left to the bottom right by moving left, right, up, and down.
func PE83() (ret int) {
	return
}

// Problem 84 - Monopoly odds
//
// In the game, Monopoly, the standard board is set up in the following way:
//
//   GO A1 	CC1 A2 	T1 	R1 	B1 	CH1 B2 	B3  JAIL
//   H2 	  								C1
//   T2 	  								U1
//   H1 	  								C2
//   CH3 	  								C3
//   R4 	  								R2
//   G3 	  								D1
//   CC3 	  								CC2
//   G2 	  								D2
//   G1 	  								D3
//   G2J F3 U2 	F2 	F1 	R3 	E3 	E2 	CH2 E1 	FP
//
// A player starts on the GO square and adds the scores on two 6-sided dice to determine the number of squares they advance in a clockwise direction. Without any further rules we would expect to visit each square with equal probability: 2.5%. However, landing on G2J (Go To Jail), CC (community chest), and CH (chance) changes this distribution.
//
// In addition to G2J, and one card from each of CC and CH, that orders the player to go directly to jail, if a player rolls three consecutive doubles, they do not advance the result of their 3rd roll. Instead they proceed directly to jail.
//
// At the beginning of the game, the CC and CH cards are shuffled. When a player lands on CC or CH they take a card from the top of the respective pile and, after following the instructions, it is returned to the bottom of the pile. There are sixteen cards in each pile, but for the purpose of this problem we are only concerned with cards that order a movement; any instruction not concerned with movement will be ignored and the player will remain on the CC/CH square.
//
//     Community Chest (2/16 cards):
//         Advance to GO
//         Go to JAIL
//     Chance (10/16 cards):
//         Advance to GO
//         Go to JAIL
//         Go to C1
//         Go to E3
//         Go to H2
//         Go to R1
//         Go to next R (railway company)
//         Go to next R
//         Go to next U (utility company)
//         Go back 3 squares.
//
// The heart of this problem concerns the likelihood of visiting a particular square. That is, the probability of finishing at that square after a roll. For this reason it should be clear that, with the exception of G2J for which the probability of finishing on it is zero, the CH squares will have the lowest probabilities, as 5/8 request a movement to another square, and it is the final square that the player finishes at on each roll that we are interested in. We shall make no distinction between "Just Visiting" and being sent to JAIL, and we shall also ignore the rule about requiring a double to "get out of jail", assuming that they pay to get out on their next turn.
//
// By starting at GO and numbering the squares sequentially from 00 to 39 we can concatenate these two-digit numbers to produce strings that correspond with sets of squares.
//
// Statistically it can be shown that the three most popular squares, in order, are JAIL (6.24%) = Square 10, E3 (3.18%) = Square 24, and GO (3.09%) = Square 00. So these three most popular squares can be listed with the six-digit modal string: 102400.
//
// If, instead of using two 6-sided dice, two 4-sided dice are used, find the six-digit modal string.
func PE84() (ret int) {
	return
}

// Problem 85 - Counting rectangles
//
// By counting carefully it can be seen that a rectangular grid measuring 3 by 2 contains eighteen rectangles:
// Although there exists no rectangular grid that contains exactly two million rectangles, find the area of the grid with the nearest solution.
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
func PE97() (ret int64) {
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
