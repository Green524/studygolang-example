package main

import (
	"fmt"
	"testing"
)

func TestSort(t *testing.T) {
	arr := []int{5, 4, 3, 9, 6, 8}
	Sort(arr)
}

type Point struct {
	X, Y int
}
type Circle struct {
	Point
	Radius int
}
type Wheel struct {
	Circle
	Spokes int
}

func TestStruct(t *testing.T) {
	var w Wheel
	w.X = 0
	w.Y = 1
	w.Radius = 10
	w.Spokes = 1
	//w = Wheel{0, 1, 10, 1}
	//w = Wheel{X: 0, Y: 1, Radius: 10, Spokes: 1}
	w1 := Wheel{Circle{Point{0, 1}, 10}, 1}
	w2 := Wheel{
		Circle: Circle{
			Point:  Point{0, 1},
			Radius: 10,
		},
		Spokes: 1,
	}
	fmt.Println(w1 == w2)
	fmt.Printf("%#v\n", w1)
	fmt.Printf("%#v\n", w2)
	w2.Radius = 11
	fmt.Println(w1 == w2)
}

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
