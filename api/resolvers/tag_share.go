package resolvers

import (
	"context"

	"github.com/mpieczaba/nimbus/tag/tag_share"
	"github.com/mpieczaba/nimbus/user"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

// Mutation

func (r *mutationResolver) TagShareDelete(ctx context.Context, tagId string, userId string) (*tag_share.TagShare, error) {
	claims, err := r.Auth.GetClaims()

	if err != nil {
		return nil, err
	}

	// Check if user is banned
	if claims["kind"].(string) == "Banned" {
		return nil, gqlerror.Errorf("You have no permissions to delete tag share!")
	}

	tagShareToDelete, err := r.Store.TagShare.GetTagShare("tag_id = ? AND user_id = ?", tagId, userId)

	if err != nil {
		return nil, err
	}

	// Check permissions
	_, err = r.Store.Tag.GetTag("id = ? AND (owner_id = ? OR ? = 'Admin')", tagShareToDelete.TagID, claims["id"].(string), claims["kind"].(string))

	if err != nil {
		return nil, err
	}

	return r.Store.TagShare.DeleteTagShare("tag_id = ? AND user_id = ?", tagId, userId)
}

// Field resolver

func (r *tagShareResolver) User(ctx context.Context, obj *tag_share.TagShare) (*user.User, error) {
	return r.Store.User.GetUser("id = ?", obj.UserID)
}
