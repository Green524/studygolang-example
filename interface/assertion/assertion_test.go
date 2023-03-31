package assertion

import (
	"fmt"
	"os"
	"testing"
)

func TestAssertion(t *testing.T) {
	assert()
}
func TestAssertion1(t *testing.T) {
	assert1()
}
func TestAssertion2(t *testing.T) {
	assert2()
}

func TestAssertion3(t *testing.T) {
	assert3()
}
func TestAssertion4(t *testing.T) {
	assert4()
}

func TestAssertion5(t *testing.T) {
	w := new(os.File)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", &w)
	(*w).Write([]byte("ff"))
	//writeString(, "Nihao")
}
