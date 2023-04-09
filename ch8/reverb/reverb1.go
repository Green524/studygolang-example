package reverb

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Println("进来了")

}

func HandleConn1(c net.Conn) {
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
}

func MustCopy(w io.Writer, r io.Reader) {
	//io.Copy(w, r) 是循环且阻塞的，这是tcp连接不会自动断开的原因
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}
