package main

import (
	"fmt"
	"runtime"
	"strconv"
	"unsafe"
)

func main() {
	fmt.Println(5.0 / 4.0)
	var u uint8 = 255
	fmt.Println(u, u+1, u*u)
	fmt.Println(1 ^ 0)
	fmt.Println(1 ^ 1)

	var x uint8 = 1<<1 | 1<<5
	//0000 0010
	//0010 0000
	//0010 0010
	var y uint8 = 1<<1 | 1<<2
	//0000 0010
	//0000 0100
	//0000 0110
	fmt.Printf("%08b\n", x)
	fmt.Printf("%08b\n", y)

	fmt.Printf("%08b\n", 3&^1)

	for i := uint(0); i < 8; i++ {
		//fmt.Println(1 << i)
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	medals := []string{"gold", "silver", "bronze"}
	for i := len(medals) - 1; i >= 0; i-- {
		fmt.Println(medals[i]) // "bronze", "silver", "gold"
	}
	var a int = 1
	var b int8 = 2
	var c int16 = 3
	var d int32 = 4
	var e int64 = 5
	var f uint = 6
	var g uint64 = 7
	var h uint32 = 8
	fmt.Println(unsafe.Sizeof(a))
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println(unsafe.Sizeof(c))
	fmt.Println(unsafe.Sizeof(d))
	fmt.Println(unsafe.Sizeof(e))
	fmt.Println(unsafe.Sizeof(f))
	fmt.Println(unsafe.Sizeof(g))
	fmt.Println(unsafe.Sizeof(h))
	fmt.Println(runtime.GOARCH)
	fmt.Println(strconv.IntSize)
	//var apples int64 = 1
	//var oranges int64 = 2
	//var compote int = apples + oranges // compile error
	i := 1e100 // a float64
	j := int(f)
	fmt.Println(i, j)

	main1()
}

func main1() {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	x := int64(0xdeadbeef)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
	// Output:
	// 3735928559 deadbeef 0xdeadbeef 0XDEADBEEF
	ascii := 'a'
	unicode := '国'
	newline := '\n'
	fmt.Printf("%d %[1]c %[1]q\n", ascii)
	fmt.Printf("%d %[1]c %[1]q\n", unicode)
	fmt.Printf("%d %[1]c %[1]q\n", newline)
	//output:
	//97 a 'a'
	//22269 国 '国'
	//10
	//'\n'
	var f float32 = 16777216
	fmt.Println(f == f+1)

	const Avogadro = 6.02214129e23
	fmt.Printf("%f\n", Avogadro)

	var z float64
	fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"
}
