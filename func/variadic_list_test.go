package _func

import (
	"fmt"
	"testing"
)

func sum(vals ...int) (total int) {
	fmt.Printf("%T\n", vals)
	for _, e := range vals {
		total += e
	}
	return
}

func TestSum(t *testing.T) {
	//如下参数列表行为将参数转变为一个数组，再将数组的切片传进去
	fmt.Println(sum(1, 2))
	fmt.Printf("%T\n", f) // "func(...int)"
	fmt.Printf("%T\n", g) // "func([]int)"
}
func f(...int) {}
func g([]int)  {}
