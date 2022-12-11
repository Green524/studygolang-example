package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	c3 := [...]uint8{1, 2}
	arrAss(&c3)
	fmt.Printf("%T %[1]v\n", c3)

	fmt.Println(bitCount(c1, c2))

	flag.Parse()

	printSHA()
}

func arrAss(b *[2]uint8) {
	for i := 0; i < len(b); i++ {
		b[i] = 0
	}
}

//初始化为0
func zero(b *[2]int) {
	*b = [2]int{}
}

var pc = [256]byte{}

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

//练习 4.1： 编写一个函数，计算两个SHA256哈希码中不同bit的数目。（参考2.6.2节的PopCount函数。)
func bitCount(b1, b2 [32]byte) (count int) {
	for i := range b1 {
		tmp := b1[i] ^ b2[i]
		count += int(pc[tmp])
	}
	return
}

//练习 4.2： 编写一个程序，默认情况下打印标准输入的SHA256编码，并支持通过命令行flag定制，输出SHA384或SHA512哈希算法。
var f = flag.Int("f", 256, "请选择一个SHA加密方式(256、384、512)")

func printSHA() {
	var s string
	fmt.Print("请输入要加密的字符串:")
	fmt.Scanf("%v", &s)
	fmt.Printf("输入：%v\n", s)
	switch *f {
	case 384:
		re := sha512.Sum384([]byte(s))
		fmt.Printf("%x\n", re)
	case 512:
		re := sha512.Sum512([]byte(s))
		fmt.Printf("%x\n", re)
	default:
		re := sha256.Sum256([]byte(s))
		fmt.Printf("%x\n", re)
	}
}
