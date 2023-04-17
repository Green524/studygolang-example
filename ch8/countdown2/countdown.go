package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		abort <- struct{}{}
	}()
	//!+
	fmt.Println("Commencing countdown.  Press return to abort.")
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("倒数完成")
		// Do nothing.
	case <-abort:
		fmt.Println("Launch aborted!")
		return
	}
	launch()

	//ch := make(chan int, 2)
	//for i := 0; i < 10; i++ {
	//	select {
	//	case x := <-ch:
	//		fmt.Println(x) // "0" "2" "4" "6" "8"
	//	case ch <- i:
	//	}
	//}
}
func launch() {
	fmt.Println("Lift off!")
}
