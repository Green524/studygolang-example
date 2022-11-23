package main

import (
	"fmt"
	"testing"
)

var global *int

func TestNew(t *testing.T) {
	//返回的是int类型的指针
	p := new(int)
	fmt.Println(*p)
	*p = 2
	fmt.Println(*p)
	fmt.Println(newInt())

	a1 := new(int)
	a2 := new(int)
	fmt.Println(a1 == a2)

	//runtime.SetFinalizer()

	fmt.Println(delta(1, 2))
}

func f() {
	var x int
	x = 1
	//局部的x变量逃逸到函数外
	//此时x不会被go的GC回收，因为x的指针被保存到全局变量global
	global = &x
}
func g() {
	/*
		局部的y变量没有逃逸到函数外
		虽然是new，但是会被go选择在栈上或堆上分配*y的空间，y变量在栈，然后又GC回收
		逃逸的变量需要分配额外的空间，对性能有着细微的影响
	*/
	y := new(int)
	*y = 1

}

//会报错cannot call non-function new (type int), declared at .\new_test.go:24:17
func delta(old, new int) int {
	new(int)
	return new - old
}

//func newInt() *int {
//	b := 1
//	*&b = 2
//	fmt.Println(b)
//	return new(int)
//}

func newInt() *int {
	var dummy int
	return &dummy
}
