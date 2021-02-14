package validators

import (
	"testing"

	"github.com/mpieczaba/nimbus/core/models"
)

func TestFileName(t *testing.T) {
	val := New()

	if err := val.Validate(models.FileInput{
		Name: "/>|:&",
	}); err == nil {
		t.Errorf("Should not validate file name!")
	}
}
