package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/auth"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/crypto/bcrypt"
)

// Mutation

func (r *mutationResolver) Login(ctx context.Context, username string, password string) (*auth.AuthPayload, error) {
	userLogin, err := r.UserStore.GetUser("username = ?", username)

	if err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Check if password is correct
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		return nil, gqlerror.Errorf("Incorrect username or password!")
	}

	// Create jwt token
	token, err := auth.NewToken(userLogin)

	if err != nil {
		return nil, gqlerror.Errorf("Internal server error!")
	}

	return &auth.AuthPayload{Token: token}, nil
}
