package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 9, 6, 8}
	Sort(arr)
}

type Point struct{ X, Y int }
type address struct {
	hostname string
	port     int
}

func TestNew(t *testing.T) {
	ps := &tree{1, nil, nil}
	fmt.Println(ps)

	ps1 := new(tree)
	fmt.Println(ps1)
	*ps1 = tree{1, nil, nil}
	fmt.Println(ps1)

	p := Point{1, 2}
	q := Point{2, 1}
	fmt.Println(p.X == q.X && p.Y == q.Y) // "false"
	fmt.Println(p == q)

	hits := make(map[address]int)
	hits[address{"golang.org", 443}]++
	fmt.Println(hits)
}
