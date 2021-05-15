package checks

import (
	"errors"
	"os"
)

// Checker is the interface for a Health Checker
type Checker interface {
	// Check returns nil if the service is okay.
	Check() error
}

// CheckFunc is a convenience type to create functions that implement
// the Checker interface
type CheckFunc func() error

// Check Implements the Checker interface to allow for any func() error method
// to be passed as a Checker
func (cf CheckFunc) Check() error {
	return cf()
}

// FileChecker checks the existence of a file and returns an error
// if the file exists.
func FileChecker(f string) Checker {
	return CheckFunc(func() error {
		if _, err := os.Stat(f); err == nil {
			return errors.New("file exists")
		}
		return nil
	})
}
