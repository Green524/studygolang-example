package main

import (
	"flag"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var verbose = flag.Bool("v", false, "show verbose progress messages")
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}
func main() {

	start := time.Now().Second()
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
	}()
	fileSizes := make(chan int64)
	var wg sync.WaitGroup
	for _, root := range roots {
		wg.Add(1)
		go walkDir(root, fileSizes, &wg)
	}
	go func() {
		wg.Wait()
		close(fileSizes)
	}()
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Microsecond)
	}
	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			log.Println("关闭maingoroutine")
			for range fileSizes {

			}
			return
		case size, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += size
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
	log.Printf("消耗%d秒", time.Now().Second()-start)
}
func printDiskUsage(nfiles, nbytes int64) {
	fmt.Printf("%d files  %.1f GB\n", nfiles, float64(nbytes)/1e9)
}
func walkDir(dir string, fileSizes chan<- int64, wg *sync.WaitGroup) {
	time.Sleep(100 * time.Millisecond)
	defer wg.Done()
	if cancelled() {
		log.Println("关闭walkDir goroutine")
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			wg.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir(subdir, fileSizes, wg)
		} else {
			fileSizes <- entry.Size()
		}
	}
}

var sema = make(chan struct{}, 20)

func dirents(dir string) []fs.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
		return nil

	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	//<-sema
	return entries
}
