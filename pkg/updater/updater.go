package updater

import (
	"sync"

	"github.com/Ankush-Goyal/go-healthcheck/pkg/checks"
)

// Updater implements a health check that is explicitly set.
type Updater interface {
	checks.Checker

	// Update updates the current status of the health check.
	Update(status error)
}

// updater implements Checker and Updater, providing an asynchronous Update
// method.
// This allows us to have a Checker that returns the Check() call immediately
// not blocking on a potentially expensive check.
type updater struct {
	mu     sync.Mutex
	status error
}

// Check implements the Checker interface
func (u *updater) Check() error {
	u.mu.Lock()
	defer u.mu.Unlock()
	return u.status
}

// Update implements the Updater interface, allowing asynchronous access to
// the status of a Checker.
func (u *updater) Update(status error) {
	u.mu.Lock()
	defer u.mu.Unlock()
	u.status = status
}

// NewStatusUpdater returns a new updater
func NewStatusUpdater() Updater {
	return &updater{}
}
