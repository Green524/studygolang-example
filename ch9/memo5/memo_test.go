package memo

import (
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
