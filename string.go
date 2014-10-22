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
var SumSubStrDiv int = 0

func IsSubStrDiv(strnum string) bool {
	d234, _ := strconv.Atoi(strnum[1:4])
	if d234%2 != 0 {
		return false
	}
	d345, _ := strconv.Atoi(strnum[2:5])
	if d345%3 != 0 {
		return false
	}
	d456, _ := strconv.Atoi(strnum[3:6])
	if d456%5 != 0 {
		return false
	}
	d567, _ := strconv.Atoi(strnum[4:7])
	if d567%7 != 0 {
		return false
	}
	d678, _ := strconv.Atoi(strnum[5:8])
	if d678%11 != 0 {
		return false
	}
	d789, _ := strconv.Atoi(strnum[6:9])
	if d789%13 != 0 {
		return false
	}
	d890, _ := strconv.Atoi(strnum[7:10])
	if d890%17 != 0 {
		return false
	}
	fmt.Println(strnum)
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
