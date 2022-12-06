package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//s := "left foot"
	//t := s
	//s += ", right foot"
	//
	//fmt.Println(s) // "left foot, right foot"
	//fmt.Println(t) // "left foot"
	//fmt.Println(string(s[0]))

	s := "Hello, 世界"
	fmt.Println(len(s)) // "13"
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\t%v\t%d\n", i, r, s[i:], size)
		i += size
	}
}
