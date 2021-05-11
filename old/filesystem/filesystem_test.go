package filesystem

import (
	"strings"
	"testing"
)

func TestWriteFile(t *testing.T) {
	fs := NewFilesystem()

	if err := fs.WriteFile("test", strings.NewReader("test")); err != nil {
		t.Errorf("File should be created!")
	}
}

func TestReadFile(t *testing.T) {
	fs := NewFilesystem()

	if _, err := fs.ReadFile("test"); err != nil {
		t.Errorf("File should be read!")
	}
}

func TestRemoveFile(t *testing.T) {
	fs := NewFilesystem()

	if err := fs.RemoveFile("test"); err != nil {
		t.Errorf("File should be removed!")
	}
}
