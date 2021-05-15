package checks

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
