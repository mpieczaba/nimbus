package utils

import (
	"strings"
	"testing"
)

func TestWriteFile(t *testing.T) {
	err := WriteFile("test", strings.NewReader("test"))

	if err != nil {
		t.Errorf("File should be created!")
	}
}

func TestReadFile(t *testing.T) {
	_, err := ReadFile("test")

	if err != nil {
		t.Errorf("File should be read!")
	}
}

func TestRemoveFile(t *testing.T) {
	err := RemoveFile("test")

	if err != nil {
		t.Errorf("File should be removed!")
	}
}
