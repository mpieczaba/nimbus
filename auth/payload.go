package auth

import (
	"github.com/mpieczaba/nimbus/user"
)

type AuthPayload struct {
	Token string     `json:"token"`
	User  *user.User `json:"user"`
}
