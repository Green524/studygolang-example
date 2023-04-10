package main

import "fmt"

func main() {
	ch := make(chan string, 3)
	ch <- "A"
	ch <- "B"
	fmt.Println(len(ch))
	ch <- "C"
	fmt.Println(cap(ch))
	ch <- "D"
}
