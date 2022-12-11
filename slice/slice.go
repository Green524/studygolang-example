package slice

import "fmt"

func main() {
	months := [...]string{1: "January", 12: "December"}
	fmt.Println(months)
	Q2 := months[4:7]
	summer := months[6:9]
	fmt.Println(Q2)     // ["April" "May" "June"]
	fmt.Println(summer) // ["June" "July" "August"]
	for _, s := range summer {
		for _, q := range Q2 {
			if s == q {
				fmt.Printf("%s appears in both\n", s)
			}
		}
	}
	//fmt.Println(summer[:20]) // panic: out of range

	endlessSummer := summer[:5] // extend a slice (within capacity)
	fmt.Println(endlessSummer)  // "[June July August September October]"

	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)
	reverse(a[:])
	fmt.Println(a) // "[5 4 3 2 1 0]"

	fmt.Println("=========================")
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"

	arr := make([]int, 1, 1)
	fmt.Println(len(arr), cap(arr))
}
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
