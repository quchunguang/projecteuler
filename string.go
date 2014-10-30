package projecteuler

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//////
func Palindrome6(N int64) bool {
	var digit [6]int64
	var i int64
	for i = 0; i < 6; i++ {
		digit[i] = N % 10
		N /= 10
	}
	if digit[5] != digit[0] || digit[4] != digit[1] || digit[3] != digit[2] {
		return false
	}
	return true
}

//////
func CancelStr(sa, sb string) (rsa, rsb string, err bool) {
	if sa[0] == sb[0] {
		rsa = string(sa[1])
		rsb = string(sb[1])
	} else if sa[0] == sb[1] {
		rsa = string(sa[1])
		rsb = string(sb[0])
	} else if sa[1] == sb[0] {
		rsa = string(sa[0])
		rsb = string(sb[1])
	} else if sa[1] == sb[1] {
		if string(sa[1]) == "0" {
			err = true
			return
		}
		rsa = string(sa[0])
		rsb = string(sb[0])
	} else {
		err = true
		return
	}
	if rsa == "0" || rsb == "0" {
		err = true
		return
	}
	err = false
	return
}

//////
func IsPalindromeBytes(s []byte) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}

//////
func StrsEquals(a, b []string) bool {
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

//////
func SumAscii(word string) (sum int) {
	for _, v := range word {
		sum += int(v - 1<<6)
	}
	return sum
}

//////
func ReverseBytes(s []byte) []byte {
	ret := make([]byte, len(s))
	copy(ret, s)
	last := len(ret) - 1
	for i := 0; i < last-i; i++ {
		ret[i], ret[last-i] = ret[last-i], ret[i]
	}
	return ret
}
func PlusBytes(a, b []byte) []byte {
	var ret, plus []byte
	var carry byte = 0
	if len(a) < len(b) {
		ret = make([]byte, len(b))
		copy(ret, b)
		plus = a
	} else {
		ret = make([]byte, len(a))
		copy(ret, a)
		plus = b
	}
	for i, j := len(ret)-1, len(plus)-1; i >= 0; i, j = i-1, j-1 {
		if j >= 0 {
			ret[i] += plus[j] - 0x30 + carry
		} else {
			ret[i] += carry
		}
		if ret[i] > 0x39 {
			ret[i] -= 10
			carry = 1
		} else {
			carry = 0
		}
	}
	if carry == 1 {
		ret = append([]byte("1"), ret...)
	}
	return ret
}

//////
func IsLychrel(N int) bool {
	num := []byte(strconv.Itoa(N))
	for i := 1; i < 50; i++ {
		num = PlusBytes(num, ReverseBytes(num))
		if IsPalindromeBytes(num) {
			return false
		}
	}
	return true
}

//////
func SortString(w string) string {
	s := strings.Split(w, "")
	sort.Strings(s)
	return strings.Join(s, "")
}

//////
func ScoreWord(word string) (ret int) {
	for _, c := range word {
		ret += int(c - 0x40)
	}
	return
}

// Convert digits string to []int
// ex. "135" -> [1, 3, 5]
func DigitsInts(digits string) (ret []int) {
	for _, c := range digits {
		ret = append(ret, int(c-0x30))
	}
	return
}

func IntsDigits(ints []int) (ret string) {
	for _, v := range ints {
		ret += strconv.Itoa(v)
	}
	return
}

// Convert digits string (as locations start with 1) to []bool
// ex. "135", 6 -> {true, false, true, false, true, false}
func TrueDigits(digits string, length int) []bool {
	zeroones := make([]bool, length)
	for i := 0; i < len(digits); i++ {
		zeroones[digits[i]-0x30-1] = true
	}
	return zeroones
}

// "123456" , {true, false, true, false, true, false} -> "135", "246"
func SplitDigits(digits string, zeroone []bool) (truestr, falsestr string) {
	if len(digits) != len(zeroone) {
		fmt.Println(digits, zeroone)
		fmt.Println("SplitDigits() needs same length!")
	}
	for i := 0; i < len(digits); i++ {
		if zeroone[i] {
			truestr += string(digits[i])
		} else {
			falsestr += string(digits[i])
		}
	}
	return
}

// "222222" -> true; "222322" -> false
func IsSameStr(s string) bool {
	head := s[0]
	for i := 1; i < len(s); i++ {
		if head != s[i] {
			return false
		}
	}
	return true
}
