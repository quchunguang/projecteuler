package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// file: "A","B","C", -> []string{"A", "B", "C"}. Trailing comer MUST have!
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

// file: 12,23,34, -> []byte{12, 23, 34}. Trailing comer MUST have!
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
		v, _ := strconv.Atoi(line[0 : len(line)-1]) // XXX,  -->  XXX
		ret = append(ret, byte(v))
	}
	return
}

// file: 4H KH KS -> [][]string{{"4H", "KH", "KS"}, ...}
func CSWs(filename string, sep string) (table [][]string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		item := strings.Split(line[:len(line)-1], sep) // trim '\n' at the end
		table = append(table, item)
	}
	return
}

// Read triangle data, separate by space
func SST(filepath string) (data [][]int) {
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		record := make([]int, i+1)
		la := strings.Fields(line)
		for j, item := range la {
			record[j], _ = strconv.Atoi(item)
		}
		data = append(data, record)
		i++
	}
	return
}
