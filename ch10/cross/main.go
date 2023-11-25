package main

import (
	"fmt"
	"reflect"
	"runtime"
)

//文档
func main() {
	fmt.Println(runtime.GOOS, runtime.GOARCH)
	v := reflect.ValueOf(3)
	fmt.Println(v.Type())
	fmt.Println(reflect.TypeOf(v.Int()))
}
