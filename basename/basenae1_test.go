package basename

import (
	"fmt"
	"testing"
)

func TestBasename1(t *testing.T) {
	fmt.Println(basename1("a/b/c.go")) // "c"
	fmt.Println(basename1("c.d.go"))   // "c.d"
	fmt.Println(basename1("abc"))
	fmt.Println(comma("111123456"))
	fmt.Println(commaBuf("-111123456.56"))
	fmt.Println(dis("bbcd", "bcdb"))

}
