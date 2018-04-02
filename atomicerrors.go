package goxcart

import (
	"sync"
)

// AtomicErrors aggregates threaded errors.
type AtomicErrors struct {
	atomicErrorsLock sync.Mutex
	atomicErrors     []error
}

// GetErrors queries aggregated errors.
func (o *AtomicErrors) GetErrors() []error {
	defer o.atomicErrorsLock.Unlock()
	o.atomicErrorsLock.Lock()
	return o.atomicErrors
}

// AddError inserts an error.
func (o *AtomicErrors) AddError(err error) {
	defer o.atomicErrorsLock.Unlock()
	o.atomicErrorsLock.Lock()
	o.atomicErrors = append(o.atomicErrors, err)
}
