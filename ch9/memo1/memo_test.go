package memo

import (
	"fmt"
	"log"
	"studygolang-example/ch9/memotest"
	"testing"
	"time"
)

func TestMemo_Get(t *testing.T) {
	m := New(memotest.HTTPGetBody)
	for url := range memotest.IncomingURLs() {
		start := time.Now()
		value, err := m.Get(url)
		if err != nil {
			log.Print(err)
		}
		fmt.Printf("%s,%s,%d bytes \n", url, time.Since(start), len(value.([]byte)))
	}
}
