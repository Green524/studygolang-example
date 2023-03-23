package _interface

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
	"time"
)

func TestByteCounter_Write(t *testing.T) {
	var c ByteCounter
	c.Write([]byte("hello"))
	fmt.Println(c) // "5", = len("hello")
	c = 0          // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c) // "12", = len("hello, Dolly")

}

func Test1(t *testing.T) {
	var w io.Writer
	w = os.Stdout         // OK: *os.File has Write method
	w = new(bytes.Buffer) // OK: *bytes.Buffer has Write method
	//w = time.Second       // compile error: time.Duration lacks Write method
	fmt.Println(w)

	var rwc io.ReadWriteCloser
	rwc = os.Stdout // OK: *os.File has Read, Write, Close methods
	//rwc = new(bytes.Buffer) // compile error: *bytes.Buffer lacks Close method
	w = rwc // OK: io.ReadWriteCloser has Write method
	//rwc = w // compile error: io.Writer lacks Close method
	fmt.Println(time.Second)

	var s string = "f。"
	switch s {
	case "f", "f。":
		fmt.Println(s)
	}
	var intSet IntSet
	var str = intSet.String()
	fmt.Println(str)

	var _ fmt.Stringer = &intSet
	//var _ fmt.Stringer = intSet //编译错误，因为intSet类型没有实现fmt.Stringer接口

	var w1 io.Writer = new(bytes.Buffer)
	fmt.Println(w1)
}
