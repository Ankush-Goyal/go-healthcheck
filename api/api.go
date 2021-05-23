package api

import (
	"errors"
	"net/http"

	"github.com/Ankush-Goyal/go-healthcheck/pkg/registry"

	"github.com/Ankush-Goyal/go-healthcheck/pkg/updater"
)

var (
	u = updater.NewStatusUpdater()
)

// DownHandler registers a manual_http_status that always returns an Error
func DownHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		u.Update(errors.New("Manual Check"))
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// UpHandler registers a manual_http_status that always returns nil
func UpHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		u.Update(nil)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
}

// init sets up the two endpoints to bring the service up and down
func init() {
	registry.Register("manual_http_status", u)
	http.HandleFunc("/debug/health/down", DownHandler)
	http.HandleFunc("/debug/health/up", UpHandler)
}
