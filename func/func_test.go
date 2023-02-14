package _func

import (
	"fmt"
	"math"
	"strings"
	"testing"
	"time"
)

func TestFunc(t *testing.T) {
	fmt.Println(hypot(3, 4))
	fmt.Println(1 * time.Minute)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }
func TestTest(t *testing.T) {
	//f := square
	//fmt.Println(f(3)) // "9"

	//f = negative
	//fmt.Println(f(3))     // "-3"
	//fmt.Printf("%T\n", f) // "func(int) int"

	//f = product // compile error: can't assign func(int, int) int to func(int) int

	//var f func(int) int
	//f(3) // 此处f的值为nil, 会引起panic错误

	fmt.Println(strings.Map(add1, "HAL-9000")) // "IBM.:111"
	fmt.Println(strings.Map(add1, "VMS"))      // "WNT"
	fmt.Println(strings.Map(add1, "Admix"))    // "Benjy"
	fmt.Println(strings.Map(func(r rune) rune { return r + 1 }, "Admix"))
	fmt.Printf("%*s<%s>\n", 1, "", "a")
}
func add1(r rune) rune { return r + 1 }
