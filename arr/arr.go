package main

import (
	"fmt"
	"reflect"
)

type Currency int

const (
	USD Currency = iota // 美元
	EUR                 // 欧元
	GBP                 // 英镑
	RMB                 // 人民币
)

func main() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB]) // "3 ￥"

	r := [...]int8{99: -1}
	fmt.Println(r[0])

	a := [2]int{1, 2}
	b := [...]int{1, 2}
	//c := [...]int{1, 2, 3}
	d := [2]int{1, 3}
	fmt.Println(a == b, a == d)

	//第一个[]是slice，第二个[]数组
	var arr [][1]int = [][1]int{{}} //必须初始化slice中的元素，否则越界
	fmt.Println(reflect.TypeOf(arr))
	fmt.Println(reflect.TypeOf(arr[0]))
	//arr1 := []int{3}
	//arr[0] = arr1
	fmt.Println(arr)

	//slice不会像数组一样初始化为对应类型的零值
	var c []int
	fmt.Println(c[0]) //panic: runtime error: index out of range [0] with length 0
}
