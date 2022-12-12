package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	args1 := map[string]int{
		"name":   1,
		"age":    2,
		"weight": 3,
		"a":      4,
		"b":      5,
	}
	args1["c"] = 1
	fmt.Println(args1)
	//长度为0，容量为5
	sorts := make([]string, 0, len(args1))
	for key := range args1 {
		sorts = append(sorts, key)
	}
	fmt.Println(sorts)
	sort.Strings(sorts)
	fmt.Println(sorts)
	for _, key := range sorts {
		fmt.Printf("type:%T\tkey:%[1]v\tvalue:%d\n", key, args1[key])
	}

	var maps map[string]int
	fmt.Println(maps == nil)
	fmt.Println(len(maps) == 0)
	//maps["carol"] = 21 // panic: assignment to entry in nil map
	if name, ok := args1["name"]; ok {
		fmt.Println(name, ok)
	}

	map1 := make(map[string]int, 2)
	map2 := make(map[string]int, 2)
	map1["a"] = 1
	map1["b"] = 2
	map2["a"] = 1
	map2["b"] = 2
	fmt.Println(Equals(map1, map2))

	//dedup()

	strings := []string{"1", "2", "3", "4", "5"}
	Add(strings)
	Add(strings)
	fmt.Println(Count(strings))
}

var m = make(map[string]int, 2)

func k(list []string) string {
	return fmt.Sprintf("%q", list)
}

func Add(list []string) {
	m[k(list)]++
}
func Count(list []string) int {
	return m[k(list)]
}

func Equals(x, y map[string]int) bool {
	if len(x) != len(y) {
		return false
	}
	for k1, v1 := range x {
		if v2, ok := y[k1]; !ok || v2 != v1 {
			return false
		}
	}
	return true
}

func dedup() {
	seen := make(map[string]bool)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		if seen[line] == false {
			seen[line] = true
			fmt.Println(line)
		}
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "dedup: %v\n", err)
		os.Exit(1)
	}
}
