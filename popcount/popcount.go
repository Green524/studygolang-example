package popcount

import (
	"fmt"
	"sync"
)

// var pc [256]byte
// pc[i] is the population count of i.
//对于pc这类需要复杂处理的初始化，可以通过将初始化逻辑包装为一个匿名函数处理，像下面这样
var pc [256]byte /* = func() (pc [256]byte) {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
	return
}()*/
var mu sync.Once

/*
下面的代码定义了一个PopCount函数，用于返回一个数字中含二进制1bit的个数。它使用init初始化函数来生成辅助表格pc，
pc表格用于处理每个8bit宽度的数字含二进制的1bit的bit个数，这样的话在处理64bit宽度的数字时就没有必要循环64次，
只需要8次查表就可以了。（这并不是最快的统计1bit数目的算法，但是它可以方便演示init函数的用法，
并且演示了如何预生成辅助表格，这是编程中常用的技术）。
*/
func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
		fmt.Printf("i: %d i/2:%d, %d = %d + %d\n", i, i/2, pc[i], pc[i/2], byte(i&1))
	}
	//11 1011
	//10 1010
	// 5 0101
	fmt.Println(pc)
}
func lazyInit() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}
func Popcount(x uint64) int {
	//练习 9.2： 重写2.6.2节中的PopCount的例子，使用sync.Once，只在第一次需要用到的时候进行初始化。（虽然实际上，对PopCount这样很小且高度优化的函数进行同步可能代价没法接受。）
	mu.Do(lazyInit)
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

// 练习 2.3 重写PopCount函数，用一个循环代替单一的表达式。比较两个版本的性能。（11.4节将展示如何系统地比较两个不同实现的性能。）
func PopcountIter(x uint64) (n int) {
	for i := 0; i < 8; i++ {
		n += int(pc[x>>(i*8)])
	}
	return
}

//练习 2.4： 用移位算法重写PopCount函数，每次测试最右边的1bit，然后统计总数。比较和查表算法的性能差异。
func PopcountBit(x uint64) (n int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			n++
		}
		x = x >> 1
	}
	return
}

//练习 2.5：表达式x&(x-1)用于将x的最低的一个非零的bit位清零。使用这个算法重写PopCount函数，然后比较性能。
//1111 1111 255
//1111 1110 254
//1111 1101 253
//1111 1100 252
//1111 1011 251
func PopcountOverride(x uint64) (n int) {
	for x != 0 {
		x &= x - 1
		n++
		fmt.Println(n)
	}
	return
}

//todo 记录一下java的bug版本 输入255 go版本输出 8，java版本输出7(在于进入循环时先进行了计算,丢掉了1bit)
/**
public static int PopcountOverride(long x){
        int n = 0;
        for (;(x &= x - 1) != 0;){
            n++;
            System.out.println(n);
        }
        return n;
    }
*/
