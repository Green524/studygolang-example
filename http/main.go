package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	for _, url := range os.Args[1:] {
		go ConFetch(url, ch)
	}
	for range os.Args[1:] {
		fmt.Println(<-ch)
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func ConFetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("fetch:%v\n", err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs   %d   %s\n    %d", secs, nbytes, url, nbytes)
}

func Fetch() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, "http") && !strings.HasPrefix(url, "https") {
			url = "http://" + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "status:%s\tfetch:%v\n", resp.Status, err)
			os.Exit(1)
		}
		//b, err := ioutil.ReadAll(resp.Body)
		b, err := io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "status:%s\tfetch: reading %s,%v\n", resp.Status, url, err)
		}
		fmt.Printf("%d", b)
	}
}
