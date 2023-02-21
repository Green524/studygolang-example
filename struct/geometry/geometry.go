package geometry

import (
	"math"
	"sync"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func X() float64 {
	return 0
}

type IntList struct {
	Value int
	Tail  *IntList
}

func (list *IntList) Sum() int {
	if list == nil {
		return 0
	}
	return list.Value + list.Tail.Sum()
}

type Values map[string][]string

func (v Values) Get(key string) []string {
	if vs := v[key]; len(vs) > 0 {
		return vs
	}
	return nil
}

// Add adds the value to key.
// It appends to any existing values associated with key.
func (v Values) Add(key, value string) {
	v[key] = append(v[key], value)
}

var (
	mu      sync.Mutex // guards mapping
	mapping = make(map[string]string)
)

//匿名struct,并初始化
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}
