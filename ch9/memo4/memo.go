package memo

import (
	"sync"
)

type Memo struct {
	f     Func
	cache map[string]*entry
	mu    sync.Mutex
}
type entry struct {
	res  result
	read chan struct{}
}
type Func func(key string) (interface{}, error)
type result struct {
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {

		e = &entry{read: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key)
		close(e.read)
	} else {
		memo.mu.Unlock()
		<-e.read
	}
	return e.res.value, e.res.err
}
