package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/axgle/mahonia"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

var min = flag.Int("min", 1, "随机数最小值，单位(ms)")
var max = flag.Int("max", 20000, "随机数最大值，单位(ms)")
var file = flag.String("f", "", "需要修改的文件")
var sec = flag.Int64("s", 2, "多少秒刷新一次")
var props = make(map[string]string)
var paragraph string
var decoder = mahonia.NewDecoder("gbk")

func main() {
	flag.Parse()
	log.Printf("区间 %d-%d", *min, *max)
	if *file == "" {
		log.Println("-f the parameter is empty.")
		return
	}
	loadProps(*file)
	ticker := time.NewTicker(time.Duration(*sec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			writeProps(*file)
		}
	}
}

func randNum(min int, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func writeProps(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0777)
	if err != nil {
		log.Fatalf("open file failed: %s \n", err.Error())
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(paragraph)
	if err == io.EOF {
		log.Println("write paragraph fail. ")
		return
	}
	err = writer.Flush()

	for k, v := range props {
		newValue := fmt.Sprintf("%s=%s\n", k, v)
		convK := decoder.ConvertString(k)
		if convK == "视频间隔" {
			randNum := randNum(*min, *max)
			newValue = fmt.Sprintf("%s=%d\n", k, randNum)
			log.Printf("write file: %s=%d\n", convK, randNum)
		}
		_, err = writer.WriteString(newValue)
		if err == io.EOF {
			log.Printf("write %s fail. \n", newValue)
			return
		}
		err = writer.Flush()
		if err != nil {
			log.Fatalf("flush file failed: %s \n", err.Error())
		}
	}
}

func loadProps(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("open file failed: %s \n", err.Error())
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	i := 0
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			return
		}
		i++
		lineKV := strings.Split(string(line), "=")
		if len(lineKV) >= 2 {
			//k := lineKV[0]
			//v := lineKV[0]
			props[lineKV[0]] = lineKV[1]
		} else {
			if strings.ContainsRune(lineKV[0], '[') && strings.ContainsRune(lineKV[0], ']') {
				paragraph = lineKV[0] + "\n"
			}
		}
	}
}
