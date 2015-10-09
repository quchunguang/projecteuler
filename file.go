// By default, txt files from http://projecteuler.org missed a newline character at the end of the files.
// That will cause issues when read files line by line.
// I add newline character for all txt files as well.
// Sublime Text will do this by simply open and save.
package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Read line by line from text file and feed to a chan of string
func ReadLine(filename string, lines chan string) {
	file, err := os.Open(filename)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
	}
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			close(lines)
			return
		}
		lines <- line[:len(line)-1]
	}
}

// "A","B","C" -> []string{"A", "B", "C"}, surround is true if quote surrounded.
func ReadWords(filename, sep string, surround bool) (ret []string) {
	lines := make(chan string)
	go ReadLine(filename, lines)
	for line := range lines {
		svs := strings.Split(line, sep)
		for _, item := range svs {
			if surround {
				ret = append(ret, item[1:len(item)-1]) // "XXX",  -->  XXX
			} else {
				ret = append(ret, item) // XXX,  -->  XXX
			}
		}
	}
	return
}

// file: 4H KH KS -> [][]string{{"4H", "KH", "KS"}, ...}, surround is true if quote surrounded.
func ReadMatrixWords(filename, sep string, surround bool) (ret [][]string) {
	lines := make(chan string)
	go ReadLine(filename, lines)
	for line := range lines {
		svs := strings.Split(line, sep)
		record := make([]string, len(svs))
		for i, item := range svs {
			if surround {
				record[i] = item[1 : len(item)-1] // "XXX"  -->  XXX
			} else {
				record[i] = item
			}
		}
		ret = append(ret, record)
	}
	return
}

// 12,23,34 -> []int{12, 23, 34}.
func ReadInts(filename, sep string) (ret []int) {
	lines := make(chan string)
	go ReadLine(filename, lines)
	for line := range lines {
		svs := strings.Split(line, sep)
		for _, item := range svs {
			v, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
				return
			}
			ret = append(ret, v) // XXX,  -->  XXX
		}
	}
	return
}

// 12,23\n45,56 -> [][]int{{12,23},{45,56}}.
func ReadMatrixInts(filename, sep string) (ret [][]int) {
	var err error
	lines := make(chan string)
	go ReadLine(filename, lines)
	for line := range lines {
		svs := strings.Split(line, sep)
		record := make([]int, len(svs))
		for i, item := range svs {
			record[i], err = strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
		ret = append(ret, record)
	}
	return
}
