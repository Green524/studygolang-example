package main

import (
	"bufio"
	"flag"
	"fmt"
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
var paragraph = ""

func main() {
	flag.Parse()
	log.Printf("区间 %d-%d", *min, *max)
	if *file == "" {
		log.Println("-f the parameter is empty.")
		return
	}
	loadProps(*file)
	fmt.Println(props)
	ticker := time.NewTicker(time.Duration(*sec) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			fmt.Println("<-ticker.C")
			writeProps(*file)
		}
	}
}

func randNum(min int, max int) int {
	rand.NewSource(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}
func writeProps(path string) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	fmt.Println("OpenFile")
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
		if k == "视频间隔" {
			newValue = fmt.Sprintf("%s=%d\n", k, randNum(*min, *max))
			log.Printf("write file: %s\n", newValue)
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
			log.Printf("read %s fail. \n", line)
			return
		}
		i++
		lineKV := strings.Split(string(line), "=")
		if len(lineKV) >= 2 {
			props[lineKV[0]] = lineKV[1]
		} else {
			if strings.ContainsRune(lineKV[0], '[') && strings.ContainsRune(lineKV[0], ']') {
				paragraph = lineKV[0] + "\n"
			}
		}
	}
}
