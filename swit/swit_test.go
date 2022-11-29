package main

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	//f, err := os.Open("")
	//f, err = os.Create("")
	//fmt.Println(f, err)
	//fmt.Println(&Ri)
	//Ri := 10
	//fmt.Println(&Ri)
	//x := 1
	//p := &x
	//fmt.Println(p)
	//fmt.Println(*p) // "1"
	//*p = 2          // equivalent to x = 2
	//fmt.Println(x)  //

	//var x, y int
	//fmt.Println(&x == &x, &x == &y, &x == nil) // "true false false"

	//fmt.Println(p1)

	//fmt.Println(g())

	//fmt.Println(f() == f())
	//var v = 1
	//incr(&v)
	//fmt.Println(incr(&v))

	var v = 1
	fmt.Println(&v)
	incrCopy(v)
	fmt.Println(incrCopy(v)) //传递到函数的v为副本
}
func incr(v *int) int {
	*v++
	return *v
}

func incrCopy(v int) int {
	fmt.Println(&v)
	v++
	return v
}

var p1 = f()

func f() *int {
	v := 1
	//返回指针v
	fmt.Println(&v)
	return &v
}

func g() int {
	x := 1
	a := &x
	return *a
}
