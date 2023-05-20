package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			//"https://go.dev/",
			"https://godoc.org",
			//"https://play.golang.org",
			"http://gopl.io",
			//"https://go.dev/",
			"https://godoc.org",
			//"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

var HTTPGetBody = httpGetBody

var HTTPGetBodyDone = httpGetBodyDone

func httpGetBody(url string) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}
func httpGetBodyDone(url string, done chan<- struct{}) (interface{}, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

type M interface {
	Get(url string) (interface{}, error)
}

type M1 interface {
	GetDoneable(url string, done chan struct{}) (interface{}, error)
}

func Sequential(t *testing.T, m M) {
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
}
func Concurrent(t *testing.T, m M) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.Get(url)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes \n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)

	}
	n.Wait()
}
func ConcurrentDone(t *testing.T, m M1, done chan struct{}) {
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			start := time.Now()
			value, err := m.GetDoneable(url, done)
			if err != nil {
				log.Print(err)
			}
			fmt.Printf("%s,%s,%d bytes \n", url, time.Since(start), len(value.([]byte)))
			n.Done()
		}(url)

	}
	n.Wait()
}
