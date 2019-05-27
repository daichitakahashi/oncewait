package oncewait

import "sync"

// OnceWaiter is an object that will perform exactly one action.
// But wait for finish of first Do.
type OnceWaiter struct {
	once      sync.Once
	completed chan struct{}
}

// Do calls the function f f and only if Do is being called for the first time
// for this instance of OnceWaiter, same as standard package's sync.Once.
// While function f is executing, second call waits for finish.
func (o *OnceWaiter) Do(f func()) {
	o.once.Do(func() {
		f()
		close(o.completed)
	})
	<-o.completed
}

// New returns fresh OnceWaiter.
func New() *OnceWaiter {
	return &OnceWaiter{completed: make(chan struct{})}
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
