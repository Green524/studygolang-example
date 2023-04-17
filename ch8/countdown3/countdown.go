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

	fmt.Println("Commencing countdown.  Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			fmt.Println("进阿里了")
			// Do nothing.
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
	//time.Tick(1 * time.Second) 只有在整个程序的生命周期都依赖它时才方便使用，不然容易goroutine泄露
	//更推荐使用以下方式
	ticker := time.NewTicker(1 * time.Second)
	<-ticker.C
	ticker.Stop()
}
func launch() {
	fmt.Println("Lift off!")
}
