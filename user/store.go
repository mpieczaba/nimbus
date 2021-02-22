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

func (store *Store) GetUserById(id string) (*User, error) {
	var user User

	if err := store.db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}

func (store *Store) GetAllUsers() ([]*User, error) {
	var users []*User

	if err := store.db.Find(&users).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all users!")
	}

	return users, nil
}

func (store *Store) SaveUser(user *User) (*User, error) {
	if err := store.db.Save(user).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or user already exists!")
	}

	return user, nil
}

func (store *Store) DeleteUser(id string) (*User, error) {
	var user User

	if err := store.db.Where("id = ?", id).First(&user).Delete(&user).Error; err != nil {
		return nil, gqlerror.Errorf("User not found!")
	}

	return &user, nil
}
