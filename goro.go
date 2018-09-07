package goro

import (
	"sync/atomic"
)

// Once is an object that will perform an action exactly one time, until reset,
// at which point, it will only perform an action once.
type Once uint32

// Do invokes f exactly one time, regardless of how many times Do is called.
func (o *Once) Do(f func()) {
	if !atomic.CompareAndSwapUint32((*uint32)(o), 0, 1) {
		return
	}
	f()
}

// Reset will allow the next invocation of Do to perform the specified action
// once again.
func (o *Once) Reset() {
	atomic.StoreUint32((*uint32)(o), 0)
}
