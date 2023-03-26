package main

import (
	"fmt"
	"syscall"
)

type test struct {
	s string
}

func New(s string) *test {
	return &test{s: s}
}

func main() {
	//err := errors.New("自定义错误")
	//fmt.Println(err.Error())
	//fmt.Println(errors.New("EOF") == errors.New("EOF")) // "false"
	//
	//fmt.Printf("%p\n", New("hello"))
	//fmt.Printf("%p\n", New("hello"))
	//fmt.Println(New("hello") == New("hello")) //false
	var err error = syscall.Errno(3)
	fmt.Println(err)
	fmt.Println(err.Error())
}
