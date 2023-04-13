package thumbnail

import (
	"fmt"
	"gopl.io/ch8/thumbnail"
	"log"
	"os"
	"sync"
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
	//如果遇到err,会导致另一个channel一直卡死或者OOM
	//解决办法是使用缓存channel,或者在卡死之后使用另一个goroutine去排空channel
	for range filenames {
		if err := <-errors; err != nil {
			return err
		}
	}
	return nil
}
func makeThumbnails5(filenames []string) (thumbnails []string, err error) {
	type item struct {
		thumbnail string
		err       error
	}
	ch := make(chan item, len(filenames))
	for _, f := range filenames {
		go func(f string) {
			var it item
			it.thumbnail, it.err = thumbnail.ImageFile(f)
			ch <- it
		}(f)
	}
	for range filenames {
		it := <-ch
		if it.err != nil {
			return nil, it.err
		}
		thumbnails = append(thumbnails, it.thumbnail)
	}
	return thumbnails, nil
}

func makeThumbnails6(filenames <-chan string) int64 {
	sizes := make(chan int64)
	var wg sync.WaitGroup
	for f := range filenames {
		wg.Add(1)
		go func(f string) {
			defer wg.Done()
			thum, err := thumbnail.ImageFile(f)
			if err != nil {
				log.Println(err)
				return
			}
			info, _ := os.Stat(thum)
			sizes <- info.Size()
		}(f)
	}

	go func() {
		wg.Wait()
		close(sizes)
	}()
	var total int64
	for size := range sizes {
		total += size
	}
	return total
}
