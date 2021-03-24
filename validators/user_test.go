package validators

import (
	"testing"

	"github.com/mpieczaba/nimbus/user"
)

func TestUsername(t *testing.T) {
	val := New()

	if err := val.Validate(user.UserInput{
		Username: "ęąśćż",
	}); err == nil {
		t.Errorf("Should not validate username!")
	}
}

func TestPassword(t *testing.T) {
	val := New()

	if err := val.Validate(user.UserInput{
		Password: "test",
	}); err == nil {
		t.Errorf("Should not validate password!")
	}
}
