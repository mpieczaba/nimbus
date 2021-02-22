package user

import (
	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type Store struct {
	db *gorm.DB
}

func NewStore(db *gorm.DB) *Store {
	store := &Store{
		db: db,
	}

	return store
}

func (s *Store) GetUserById(id string) (*User, error) {
	var user User

	if err := s.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (s *Store) GetUserByUsername(username string) (*User, error) {
	var user User

	if err := s.db.Where("username = ?", username).First(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (s *Store) GetAllUsers() ([]*User, error) {
	var users []*User

	if err := s.db.Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all users!")
	}

	return users, nil
}

func (s *Store) SaveUser(user *User) (*User, error) {
	if err := s.db.Save(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (s *Store) DeleteUser(id string) (*User, error) {
	var user User

	if err := s.db.Where("id = ?", id).First(&user).Delete(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}