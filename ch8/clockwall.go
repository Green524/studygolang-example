package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func main() {
	//args := os.Args[1:]
	args := []string{"NewYork=localhost:8010", "Tokyo=localhost:8020", "London=localhost:8030"}
	a := parseArgs(args)
	chs := make([]chan string, len(a))
	i := 0
	for _, v := range a {
		chs[i] = make(chan string)
		go connect(chs[i], v)
		i++
	}
	for {
		for i := range chs {
			fmt.Print(<-chs[i] + "\t")
		}
		fmt.Println()
	}
}

func connect(ch chan string, address string) {
	defer close(ch)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	sc := bufio.NewScanner(conn)
	for sc.Scan() {
		ch <- sc.Text()
	}
	time.Sleep(1 * time.Second)
}
func parseArgs(args []string) (a map[string]string) {
	a = make(map[string]string)
	for _, arg := range args {
		//NewYork=localhost:8010
		var mapping = strings.Split(arg, "=")
		a[mapping[0]] = mapping[1]
	}
	return a
}
