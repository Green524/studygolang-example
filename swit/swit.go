package swit

import (
	"fmt"
	"math/rand"
	"time"
)

var arr [3]string = [3]string{"heads", "tails", "default"}

func main() {
	fmt.Println(Signum(-1))
	fmt.Println(Signum(0))
	fmt.Println(Signum(1))
}

type Point struct {
	X, Y int
}

var p Point

func Signum(x int) int {
	switch true {
	case x > 0:
		fallthrough
	default:
		return 0
	case x < 0:
		return -1
	}
}

func coinflip() {
	var heads, tails int
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(3)
	switch arr[num] {
	case "heads":
		heads++
		fmt.Println("heads")
	case "tails":
		tails++
		fmt.Println("tails")
	default:
		fmt.Println("landed on edge!")
	}
}
