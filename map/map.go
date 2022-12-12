package main

import (
	"fmt"
	"reflect"
)

func main() {
	args := make(map[string]int)
	fmt.Println(reflect.TypeOf(args))
	args1 := map[string]int{
		"name": 1,
		"age":  2,
	}
	fmt.Println(reflect.TypeOf(args1))
	fmt.Println(args1)
	delete(args1, "name")
	fmt.Println(args1)
	args1["weight"] = 150
	for key, value := range args1 {
		fmt.Println(key, value)

	}
}
