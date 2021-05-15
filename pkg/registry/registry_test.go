package registry

import "testing"

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
