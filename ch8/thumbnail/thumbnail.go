package thumbnail

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
	"log"
)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}

}

func makeThumbnails2(filenames []string) {
	for _, f := range filenames {
		go thumbnail.ImageFile(f)
	}
}

func makeThumbnails3(filenames []string) {
	chs := make(chan int)
	for _, f := range filenames {
		go func(f string) {
			log.Println(f)
			thumbnail.ImageFile(f)
			log.Println("结束")
			chs <- 0
		}(f)
	}
	for range filenames {
		fmt.Println(<-chs)
	}
}

func makeThumbnails4(filenames []string) error {
	errors := make(chan error)
	for _, f := range filenames {
		go func(f string) {
			_, err := thumbnail.ImageFile(f)
			errors <- err
		}(f)
	}
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}
