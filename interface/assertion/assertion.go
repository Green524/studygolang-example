package assertion

import (
	"fmt"
	"io"
	"os"
	count "studygolang-example/interface"
)

func assert() {
	var w io.Writer = os.Stdout
	fmt.Println(w)
	//x.(T) 如果断言类型(T)为具体类型，那么会判断x是否与T相同（x就是T）
	f, ok := w.(*os.File) // success:  ok, f == os.Stdout
	fmt.Println(f, ok)
	b, ok := w.(*count.ByteCounter) // failure: !ok, b == nil
	fmt.Println(b, ok)

}

func assert1() {
	var w io.Writer
	w = os.Stdout
	//x.(T) 如果断言类型(T)为接口类型，那么会判断x是否满足T
	rw, ok := w.(io.ReadWriter) // success: *os.File has both Read and Write
	fmt.Printf("%v%T\n", ok, rw)
	w = new(count.ByteCounter)
	rw, ok = w.(io.ReadWriter) // panic: *ByteCounter has no Read method
	fmt.Printf("%v%T\n", ok, rw)

}
func assert2() {
	var w io.Writer
	rw := w.(io.ReadWriter) //interface is nil, not io.ReadWriter
	fmt.Printf("%v\n", rw)

}
func assert3() {
	var w io.Writer = os.Stdout
	rw := w.(io.ReadWriter) //成功，*os.File 满足 io.ReadWriter
	fmt.Printf("%v\n", rw)

}
