// Functions PEXXX_svg() here write out svg illustrations with name PEXXX.svg.
// Using web browsers with svg support (chrome, firefox, etc.) to open it.
// linux:   `x-www-browser PEXXX.svg`
package projecteuler

import (
	"bufio"
	"fmt"
	"github.com/ajstarks/svgo"
	"os"
)

func PE81_svg(data [][]int, sd [][]datai) {
	length := len(data)
	for i, j := length-1, length-1; i >= 0 && j >= 0; {
		if sd[i][j].d == LEFT_TO_HERE {
			sd[i][j].d |= SELECTED_FLAG
			j--
		} else {
			sd[i][j].d |= SELECTED_FLAG
			i--
		}
	}
	DrawMatrix("PE81.svg", data, sd)
}

func PE82_svg(data [][]int, sd [][]datai) {
	var min, mini int = MaxInt, 0
	length := len(data)
	for i := 0; i < length; i++ {
		if sd[i][length-1].sm < min {
			min = sd[i][length-1].sm
			mini = i
		}
	}
	for i, j := mini, length-1; j >= 0; {
		if sd[i][j].d == LEFT_TO_HERE {
			sd[i][j].d |= SELECTED_FLAG
			j--
		} else if sd[i][j].d == ABOVE_TO_HERE {
			sd[i][j].d |= SELECTED_FLAG
			i--
		} else if sd[i][j].d == BELOW_TO_HERE {
			sd[i][j].d |= SELECTED_FLAG
			i++
		} else {
			fmt.Println("err", i, j)
		}
	}
	DrawMatrix("PE82.svg", data, sd)
}

func DrawMatrix(filename string, data [][]int, sd [][]datai) {
	fo, err := os.Create(filename)
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
	var ss string
	for i := 0; i < length; i++ {
		for j := 0; j < length; j++ {
			if sd[i][j].d&LEFT_TO_HERE != 0 {
				ss = ">"
			}
			if sd[i][j].d&ABOVE_TO_HERE != 0 {
				ss = "∨"
			}
			if sd[i][j].d&BELOW_TO_HERE != 0 {
				ss = "∧"
			}
			if sd[i][j].d&RIGHT_TO_HERE != 0 {
				ss = "<"
			}
			if sd[i][j].d&SELECTED_FLAG != 0 {
				canvas.Text(wi*j+16, hi*i+8, ss,
					"fill:red;font-size:8pt;text-anchor:middle")
			} else {
				canvas.Text(wi*j+16, hi*i+8, ss,
					"fill:white;font-size:8pt;text-anchor:middle")
			}
		}
	}
	canvas.End()

	if err = w.Flush(); err != nil {
		panic(err)
	}
}
