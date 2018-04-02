package goxcart

import (
	"sync"
)

// AtomicContainerIDs provides thread-safe container ID records.
type AtomicContainerIDs struct {
	containerIDsLock sync.Mutex
	containerIDs     map[string]bool
}

// NewAtomicContainerIDs constructs an AtomicContainerIDs.
func NewAtomicContainerIDs() AtomicContainerIDs {
	var atomicContainerIDs AtomicContainerIDs
	atomicContainerIDs.containerIDs = make(map[string]bool)
	return atomicContainerIDs
}

// GetIDs queries the container IDs.
func (o *AtomicContainerIDs) GetIDs() []string {
	defer o.containerIDsLock.Unlock()

	var ids []string

	o.containerIDsLock.Lock()

	for id := range o.containerIDs {
		ids = append(ids, id)
	}

	return ids
}

// AddID inserts a container ID.
func (o *AtomicContainerIDs) AddID(id string) {
	defer o.containerIDsLock.Unlock()
	o.containerIDsLock.Lock()
	o.containerIDs[id] = true
}

// RemoveID deletes a container ID.
func (o *AtomicContainerIDs) RemoveID(id string) {
	defer o.containerIDsLock.Unlock()
	o.containerIDsLock.Lock()
	delete(o.containerIDs, id)
}
