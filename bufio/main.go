package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//counts[input.Text()]++
		line := input.Text()
		fmt.Println(counts[line])
		fmt.Println(counts[line])
		counts[line] = counts[line] + 1
		fmt.Println(counts[line])
	}
	fmt.Println(counts)
	for line, n := range counts {
		fmt.Println(line, n)
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
