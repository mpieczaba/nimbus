package user

import (
	"github.com/mpieczaba/nimbus/api/models"
	"github.com/mpieczaba/nimbus/utils/paginator"

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

func (s *UserStore) GetUser(query interface{}, args ...interface{}) (*User, error) {
	var user User

	if err := s.db.Where(query, args...).First(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (s *UserStore) GetAllUsers(after, before *string, first, last *int) (*UserConnection, error) {
	var userConnection UserConnection
	var users []*User

	if err := s.db.Model(User{}).Scopes(paginator.Paginate(after, before, first, last)).Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all users!")
	}

	if len(users) > 0 {
		userConnection.Nodes = users

		for _, user := range users {
			cursor, err := paginator.EncodeCursor(user.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all users!")
			}

			userConnection.Edges = append(userConnection.Edges, &UserEdge{
				Cursor: cursor,
				Node:   user,
			})
		}

		pageInfo := models.PageInfo{
			HasNextPage:     false,
			HasPreviousPage: false,
		}

		if err := s.db.Model(User{}).Scopes(paginator.GetBefore(users[0].ID)).First(&User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Model(User{}).Scopes(paginator.GetAfter(users[len(users)-1].ID)).First(&User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}

		userConnection.PageInfo = &pageInfo
	}

	return &userConnection, nil
}

func (s *UserStore) CreateUser(user *User) (*User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (s *UserStore) UpdateUser(user *User) (*User, error) {
	if err := s.db.Save(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (s *UserStore) DeleteUser(query interface{}, args ...interface{}) (*User, error) {
	var user User

	if err := s.db.Where(query, args...).First(&user).Delete(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}
