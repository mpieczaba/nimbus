package resolvers

//go:generate go run github.com/99designs/gqlgen

import (
	"github.com/mpieczaba/nimbus/api/generated"
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/cv"
	"github.com/mpieczaba/nimbus/file"
	"github.com/mpieczaba/nimbus/file/file_share"
	"github.com/mpieczaba/nimbus/file/file_tag"
	"github.com/mpieczaba/nimbus/filesystem"
	"github.com/mpieczaba/nimbus/tag"
	"github.com/mpieczaba/nimbus/tag/tag_share"
	"github.com/mpieczaba/nimbus/user"
	"github.com/mpieczaba/nimbus/validators"
)

type Store struct {
	User      *user.Store
	File      *file.Store
	FileShare *file_share.Store
	FileTag   *file_tag.Store
	Tag       *tag.Store
	TagShare  *tag_share.Store
}

type Resolver struct {
	Store      *Store
	Auth       *auth.Auth
	CV         *cv.CV
	Filesystem *filesystem.Filesystem
	Validator  *validators.Validator
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
