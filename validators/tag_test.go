package validators

import (
	"testing"

	"github.com/mpieczaba/nimbus/tag"
)

func TestTagName(t *testing.T) {
	val := New()

	if err := val.Validate(tag.TagInput{
		Name: "test tag name!",
	}); err == nil {
		t.Errorf("Should not validate tag name!")
	}
}
