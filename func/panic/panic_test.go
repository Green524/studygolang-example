package panic

import (
	"fmt"
	"os"
	"runtime"
	"testing"
)

func TestPanic(t *testing.T) {
	switch s := "abc"; s {
	case "Spades": // ...
	case "Hearts": // ...
	case "Diamonds": // ...
	case "Clubs": // ...
	default:
		panic(fmt.Sprintf("invalid suit %q", s)) // Joker?
	}
}

func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

func TestF(t *testing.T) {
	f(3)
}

func TestPrintStack(t *testing.T) {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
