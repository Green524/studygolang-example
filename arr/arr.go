package main

import "fmt"

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

}
