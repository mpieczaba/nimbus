package store

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/store/scopes/filters"
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

	if err := s.db.Scopes(scopes.Paginate(
		s.db.Scopes(filters.FilterByUsername(username)).Model(models.User{}),
		after, before, first, last,
	)).Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all users!")
	}

	pageInfo := utils.GetEmptyPageInfo()

	if len(users) > 0 {
		for _, user := range users {
			userConnection.Edges = append(userConnection.Edges, &models.UserEdge{
				Cursor: utils.EncodeCursor(user.ID),
				Node:   user,
			})
		}

		pageInfo.StartCursor = &userConnection.Edges[0].Cursor
		pageInfo.EndCursor = &userConnection.Edges[len(userConnection.Edges)-1].Cursor

		if err := s.db.Scopes(scopes.HasPreviousPage(
			s.db.Scopes(filters.FilterByUsername(username)).Model(models.User{}),
			users[0].ID,
		)).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(scopes.HasNextPage(
			s.db.Scopes(filters.FilterByUsername(username)).Model(models.User{}),
			users[len(users)-1].ID,
		)).First(&models.User{}).Error; err == nil {
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
