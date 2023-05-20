package memo

import (
	"studygolang-example/ch9/memotest"
	"testing"
)

func Test(t *testing.T) {
	m := New(memotest.HTTPGetBody)
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := New(memotest.HTTPGetBody)
	memotest.Concurrent(t, m)
}
