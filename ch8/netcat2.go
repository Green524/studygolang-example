package main

import (
	"log"
	"net"
	"os"
	"studygolang-example/ch8/reverb"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	go reverb.MustCopy(os.Stdout, conn)
	reverb.MustCopy(conn, os.Stdin)
}
