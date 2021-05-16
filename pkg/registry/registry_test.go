package registry

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	t.Run("Registry", func(t *testing.T) {
		registry := NewRegistry()
		if registry == nil {
			t.Errorf("Empty registry returned")
		}
		t.Run("RegisterFunc", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				key := "test_func_key"
				registry.RegisterFunc(key, func() error {
					return nil
				})
				_, ok := registry.registeredChecks[key]
				if !ok {
					t.Error("Addition of new checker function failed")
				}
			})
		})
		t.Run("Register", func(t *testing.T) {
			t.Run("Success", func(t *testing.T) {
				key := "test_key"
				registry.RegisterFunc(key, func() error {
					return nil
				})
				_, ok := registry.registeredChecks[key]
				if !ok {
					t.Error("Addition of new checker function failed")
				}
			})
		})
		t.Run("CheckStatus", func(t *testing.T) {
			status := registry.CheckStatus()
			if len(status) != 0 {
				t.Error("Status returned error")
			}
		})
	})
}

func TestStatusHandler(t *testing.T) {
	t.Run("GET", func(t *testing.T) {
		req, err := http.NewRequest("GET", "/debug/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(StatusHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusOK {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusOK)
		}
	})

	t.Run("POST", func(t *testing.T) {
		req, err := http.NewRequest("POST", "/debug/health", nil)
		if err != nil {
			t.Fatal(err)
		}

		// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(StatusHandler)

		// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
		// directly and pass in our Request and ResponseRecorder.
		handler.ServeHTTP(rr, req)

		// Check the status code is what we expect.
		if status := rr.Code; status != http.StatusNotFound {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, http.StatusNotFound)
		}
	})

}
