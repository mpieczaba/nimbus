package auth

import (
	"os"
	"time"

	"github.com/mpieczaba/nimbus/user"

	"github.com/form3tech-oss/jwt-go"
)

func NewToken(user *user.User) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
