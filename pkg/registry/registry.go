package registry

import (
	"sync"

	"github.com/Ankush-Goyal/go-healthcheck/pkg/checks"
)

// A Registry is a collection of checks. Most applications will use the global
// registry defined in DefaultRegistry. However, unit tests may need to create
// separate registries to isolate themselves from other tests.
type Registry struct {
	mu               sync.RWMutex
	registeredChecks map[string]checks.Checker
}

// NewRegistry creates a new registry. This isn't necessary for normal use of
// the package, but may be useful for unit tests so individual tests have their
// own set of checks.
func NewRegistry() *Registry {
	return &Registry{
		registeredChecks: make(map[string]checks.Checker),
	}
}

// Register associates the checker with the provided name.
func (registry *Registry) Register(name string, check checks.Checker) {
	registry.mu.Lock()
	defer registry.mu.Unlock()
	_, ok := registry.registeredChecks[name]
	if ok {
		panic("Check already exists: " + name)
	}
	registry.registeredChecks[name] = check
}

// RegisterFunc allows the convenience of registering a checker directly from
// an arbitrary func() error.
func (registry *Registry) RegisterFunc(name string, check func() error) {
	registry.Register(name, checks.CheckFunc(check))
}

// CheckStatus returns a map with all the current health check errors
func (registry *Registry) CheckStatus() map[string]string {
	registry.mu.RLock()
	defer registry.mu.RUnlock()
	statusKeys := make(map[string]string)
	for k, v := range registry.registeredChecks {
		err := v.Check()
		if err != nil {
			statusKeys[k] = err.Error()
		}
	}

	return statusKeys
}
