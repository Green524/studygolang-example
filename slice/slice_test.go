package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAppend(t *testing.T) {
	//append函数
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i, i+1)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}

	x = append(x, 1)
	x = append(x, 2, 3)
	x = append(x, 4, 5, 6)
	x = append(x, x...) // append the slice x
	fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
	fmt.Println("========================")
	var stack []int
	stack = append(stack, 1)
	top := stack[len(stack)-1]
	fmt.Println(top)
	stack = stack[:len(stack)-1] //去除元素之后收缩stack
	fmt.Println(stack)
	fmt.Println("========================")
	a := []int{1, 2, 3, 4, 5}
	a = remove(a, 1)
	fmt.Println(a)
}

func TestReverse1(t *testing.T) {
	a := [5]int{0, 1, 2, 3, 4}
	fmt.Println(reflect.TypeOf(a))
	fmt.Println(a)
	reverse1(&a)
	fmt.Println(a)
}

func TestRepear(t *testing.T) {
	string := []string{"abc", "abc", "abc", "aaa"}
	string = repeat(string)
	fmt.Println(string)
}

func TestSpace(t *testing.T) {
	b := []byte{1, 2, ' ', '1', ' ', ' ', 2, 3, ' ', ' '}
	fmt.Println(b)
	b = space(b)
	fmt.Println(b)
}
func TestRotate(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 2) //[2 3 4 5 0 1]
	fmt.Println(s)

	s = []int{0, 1, 2, 3, 4, 5}
	s = rotateRight(s, 4) //[2 3 4 5 0 1]
	fmt.Println(s)
}
