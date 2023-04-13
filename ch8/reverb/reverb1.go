package reverb

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	wg.Done()
	//if tcpConn, ok := c.(*net.TCPConn); ok {
	//	log.Fatal("关闭")
	//	tcpConn.CloseWrite()
	//}

}

func HandleConn1(c net.Conn, wg *sync.WaitGroup) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second, wg)
	}
}
func HandleConn(conn net.Conn, tz string) {
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
func MustCopy(w io.Writer, r io.Reader) {
	//io.Copy(w, r) 是循环且阻塞的，这是tcp连接不会自动断开的原因
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
