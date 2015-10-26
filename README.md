projecteuler
============

Solve problems from [Project Euler](http://projecteuler.net) with golang.
Just for helping sleep.

Install
--------

```sh
go get github.com/quchunguang/projecteuler
go install github.com/quchunguang/projecteuler/projecteuler
```

Run
---

Run the solver (problem 10 for example),

```sh
projecteuler -id 10
```

For the problem description, (problem 10 for example), run

```sh
godoc github.com/quchunguang/projecteuler PE10
```

You may also direct access official version at [Problem 10](https://projecteuler.net/problem=10).

Get command line details by ` projecteuler -h `.

```
Usage of projecteuler:
  -f string
        Additional data file. Only the first one works in [-n|-f].
        (default target to the data file come with source)
  -h    Usage information.
  -id int
        Problem Id. (default 1)
  -n int
        N. Only the first one works in [-n|-f]. (default is the problem setting, depend on problem id given) (default -1)

IMPORT: Ensure there is a newline at the end of the file if the file is
 downloaded from projecteuler.org directly.
```
