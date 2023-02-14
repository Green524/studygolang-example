package _func

import (
	"fmt"
	"testing"
)

//匿名函数

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

func TestSquares(t *testing.T) {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
}
