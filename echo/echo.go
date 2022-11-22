package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")

var sep = flag.String("s", " ", "separator")

/*
  -n    omit trailing newline
  -s string
        separator (default " ")

*/
func main() {
	fmt.Println(*n, *sep)
	flag.Parse()
	fmt.Println(*n, *sep)
	fmt.Println(strings.Join(flag.Args(), *sep))
	if !*n {
		fmt.Println()
	}
}
