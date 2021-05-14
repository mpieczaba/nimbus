package auth

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mpieczaba/nimbus/user"

	"github.com/form3tech-oss/jwt-go"
)

type Claims struct {
	jwt.StandardClaims
	ID       string        `json:"id"`
	Username string        `json:"usr"`
	Kind     user.UserKind `json:"kind"`
}

func NewToken(user *user.User) (string, error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		ID:       user.ID,
		Username: user.Username,
		Kind:     user.Kind,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
		},
	})

	return t.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func CheckToken(tokenString string) (*Claims, error) {
	if !strings.HasPrefix(tokenString, "Bearer ") {
		return nil, fmt.Errorf("Token is not a JWT token!")
	}

	t := strings.Split(tokenString, "Bearer ")[1]

	token, err := jwt.ParseWithClaims(t, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("Token is invalid!")
	}

	return token.Claims.(*Claims), nil
}
