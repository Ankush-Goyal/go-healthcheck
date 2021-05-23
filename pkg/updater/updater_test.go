package updater

import (
	"errors"
	"testing"
)

func TestNewStatusUpdater(t *testing.T) {
	t.Run("New updater", func(t *testing.T) {
		u := NewStatusUpdater()
		t.Run("Update Error", func(t *testing.T) {
			u.Update(errors.New("Test error check"))
			if u.Check() == nil || u.Check().Error() != "Test error check" {
				t.Error("update error not working")
			}
		})
		t.Run("Update Success", func(t *testing.T) {
			u.Update(nil)
			if u.Check() != nil {
				t.Error("error should not have occured", u.Check())
			}
		})
	})

}
