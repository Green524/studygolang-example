package main

import (
	"io"
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
	//done := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		log.Println("done")
		//done <- struct{}{}
	}()
	reverb.MustCopy(conn, os.Stdin)
	conn.Close()
	//<-done
}
