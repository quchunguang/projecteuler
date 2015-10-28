// Command line tool for calling github.com/quchunguang/projecteuler solver.
package main

import (
	"flag"
	"fmt"
	"github.com/quchunguang/projecteuler"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
)

const maxid = 521
const pkgPath = "github.com/quchunguang/projecteuler"
const cmdPath = pkgPath + "/projecteuler"
const godocName = "go doc"

type options struct {
	id    int
	n     int
	file  string
	help  bool
	about bool
}

// Print the default command line with given project id.
func PrintInfo(id int) {
	fmt.Println("Project id:\t\t", id)

	fName := "PE" + strconv.Itoa(id)
	fmt.Println("Project description:\t", godocName, pkgPath, fName)

	// Entry function of solver
	callType := reflect.ValueOf(projecteuler.Solvers[id].Caller).
		Type().String()
	strf := strings.Replace(callType, "func", "PE"+strconv.Itoa(id), 1)
	fmt.Println("Solver function:\t", strf)

	// Demo command line with default arguments
	if projecteuler.Solvers[id].Arg == nil {
		fmt.Println("Calling Command:\t", "projecteuler -id", id)
	} else if value, ok := projecteuler.Solvers[id].Arg.(string); ok {
		value = filepath.Join(os.Getenv("GOPATH"), "src", cmdPath, value)
		fmt.Println("Calling Command:\t", "projecteuler -id", id, "-file", value)
	} else if value, ok := projecteuler.Solvers[id].Arg.(int); ok {
		fmt.Println("Calling Command:\t", "projecteuler -id", id, "-n", value)
	} else {
		fmt.Println("[ERROR] BUG, Not supported argument type.")
		os.Exit(5)
	}

	fmt.Println("Comments:\t\t", projecteuler.Solvers[id].Comment)
	fmt.Println("Solved:\t\t\t", projecteuler.Solvers[id].Finished)

	// Total collections
	fmt.Println("\nTotal Projects:\t\t", maxid)
	totalSolved := 0
	for _, i := range projecteuler.Solvers {
		if i.Finished {
			totalSolved++
		}
	}
	fmt.Println("Total Solved:\t\t", totalSolved)
	fmt.Printf("Finished (%%):\t\t %4.1f\n",
		float32(totalSolved)/float32(maxid)*100.0)
}

// Call a solver function given project id and argument.
// If there is one argument, it could be any type.
// If pass nil, means using default argument given in `projecteuler.Solvers` or the solver
// function need no argument at all.
func Call(id int, arg interface{}) int {
	if projecteuler.Solvers[id].Arg != nil && arg == nil {
		arg = projecteuler.Solvers[id].Arg
		if value, ok := arg.(string); ok {
			// check if the argument is a file
			if strings.HasSuffix(value, ".txt") {
				p := filepath.Join(os.Getenv("GOPATH"), "src", cmdPath, value)
				if !ExistPath(p) {
					fmt.Println("[ERROR] Parameter not a valid path.")
					flag.Usage()
					os.Exit(1)
				}
				arg = p
			}
		}
	}
	f := reflect.ValueOf(projecteuler.Solvers[id].Caller)
	nArg := f.Type().NumIn()
	if nArg == 0 && arg != nil || nArg == 1 && arg == nil || nArg > 1 {
		fmt.Println("[ERROR] The number of parameters is not adapted.")
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
		fmt.Println("[ERROR] -file: No such file!")
		return false
	}
	if finfo.IsDir() {
		fmt.Println("[ERROR] -file: Not a file!")
		return false
	}
	return true
}

func main() {
	var opts options

	flag.IntVar(&opts.id, "id", 1, "Project id.")

	flag.IntVar(&opts.n, "n", -1, "N. Only the first one works in [-n|-file]. (default is the project setting, depend on project id given)")

	flag.StringVar(&opts.file, "file", "", "Additional data file. Only the first one works in [-n|-file]. (default target to the data file come with source)")

	flag.BoolVar(&opts.help, "h", false, "Usage information. IMPORT: Ensure there is a newline at the end of the file if the file is downloaded from projecteuler.org directly.")

	flag.BoolVar(&opts.about, "about", false, "Print the default command line with given project id.")

	// parse command line arguments
	flag.Parse()

	// process arguments -h
	if opts.help {
		flag.Usage()
		return
	}

	// check project id
	if opts.id < 1 || opts.id >= len(projecteuler.Solvers) ||
		!projecteuler.Solvers[opts.id].Finished {
		fmt.Println("[ERROR] No such project id or net solved yet!")
		os.Exit(3)
	}

	// process argument -about
	if opts.about {
		PrintInfo(opts.id)
		return
	}

	// process arguments -n -file
	var arg interface{}
	if opts.n != -1 {
		arg = opts.n
	} else if opts.file != "" {
		p := opts.file
		if !filepath.IsAbs(p) {
			abs, _ := os.Getwd()
			p = filepath.Join(abs, p)
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
	answer := Call(opts.id, arg)
	fmt.Println(answer)
}
