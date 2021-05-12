package user

import (
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

func (s *UserStore) GetAllUsers() ([]*User, error) {
	var users []*User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all users!")
	}

	return users, nil
}

func (s *UserStore) CreateUser(user *User) (*User, error) {
	if err := s.db.Create(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}
