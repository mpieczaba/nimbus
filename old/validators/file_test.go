package validators

import (
	"testing"

	"github.com/mpieczaba/nimbus/file"
)

func TestFileName(t *testing.T) {
	val := New()

	if err := val.Validate(file.FileInput{
		Name: "/>|:&",
	}); err == nil {
		t.Errorf("Should not validate file name!")
	}
}
