package main

import (
	"flag"
	"io"
	"log"
	"net"
	"studygolang-example/ch8/reverb"
	"time"
)

var port = flag.String("port", "8000", "端口号")

func main() {
	flag.Parse()
	log.Printf("端口号：%s", *port)
	listener, err := net.Listen("tcp", "localhost:"+*port)
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		//go handleConn(conn, timezone(*port))
		go reverb.HandleConn1(conn)
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
func handleConn(conn net.Conn, tz string) {
	defer conn.Close()
	for {
		str := tz + "=" + time.Now().Format("15:04:05\n")
		log.Println(str)
		_, err := io.WriteString(conn, str)
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
