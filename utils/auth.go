package utils

import (
	"os"
	"strings"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"github.com/vektah/gqlparser/v2/gqlerror"
)

func Auth(c *fiber.Ctx) (jwt.MapClaims, error) {
	auth := c.Get(fiber.HeaderAuthorization)

	if strings.HasPrefix(auth, "Bearer ") {
		t := strings.Split(auth, "Bearer ")[1]

		token, _ := jwt.Parse(t, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if token.Valid {
			return token.Claims.(jwt.MapClaims), nil
		}

	}

	return nil, gqlerror.Errorf("User must be signed in!")
}
