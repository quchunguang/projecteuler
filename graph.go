package projecteuler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

//////
func FindPathMax(data [][]int, i, j int) (ret int) {
	if i == len(data)-1 {
		return data[i][j]
	}
	l := FindPathMax(data, i+1, j)
	r := FindPathMax(data, i+1, j+1)
	if l > r {
		ret = l + data[i][j]
	} else {
		ret = r + data[i][j]
	}
	return
}

type datai struct {
	sm    int  // Biggest sum
	right bool // Previous item go to here from right path?(or left)
}

func FindPathMax2(data [][]int) (ret int) {
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

func MaxPathSum(filepath string, N int) (ret int) {
	// Create 2D slice
	data := make([][]int, N)
	for i := 0; i < N; i++ {
		data[i] = make([]int, i+1)
	}

	// Read file and fill slice
	file, err := os.Open(filepath)
	defer file.Close()
	if err != nil {
		log.Fatal(err)
	}
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		la := strings.Fields(line)
		for j, item := range la {
			data[i][j], _ = strconv.Atoi(item)
		}
		i++
	}
	// Find biggest path
	// ret = FindPathMax(data, 0, 0)
	ret = FindPathMax2(data)
	return
}
