package projecteuler

//////
type SortIntStr []string

func (s SortIntStr) Less(i, j int) bool {
	if len(s[i]) != len(s[j]) {
		return len(s[i]) < len(s[j])
	}
	for c := 0; c < len(s[i]); c++ {
		if s[i][c] == s[j][c] {
			continue
		}
		return s[i][c] < s[j][c]
	}
	return false // ==
}

func (s SortIntStr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s SortIntStr) Len() int {
	return len(s)
}

//////
type ByteSlice []byte

func (a ByteSlice) Len() int           { return len(a) }
func (a ByteSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByteSlice) Less(i, j int) bool { return a[i] < a[j] }
