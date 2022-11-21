package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	dup2()
}

func dup3() {
	counts := make(map[string]int)
	fmt.Println(os.Args)
	for _, filename := range os.Args[1:] {
		data, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3:%v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for index, line := range counts {
		fmt.Printf("%d\t%s\n", line, index)
	}
}

func dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(counts, os.Stdin)
	} else {
		for _, file := range files {
			f, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(counts, f)
			f.Close()
		}
	}

	for line, n := range counts {
		fmt.Printf("%d\t%s\n", n, line)
	}
}

func countLines(counts map[string]int, f *os.File) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()

		if counts[line] > 0 {
			fmt.Printf("重复行：%s\t文件:%s\n", line, f.Name())
		} else {
			counts[line]++
		}
	}
}
