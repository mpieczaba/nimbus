package validators

import (
	"testing"

	"github.com/mpieczaba/nimbus/core/models"
)

func TestTagName(t *testing.T) {
	val := New()

	if err := val.Validate(models.TagInput{
		Name: "test tag name!",
	}); err == nil {
		t.Errorf("Should not validate tag name!")
	}
}
