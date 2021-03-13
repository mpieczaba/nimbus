package auth

import (
	"os"
	"strings"
	"time"

	"github.com/mpieczaba/nimbus/user"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Auth struct {
	ctx       *fiber.Ctx
	jwtSecret string
}

func NewAuth(ctx *fiber.Ctx) *Auth {
	return &Auth{
		ctx:       ctx,
		jwtSecret: os.Getenv("JWT_SECRET"),
	}
}

func (a *Auth) NewToken(user *user.User) (string, error) {
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = user.ID
	claims["username"] = user.Username
	claims["kind"] = user.Kind
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (a *Auth) GetClaims() (jwt.MapClaims, error) {
	auth := a.ctx.Get(fiber.HeaderAuthorization)

	if strings.HasPrefix(auth, "Bearer ") {
		t := strings.Split(auth, "Bearer ")[1]

		token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			return []byte(a.jwtSecret), nil
		})

		if token.Valid {
			return token.Claims.(jwt.MapClaims), nil
		}

	}

	return nil, gqlerror.Errorf("User must be signed in!")
}
