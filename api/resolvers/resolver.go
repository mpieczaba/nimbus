package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/validators"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Resolver struct {
	Ctx       *fiber.Ctx
	DB        *gorm.DB
	UserStore *user.Store
	FileStore *file.Store
	Validator *validators.Validator
}

func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }

func (r *Resolver) User() generated.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }

func (r *Resolver) File() generated.FileResolver { return &fileResolver{r} }

type fileResolver struct{ *Resolver }

func (r *Resolver) Tag() generated.TagResolver { return &tagResolver{r} }

type tagResolver struct{ *Resolver }

func (r *Resolver) TagShare() generated.TagShareResolver { return &tagShareResolver{r} }

type tagShareResolver struct{ *Resolver }

func (r *Resolver) FileShare() generated.FileShareResolver { return &fileShareResolver{r} }

type fileShareResolver struct{ *Resolver }
