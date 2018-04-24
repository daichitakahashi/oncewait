package oncewait

import "sync"

type onceSet struct {
	once    *sync.Once
	channel chan struct{}
}

// OnceWaiter is
type OnceWaiter struct {
	once      sync.Once
	completed chan struct{}
}

// Do is
func (o *OnceWaiter) Do(f func()) {
	o.once.Do(func() {
		f()
		close(o.completed)
	})
	<-o.completed
}

// Factory is
type Factory struct {
	sets sync.Map
}

// Get is
func (f *Factory) Get(key string) *OnceWaiter {
	v, _ := f.sets.LoadOrStore(key, &OnceWaiter{completed: make(chan struct{})})
	return v.(*OnceWaiter)
}

// Refresh is
func (f *Factory) Refresh(key string) {
	f.sets.Store(key, &OnceWaiter{completed: make(chan struct{})})
}
