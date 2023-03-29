package count

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
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

func TestInterfaceValue(t *testing.T) {
	var w io.Writer
	w = os.Stdout
	w.Write([]byte("hello"))
	w = new(bytes.Buffer)
	w.Write([]byte("hello"))
	w = nil //将type 和value 都置为nil
	//var x interface{} = time.Now()
	//fmt.Println(x)
	var x interface{} = []int{1, 2, 3}
	//fmt.Println(x == x)
	x = "nil"
	var x1 interface{} = "nil"
	fmt.Println(x == x)
	fmt.Println(x == x1)

}

const debug = false

func TestInterfaceValue1(t *testing.T) {
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}
	f(buf)
}

func f(out io.Writer) {
	//...do something...
	if out != nil {
		fmt.Println(out != nil)
		out.Write([]byte("done!\n"))
	}
}

func TestSortInterface(t *testing.T) {
	var names []string = []string{"c", "b", "f", "a"}
	//sort.Sort(sort.StringSlice(names))
	sort.Strings(names)
	fmt.Println(names)
}

func TestPrintTracks(t *testing.T) {
	//printTracks(tracks)
	fmt.Println("Martin Solveig" < "Maby")
	fmt.Println("Martin Solveig" < "Melilah")
	fmt.Println(strings.Compare("Martin Solveig", "Moby"))
	fmt.Println(strings.Compare("Martin Solveig", "Melilah"))
}

func TestByArtist(t *testing.T) {
	sort.Sort(byYear(tracks))
	printTracks(tracks)
	//反转排序
	sort.Sort(sort.Reverse(byYear(tracks)))
	printTracks(tracks)
	//自定义排序规则
	sort.Sort(customSort{tracks, func(x, y *Track) bool {
		if x.Title != y.Title {
			return x.Title < y.Title
		}
		if x.Year != y.Year {
			return x.Year < y.Year
		}
		if x.Length != y.Length {
			return x.Length < y.Length
		}
		return false
	}})
	printTracks(tracks)
	values := []int{3, 1, 2, 4}
	//检查是否是已排序
	fmt.Println(sort.IntsAreSorted(values))
	//int 排序方式
	sort.Ints(values)
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
	sort.Sort(sort.Reverse(sort.IntSlice(values)))
	fmt.Println(values)
	fmt.Println(sort.IntsAreSorted(values))
}
