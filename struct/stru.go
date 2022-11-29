package main

import (
	"fmt"
)

type Person struct {
	name string
}

var arr [3]string = [3]string{"heads", "tails", "default"}

func main() {
	person := Person{name: ""}
	person.name = "chenum"

	//arr[0] = "你好"
	arr[0] += "你好"

	fmt.Println(arr[0])
	var x, y int
	x = 1
	y = 2
	x, y = y, x
	fmt.Println(x, y)
	fmt.Println("-------")
	fmt.Println(gcd(1, 2))
	fmt.Println(1 % 2)

	fmt.Println("--------")
	fmt.Println(fib(3))
	a, b, c := 1, 2, 3
	fmt.Println(a, b, c)

	fmt.Println("--------")
	m := make(map[string]string)
	var ok bool
	var res string
	res, ok = m["a"]
	fmt.Println(ok, res)
	m["a"] = "cpq"
	_, ok = m[""]         // map返回2个值
	_, ok = m["a"], false // ok 有值，也是false
	fmt.Println(ok)
	_ = m[""] // map返回1个值
	//数据类型判断
	gtype(m)

	//下面语句包含隐式的赋值
	medals := []string{"gold", "sliver", "bronze"}
	//类似下面
	//medals[0] = "gold"
	//medals[1] = "sliver"
	//medals[2] = "bronze"
	fmt.Println(medals)
}
func gtype(i interface{}) {
	switch i.(type) {
	case string:
		break
	case int:
		break
	}
}

func gcd(x, y int) int {
	for y != 0 {
		//1,2 = 2, 1 % 2
		x, y = y, x%y
		fmt.Println(x, y)
	}
	return x
}

func fib(n int) int {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		fmt.Printf("x:%d y:%d x+y:%d\n", x, y, x+y)
		x, y = y, x+y
	}
	return x
}
