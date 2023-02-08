package _func

import (
	"fmt"
	"math"
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
