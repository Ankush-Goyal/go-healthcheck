package registry

import (
	"encoding/json"
	"fmt"
	"net/http"
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

var reg *Registry

// NewRegistry creates a new registry. This isn't necessary for normal use of
// the package, but may be useful for unit tests so individual tests have their
// own set of checks.
func NewRegistry() *Registry {
	reg = &Registry{
		registeredChecks: make(map[string]checks.Checker),
	}
	return reg
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

// StatusHandler returns a JSON blob with all the currently registered Health Checks
// and their corresponding status.
// Returns 503 if any Error status exists, 200 otherwise
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	                                                                        if r.Method == "GET" {
		statuses := reg.CheckStatus()
		status := http.StatusOK

		// If there is an error, return 503
		if len(statuses) != 0 {
			status = http.StatusServiceUnavailable
		}

		statusResponse(w, status, statuses)
	} else {
		http.NotFound(w, r)
	}
}

// statusResponse completes the request with a response describing the health
// of the service.
func statusResponse(w http.ResponseWriter, status int, checks map[string]string) {
	p, _ := json.Marshal(checks)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", fmt.Sprint(len(p)))
	w.WriteHeader(status)
	w.Write(p)
}

func init() {
	reg = NewRegistry()
	http.HandleFunc("/debug/health", StatusHandler)
}
