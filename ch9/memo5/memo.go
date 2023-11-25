package memo

import "fmt"

type Memo struct {
	requests chan request
}
type entry struct {
	res  result
	read chan struct{}
}

type Func func(key string, done chan<- struct{}) (interface{}, error)
type result struct {
	value interface{}
	err   error
}
type request struct {
	key      string
	response chan result
	done     chan struct{}
}

func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request)}
	go memo.server(f)
	return memo
}
func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
	for req := range memo.requests {
		go memo.done(cache, req.key, req.done)
		e := cache[req.key]
		if e == nil {
			e = &entry{read: make(chan struct{})}
			cache[req.key] = e
			go e.call(f, req.key, req.done)
		}
		go e.deliver(req.response)
	}
}
func (memo *Memo) done(cache map[string]*entry, key string, done <-chan struct{}) {
	fmt.Println(cache)
	<-done
	delete(cache, key)
	memo.Close()
	fmt.Println("清除缓存", cache)
}
func (e *entry) call(f Func, key string, done chan<- struct{}) {
	e.res.value, e.res.err = f(key, done)
	close(e.read)
}
func (e *entry) deliver(response chan<- result) {
	<-e.read
	response <- e.res
}
func (memo *Memo) Get(key string) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key: key, response: response}
	res := <-response
	return res.value, res.err
}
func (memo *Memo) GetDoneable(key string, done chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key: key, response: response, done: done}
	res := <-response
	return res.value, res.err
}
func (memo *Memo) Close() {
	close(memo.requests)
}
