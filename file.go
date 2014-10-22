package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

//////
// file: "A","B","C", -> []string. Trailing comer MUST have!
func CSW(filename string) (words []string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString(',')
		if err == io.EOF {
			break
		}
		words = append(words, line[1:len(line)-2]) // "XXX",  -->  XXX
	}
	return
}

//////
// file: 12,23,34, -> []string. Trailing comer MUST have!
func CSV(filename string) (ret []byte) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString(',')
		if err == io.EOF {
			break
		}
		v, _ := strconv.Atoi(line[0 : len(line)-1])
		ret = append(ret, byte(v)) // "XXX",  -->  XXX
	}
	return
}
