package main

import (
	"fmt"
	"math/rand"
	"time"
)

var arr [3]string = [3]string{"heads", "tails", "default"}
var Ri = 1

func main() {
	fmt.Println(Signum(-1))
	fmt.Println(Signum(0))
	fmt.Println(Signum(1))

	coinflip()
}

type Point struct {
	X, Y int
}

var p Point

const abc int = 1

func Signum(x int) int {
	//go 的switch自带return,如果需要case往下走需要手动加关键字 fallthrough
	//switch true 和其它语言中 if() else if() 类似,会逐个判断
	switch true {
	case x > 0:
		fallthrough
	default:
		return 0
	case x < 0:
		return -1
	}
}

func coinflip() {
	var heads, tails int
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3)
	switch arr[num] {
	case "heads":
		heads++
		fmt.Println("heads")
	case "tails":
		tails++
		fmt.Println("tails")
	default:
		fmt.Println("landed on edge!")
	}
}
