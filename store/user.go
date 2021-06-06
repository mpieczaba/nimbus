package store

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/utils"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (s *UserStore) GetUser(query interface{}, args ...interface{}) (*models.User, error) {
	var user models.User

	if err := s.db.Where(query, args...).First(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (s *UserStore) GetAllUsers(after, before *string, first, last *int, username *string) (*models.UserConnection, error) {
	var userConnection models.UserConnection
	var users []*models.User

	if err := s.db.Scopes(
		scopes.NameLike(models.User{}, "username", username),
		scopes.Paginate(after, before, first, last),
	).Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all users!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(users) > 0 {
		userConnection.Nodes = users

		for _, user := range users {
			cursor, err := utils.EncodeCursor(user.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all users!")
			}

			userConnection.Edges = append(userConnection.Edges, &models.UserEdge{
				Cursor: cursor,
				Node:   user,
			})
		}

		if err := s.db.Scopes(
			scopes.NameLike(models.User{}, "username", username),
			scopes.GetBefore(users[0].ID),
		).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(
			scopes.NameLike(models.User{}, "username", username),
			scopes.GetAfter(users[len(users)-1].ID),
		).First(&models.User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	userConnection.PageInfo = &pageInfo

	return &userConnection, nil
}

func (s *UserStore) CreateUser(user *models.User) (*models.User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (s *UserStore) UpdateUser(user *models.User) (*models.User, error) {
	if err := s.db.Save(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (s *UserStore) DeleteUser(query interface{}, args ...interface{}) (*models.User, error) {
	var user models.User

	if err := s.db.Where(query, args...).First(&user).Delete(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}
