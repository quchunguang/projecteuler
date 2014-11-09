package projecteuler

import (
	"bufio"
	"github.com/ajstarks/svgo"
	"os"
	"strconv"
)

func PE81_svg(data [][]int, sd [][]datai) {
	// Create output file
	fo, err := os.Create("PE81.svg")
	if err != nil {
		panic(err)
	}
	defer fo.Close()
	w := bufio.NewWriter(fo)

	length := len(data)
	wi := 30
	hi := 10
	canvas := svg.New(w)
	canvas.Start(wi*length, hi*length)
	canvas.Rect(0, 0, wi*length, hi*length, canvas.RGB(0, 0, 0))
	for i, j := length-1, length-1; i >= 0 && j >= 0; {
		if sd[i][j].d == 0 {
			sd[i][j].d = 999
			j--
		} else {
			sd[i][j].d = 999
			i--
		}
	}
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if sd[i][j].d == 999 {
				canvas.Text(wi*i+9, hi*j, strconv.Itoa(data[i][j]), "fill:red;font-size:8pt;text-anchor:middle")
			} else {
				canvas.Text(wi*i+9, hi*j, strconv.Itoa(data[i][j]), "fill:white;font-size:8pt;text-anchor:middle")
			}
		}
	}
	canvas.End()

	if err = w.Flush(); err != nil {
		panic(err)
	}
}
