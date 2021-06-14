package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Mutation

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*models.AuthPayload, error) {
	userLogin, err := r.Store.User.GetUser("username = ?", username)

	if err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Check if password is correct
	if err = bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Create jwt token
	token, err := auth.NewToken(userLogin)

	if err != nil {
		return nil, gqlerror.Errorf("Internal server error!")
	}

	return &models.AuthPayload{
		Token: token,
		User:  userLogin,
	}, nil
}
