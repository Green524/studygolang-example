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
}
func gcd(x, y int) int {
	for y != 0 {
		//1,2 = 2, 1 % 2
		x, y = y, x%y
		fmt.Println(x, y)
	}
	return x
}
