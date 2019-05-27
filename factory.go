package oncewait

import "sync"

// Factory offers OnceWaiter which is identified according to key string.
type Factory struct {
	sets sync.Map
}

// Get returns instance of OnceWaiter associated with the key.
// In the first Get call, new instance is set.
func (f *Factory) Get(key string) *OnceWaiter {
	v, _ := f.sets.LoadOrStore(key, New())
	return v.(*OnceWaiter)
}

// Refresh replaces instance of OnceWaiter associated with the key.
// Usually Refresh should be called inside of Do's function f.
func (f *Factory) Refresh(key string) {
	f.sets.Store(key, New())
}
