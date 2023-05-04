package main

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	dir := "D:\\BaiduNetdiskDownload"
	ch := make(chan DirInfo)
	go walkDir(dir, ch)

	result := make(map[string]int64)
	for {
		dirinfo := <-ch
		e, ok := result[dirinfo.dirname]
		if ok {
			e += dirinfo.nbytes
			result[dirinfo.dirname] = e
		} else {
			result[dirinfo.dirname] = e
		}
	}
}

type DirInfo struct {
	dirname string
	nbytes  int64
}

func walkDir(dir string, ch chan<- DirInfo) {
	entries := dirents(dir)
	for _, entrie := range entries {

		if entrie.IsDir() {
			subdir := filepath.Join(dir, entrie.Name())
			walkDir(subdir, ch)
		} else {
			subdir := filepath.Join(dir, entrie.Name())
			var dirinfo = DirInfo{subdir, entrie.Size()}
			ch <- dirinfo
		}
	}
}

func dirents(dir string) []fs.FileInfo {
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprint(os.Stderr, "\v\t", err)
		return nil
	}
	return entries
}
