package main

import (
	"flag"
	"log"
	"net"
	"studygolang-example/ch8/reverb"
	"sync"
)

var port = flag.String("port", "8000", "端口号")

func main() {
	flag.Parse()
	log.Printf("端口号：%s", *port)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	var wg sync.WaitGroup
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		wg.Add(1)
		//go reverb.HandleConn(conn, timezone(*port))
		go reverb.HandleConn1(conn, &wg)
		log.Printf("wait close %v\n", conn.RemoteAddr())
		wg.Wait()
		err = conn.Close()
		log.Printf("closed %v", conn.RemoteAddr())
		if err != nil {
			log.Fatal(err)
			return
		}

	}
}

func timezone(port string) string {
	switch port {
	case "8010":
		return "US/Eastern"
	case "8020":
		return "Asia/Tokyo"
	case "8030":
		return "Europe/London"
	default:
		return "Asia/China"
	}
}
