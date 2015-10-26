// Command line tool for calling github.com/quchunguang/projecteuler solver.
package main

import (
	"flag"
	"fmt"
	"github.com/quchunguang/projecteuler"
	"os"
	"path"
	"reflect"
	"strings"
)

// Id   option -id target the problem to run.
// N    option -n given N (OPTIONEL).
// File option -f file path to the datafile (OPTIONEL).
type Options struct {
	Id   int
	N    int
	File string
}

// Functions projecteuler.PExxx() will get one or no argument and could be
// any type, return value MUST be int type and no more, holding the answer.
type Solver struct {
	Caller   interface{} // Function handle of solver.
	Arg      interface{} // Default argument used by solver.
	Finished bool        // If the problem had solved.
}

// List all solver function handler and default argument.
var Solvers = []Solver{
	{nil, nil, false}, // 0 - Hold place, No this problem!
	{projecteuler.PE1, int(1e3), true},
	{projecteuler.PE2, int(4e6), true},
	{projecteuler.PE3, int(600851475143), true},
	{projecteuler.PE4, nil, true},
	{projecteuler.PE5, int(20), true},
	{projecteuler.PE6c, int(100), true},
	{projecteuler.PE7, int(10001), true},
	{projecteuler.PE8, int(13), true},
	{projecteuler.PE9, int(1000), true},
	{projecteuler.PE10b, nil, true},
	{projecteuler.PE11, nil, true},
	{projecteuler.PE12, int(500), true},
	{projecteuler.PE13, "p013_bignumbers.txt", true},
	{projecteuler.PE14, int(1e6), true},
	{projecteuler.PE15, int(20), true},
	{projecteuler.PE16, int(1000), true},
	{projecteuler.PE17, int(1000), true},
	{projecteuler.PE18, "p018_path.txt", true},
	{projecteuler.PE19b, nil, true},
	{projecteuler.PE20, int(100), true},
	{projecteuler.PE21, int(1e4), true},
	{projecteuler.PE22, "p022_names.txt", true},
	{projecteuler.PE23, nil, true},
	{projecteuler.PE24, int(1e6), true},
	{projecteuler.PE25, int(1000), true},
	{projecteuler.PE26, int(1000), true},
	{projecteuler.PE27, nil, true},
	{projecteuler.PE28, int(1001), true},
	{projecteuler.PE29, int(100), true},
	{projecteuler.PE30, int(5), true},
	{projecteuler.PE31, int(200), true},
	{projecteuler.PE32, nil, true},
	{projecteuler.PE33, nil, true},
	{projecteuler.PE34, nil, true}, // TODO: this function will never stop
	{projecteuler.PE35, int(1e6), true},
	{projecteuler.PE36, int(1e6), true},
	{projecteuler.PE37, nil, true},
	{projecteuler.PE38, nil, true},
	{projecteuler.PE39, int(1000), true},
	{projecteuler.PE40, int(1e6), true},
	{projecteuler.PE41, nil, true}, // TODO: this function will never stop
	{projecteuler.PE42, "p042_words.txt", true},
	{projecteuler.PE43, nil, true},
	{projecteuler.PE44, nil, true},
	{projecteuler.PE45, nil, true}, // TODO: this function will never stop
	{projecteuler.PE46, nil, true},
	{projecteuler.PE47, int(4), true},
	{projecteuler.PE48, int(1000), true},
	{projecteuler.PE49, nil, true},
	{projecteuler.PE50, int(1e6), true},
	{projecteuler.PE51, nil, true},
	{projecteuler.PE52, nil, true},
	{projecteuler.PE53, int(1e6), true},
	{projecteuler.PE54, "p054_poker.txt", true},
	{projecteuler.PE55, int(1e4), true},
	{projecteuler.PE56, int(100), true},
	{projecteuler.PE57, int(1000), true},
	{projecteuler.PE58, nil, true}, // TODO: Run time about 1 hour
	{projecteuler.PE59, "p059_cipher.txt", true},
	{projecteuler.PE60, nil, true}, // TODO: this function will never stop
	{projecteuler.PE61, nil, true}, // TODO: Only print out answer
	{projecteuler.PE62, nil, true},
	{projecteuler.PE63, nil, true},
	{projecteuler.PE64, int(10000), true},
	{projecteuler.PE65, nil, true},
	{projecteuler.PE66, int(1000), true},
	{projecteuler.PE67, "p067_triangle.txt", true},
	{projecteuler.PE68, nil, true}, // TODO: Only print out answer
	{projecteuler.PE69, int(1e6), true},
	{projecteuler.PE70, int(1e7), true}, // TODO: Run time about 1 hour at 83%.
	{projecteuler.PE71, int(1e6), true},
	{projecteuler.PE72, int(1e6), true},
	{projecteuler.PE73, int(12000), true}, // TODO: Run time about 5 minutes
	{projecteuler.PE74, int(1e6), true},
	{projecteuler.PE75, int(1500000), true},
	{projecteuler.PE76, int(100), true},
	{projecteuler.PE77, int(5000), true},
	{projecteuler.PE78, int(1e6), true},
	{projecteuler.PE79, "p079_keylog.txt", false},
	{projecteuler.PE80, nil, false},
	{projecteuler.PE81, "p081_matrix.txt", true},
	{projecteuler.PE82, "p082_matrix.txt", true},
	{projecteuler.PE83, "p083_matrix.txt", false},
	{projecteuler.PE84, nil, false},
	{projecteuler.PE85, int(2e6), true},
	{projecteuler.PE86, int(1e2), false},
	{projecteuler.PE87, nil, false},
	{projecteuler.PE88, nil, false},
	{projecteuler.PE89, nil, false},
	{projecteuler.PE90, nil, false},
	{projecteuler.PE91, nil, false},
	{projecteuler.PE92, nil, false},
	{projecteuler.PE93, nil, false},
	{projecteuler.PE94, nil, false},
	{projecteuler.PE95, nil, false},
	{projecteuler.PE96, "p096_sudoku.txt", false},
	{projecteuler.PE97, nil, true},
	{projecteuler.PE98, "p098_words.txt", true},
	{projecteuler.PE99, "p099_base_exp.txt", true},
	{projecteuler.PE100, nil, false},
}

// Call a solver function given problem Id and argument.
// If there is one argument, it could be any type.
// If pass nil, means using default argument given in `Solvers` or the solver
// function need no argument at all.
func Call(Id int, arg interface{}) int {
	if Solvers[Id].Arg != nil && arg == nil {
		arg = Solvers[Id].Arg
		if value, ok := arg.(string); ok {
			// check if the argument is a file
			if strings.HasSuffix(value, ".txt") {
				p := path.Join(os.Getenv("GOPATH"), "src", "github.com/quchunguang/projecteuler/projecteuler", value)
				if !ExistPath(p) {
					fmt.Println("[ERROR] Parameter not a valid path.")
					flag.Usage()
					os.Exit(1)
				}
				arg = p
			}
		}
	}
	f := reflect.ValueOf(Solvers[Id].Caller)
	nArg := f.Type().NumIn()
	if nArg == 0 && arg != nil || nArg == 1 && arg == nil || nArg > 1 {
		fmt.Println("[ERROR] The number of params is not adapted.")
		flag.Usage()
		os.Exit(2)
	}
	in := make([]reflect.Value, 1)
	var result []reflect.Value
	if arg != nil {
		in[0] = reflect.ValueOf(arg)
		result = f.Call(in)
	} else {
		result = f.Call(nil)
	}
	return int(result[0].Int())
}

// Check if given pathname is exist and target to a regular file.
func ExistPath(p string) bool {
	finfo, err := os.Stat(p)
	if err != nil {
		fmt.Println("[ERROR] -f: No such file!")
		return false
	}
	if finfo.IsDir() {
		fmt.Println("[ERROR] -f: Not a file!")
		return false
	}
	return true
}

func main() {
	// parse command line arguments
	var opts Options
	flag.IntVar(&opts.Id, "id", 1, "Problem Id.")
	flag.IntVar(&opts.N, "n", -1, "N. Only the first one works in [-n|-f]. (default is the problem setting, depend on problem id given)")
	flag.StringVar(&opts.File, "f", "", "Additional data file. Only the first one works in [-n|-f]. (default target to the data file come with source)")
	h := flag.Bool("h", false, "Usage information. IMPORT: Ensure there is a newline at the end of the file if the file is downloaded from projecteuler.org directly.")

	flag.Parse()

	if *h {
		flag.Usage()
		return
	}

	// check problem id
	if opts.Id < 1 || opts.Id > len(Solvers) || !Solvers[opts.Id].Finished {
		fmt.Println("[ERROR] No such problem or not solved yet!")
		flag.Usage()
		os.Exit(3)
	}

	// process arguments
	var arg interface{}
	if opts.N != -1 {
		arg = opts.N
	} else if opts.File != "" {
		p := opts.File
		if !path.IsAbs(p) {
			abs, _ := os.Getwd()
			p = path.Join(abs, p)
		}
		if !ExistPath(p) {
			flag.Usage()
			os.Exit(4)
		}
		arg = p
	} else {
		arg = nil
	}

	// calling solver
	answer := Call(opts.Id, arg)
	fmt.Println(answer)
}
