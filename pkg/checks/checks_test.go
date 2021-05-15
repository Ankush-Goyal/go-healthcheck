package checks

import (
	"os"
	"testing"
)

func TestFileChecker(t *testing.T) {
	dir := t.TempDir()
	t.Cleanup(func() {
		os.RemoveAll(dir)
	})
	if err := FileChecker(dir).Check(); err == nil {
		t.Errorf(dir, " was expected as exists")
	}

	if err := FileChecker("NoSuchFile").Check(); err != nil {
		t.Errorf("NoSuchFile was expected as not exists, error:%v", err)
	}
}
