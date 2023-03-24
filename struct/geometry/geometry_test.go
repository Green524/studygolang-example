package geometry

import (
	"fmt"
	"testing"
)

func TestDistance(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}
	fmt.Println(Distance(p, q)) // "5", function call
	fmt.Println(p.Distance(q))  // "5", method call
	//fmt.Println(p.X()) 报错

	c := &Point{1, 2}
	fmt.Println(c.X)

	i := IntList{}
	fmt.Println(i.Sum())
}
func TestValues_Get(t *testing.T) {
	m := Values{"lang": {"en"}} // direct construction
	m.Add("item", "1")
	m.Add("item", "2")

	fmt.Println(m.Get("lang")) // "en"
	fmt.Println(m.Get("q"))    // ""
	fmt.Println(m.Get("item")) // "1"      (first value)
	fmt.Println(m["item"])     // "[1 2]"  (direct map access)

	m = nil
	fmt.Println(m.Get("item")) // ""
	m.Add("item", "3")         // panic: assignment to entry in nil map

}

func TestCache(t *testing.T) {
	fmt.Println(cache)
	p := Point{1, 2}
	distance := p.Distance
	fmt.Println(distance(p))
}
