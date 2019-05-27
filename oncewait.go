package oncewait

import "sync"

// OnceWaiter is an object that will perform exactly one action.
// But wait for finish of first Do.
type OnceWaiter struct {
	once      sync.Once
	completed chan struct{}
}

// New returns fresh OnceWaiter.
func New() *OnceWaiter {
	return &OnceWaiter{completed: make(chan struct{})}
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
