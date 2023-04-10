package thumbnail

import (
	"fmt"
	"testing"
)

var filenames = []string{"C:\\Users\\cher\\Pictures\\DSC_0014.JPG", "C:\\Users\\cher\\Pictures\\DSC_0016.JPG"}

func TestMakeThumbnails(t *testing.T) {
	makeThumbnails(filenames) // 0.45s
}

func TestMakeThumbnails2(t *testing.T) {
	makeThumbnails2(filenames)
}
func TestMakeThumbnails3(t *testing.T) {
	makeThumbnails3(filenames)
}
func TestMakeThumbnails4(t *testing.T) {
	fmt.Println(makeThumbnails4(filenames))
}
