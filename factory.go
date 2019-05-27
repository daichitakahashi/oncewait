package oncewait

import "sync"

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
