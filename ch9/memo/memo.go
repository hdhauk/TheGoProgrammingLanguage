// Package memo provides a concurrency-unsafe
// memoization of a function of type Func.
package memo

import "sync"

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready

}

// Memo caches the results of calling a Func.
type Memo struct {
	f     Func
	mu    sync.Mutex //guards cache
	cache map[string]*entry
}

// Func is the type of the function to memoize.
type Func func(key string) (interface{}, error)

type result struct {
	value interface{}
	err   error
}

// New creates a new Memo and returns the address of that Memo
func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]*entry)}
}

// Get returns either a precomputed output if it exists, otherwise it
// computes a new result, adds it to the memo and return the result
func (memo *Memo) Get(key string) (interface{}, error) {
	memo.mu.Lock()
	e := memo.cache[key]
	if e == nil {
		// First request for this key
		// This goroutine becomes responsible for computing
		// the value and broadcasting the ready condition.
		e = &entry{ready: make(chan struct{})}
		memo.cache[key] = e
		memo.mu.Unlock()

		e.res.value, e.res.err = memo.f(key) // Do actual computation
		close(e.ready)                       // broadcast ready condition
	} else {
		// This is a repeated request
		memo.mu.Unlock()
		<-e.ready //Wait for entry to be ready
	}
	return e.res.value, e.res.err
}
