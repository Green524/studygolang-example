package memo

import (
	"fmt"
	"runtime"
	"studygolang-example/ch9/memotest"
	"testing"
	"time"
)

//func Test(t *testing.T) {
//	m := New(memotest.HTTPGetBody)
//	memotest.Sequential(t, m)
//}

func TestConcurrent(t *testing.T) {
	m := New(memotest.HTTPGetBodyDone)
	done := make(chan struct{})
	go memotest.ConcurrentDone(t, m, done)
	time.Sleep(1 * time.Second)
	done <- struct{}{}
	//m.Close()
}

func Test(t *testing.T) {
	//最多同时有2个操作系统线程调度goroutine
	runtime.GOMAXPROCS(2)
	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
