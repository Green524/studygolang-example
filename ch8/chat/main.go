package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string

var (
	entering = make(chan client)
	levaing  = make(chan client)
	messages = make(chan string)
	delay    = 5 * time.Second
)

func main() {
	lis, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	go broadcaster()
	for {
		conn, err := lis.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleConn(conn)
	}
}
func broadcaster() {
	clients := make(map[client]bool)
	for {
		select {
		case msg := <-messages:
			for cli := range clients {
				cli <- msg
			}
		case cli := <-entering:
			clients[cli] = true
		case cli := <-levaing:
			delete(clients, cli)
			close(cli)
		}
	}
}
func handleConn(conn net.Conn) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	//下面俩行代码的顺序区别在于是否在客户端连接时给客户端自己发送 who + " has arrived"
	entering <- ch
	messages <- who + " has arrived"

	input := bufio.NewScanner(conn)
	timer := time.NewTimer(delay)
	go autoClose(conn, who, timer)
	// 异步autoClose关闭连接之后 Scanner 扫描结果为false ，程序继续往下走
	for input.Scan() {
		messages <- who + ": " + input.Text()
		timer.Reset(delay)
	}
	messages <- who + " has left"
	levaing <- ch
	conn.Close()
}
func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg)
	}
}
func autoClose(conn net.Conn, who string, timer *time.Timer) {
	for {
		select {
		case <-timer.C:
			timer.Stop()
			fmt.Fprintln(conn, who+" timeout auto close!")
			conn.Close()
			//为什么不往channel发送信息，然后clientWriter写入客户端？
			//因为conn 流管道关闭操作可能比写入早
		default:
		}
	}
}
