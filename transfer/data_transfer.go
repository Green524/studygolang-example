package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//data, err := ioutil.ReadFile("D:\\IdeaProjects\\datamigratio-script\\2022-03-11（中文）.log")
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "发生错误%v", err)
	//}
	//for _, line := range strings.Split(string(data), "\n") {
	//	fmt.Println(line)
	//}
	reader()
}

func reader() {
	file, err := os.Open("D:\\IdeaProjects\\datamigratio-script\\logs\\2022-03-11（中文）.log")
	if err != nil {
		fmt.Fprintf(os.Stderr, "发生错误%v", err)
	}
	reader := bufio.NewReader(file)
	line, _, _ := reader.ReadLine()
	for line != nil {
		fmt.Println(string(line))
		line, _, _ = reader.ReadLine()

	}
}
