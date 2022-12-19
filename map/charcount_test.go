package main

import (
	"fmt"
	"testing"
)

func TestCharcount(t *testing.T) {
	addEdge("abc", "你好")

	fmt.Println(hasEdge("a", "你好"))
}
