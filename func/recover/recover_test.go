package recover

import (
	"fmt"
	"golang.org/x/net/html"
	"testing"
)

func soleTitle(doc *html.Node) (title string, err error) {
	type bailout struct{}
	defer func() {
		switch p := recover(); p {
		case nil: // no panic
		case bailout{}: // "expected" panic
			err = fmt.Errorf("multiple title elements")
		default:
			panic(p) // unexpected panic; carry on panicking
		}
	}()
	// Bail out of recursion if we find more than one nonempty title.
	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "title" &&
			n.FirstChild != nil {
			if title != "" {
				panic(bailout{}) // multiple titleelements
			}
			title = n.FirstChild.Data
		}
	}, nil)
	if title == "" {
		return "", fmt.Errorf("no title element")
	}
	return title, nil
}

func forEachNode(doc *html.Node, f func(n *html.Node), t interface{}) {

}

func TestRecover(t *testing.T) {
	fmt.Println(Parse("nil"))
}

type Syntax int

func Parse(input string) (s *Syntax, err error) {
	defer func() {
		if p := recover(); p != nil {
			err = fmt.Errorf("internal error: %v", p)
		}
	}()
	// ...parser...
	panic(input)
}

//练习5.19： 使用panic和recover编写一个不包含return语句但能返回一个非零值的函数。
func TestRecover1(t *testing.T) {
	fmt.Println(Recover1(20))
}
func Recover1(x int) (result int) {
	defer func() {
		recover()
		result = x
	}()
	panic(x)
}
