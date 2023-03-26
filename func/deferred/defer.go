package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"testing"
	"time"
)

func bigSlowOperation() {
	defer trace("bigSlowOperation")() // don't forget the extra parentheses
	// ...lots of work…
	time.Sleep(10 * time.Second) // simulate slow operation by sleeping
}
func trace(msg string) func() {
	start := time.Now()
	log.Printf("enter %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func TestBigSlowOperation(t *testing.T) {
	//bigSlowOperation()

	_ = double(4)
	// Output:
	// "double(4) = 8"

	fmt.Println(triple(4)) // "12"
}

func double(x int) (result int) {
	defer func() { fmt.Printf("double(%d) = %d\n", x, result) }()
	return x + x
}

func triple(x int) (result int) {
	defer func() { result += x }()
	return double(x)
}

// Fetch downloads the URL and returns the
// name and length of the local file.
func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	local := path.Base(resp.Request.URL.Path)
	if local == "/" {
		local = "index.html"
	}
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	p := make([]byte, 0, 28)
	c, err := resp.Body.Read(p[len(p):cap(p)])
	//p, _ = ioutil.ReadAll(resp.Body)
	fmt.Println(string(p[:len(p)+c]))
	// Close file, but prefer error from Copy, if any.
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}

var url = flag.String("url", "localhost:8000", "请求一个URL并返回内容")

func main() {
	flag.Parse()
	fetch("http://localhost:8000/list")
	//p := make([]byte, 0, 512)
	//fmt.Println(len(p), cap(p))
	//fmt.Println(len(p[len(p):cap(p)]), cap(p[:]))
	//fmt.Println(len(p[:]), cap(p[:]))
}
