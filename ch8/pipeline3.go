package main

import (
	"fmt"
	"time"
)

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go Counter(naturals)
	go Squarer(naturals, squares)
	Printer(squares)
}

func Counter(out chan<- int) {
	for x := 0; x < 11; x++ {
		out <- x
		time.Sleep(1 * time.Second)
	}
	close(out)
}
func Squarer(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
		fmt.Println("计算：", x)
	}
	close(out)
}
func Printer(in <-chan int) {
	for x := range in {

		fmt.Println(x)
	}
}
