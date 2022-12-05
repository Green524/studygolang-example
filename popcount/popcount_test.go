package popcount

import (
	"fmt"
	"testing"
)

func TestPopcount(t *testing.T) {
	fmt.Println(Popcount(255))
}
func TestPopcountIter(t *testing.T) {
	fmt.Println(PopcountIter(255))
}

func TestPopcountBit(t *testing.T) {
	fmt.Println(PopcountBit(255))
}
func TestPopcountOverride1(t *testing.T) {
	fmt.Println(PopcountOverride(255))
}
