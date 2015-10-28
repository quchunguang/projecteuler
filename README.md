projecteuler
============

Solve projects from [Project Euler](http://projecteuler.net) with golang.
Just for helping sleep.

Install
--------

```sh
go get github.com/quchunguang/projecteuler
go install github.com/quchunguang/projecteuler/projecteuler
```

Run
---

Run the solver (project id=10 for example), `projecteuler -id 10`

Get information of the solver (project id=10 for example), `projecteuler -id 10 -about`

For the project description, (project id=10 for example), `godoc github.com/quchunguang/projecteuler PE10`

You may also direct access official version at [Problem 10](https://projecteuler.net/project=10).

Get details of `projecteuler` by `projecteuler -h`.

```
Usage of projecteuler:
  -f string
        Additional data file. Only the first one works in [-n|-f].
        (default target to the data file come with source)
  -h    Usage information.
  -id int
        Problem Id. (default 1)
  -n int
        N. Only the first one works in [-n|-f].
        (default is the project setting, depend on project id given)

NOTE: (TODO)Ensure there is a newline at the end of the file if the file is
 downloaded from projecteuler.org directly.
```
