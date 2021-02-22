package resolvers

import (
	"context"
	"os"
	"time"

	"github.com/mpieczaba/nimbus/core/models"
	"github.com/mpieczaba/nimbus/user"

	"github.com/form3tech-oss/jwt-go"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Mutation

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*models.AuthPayload, error) {
	var userLogin user.User

	if err := r.DB.Set("gorm:auto_preload", true).Where("username = ?", username).First(&userLogin).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Create jwt token
	t := jwt.New(jwt.SigningMethodHS256)

	claims := t.Claims.(jwt.MapClaims)

	claims["id"] = userLogin.ID
	claims["username"] = userLogin.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	token, err := t.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return nil, gqlerror.Errorf("Internal server error!")
	}

	authPayload := models.AuthPayload{Token: token}

	return &authPayload, nil
}
