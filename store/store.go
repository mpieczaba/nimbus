package store

import "gorm.io/gorm"

type Store struct {
	User             *UserStore
	File             *FileStore
	FileCollaborator *FileCollaboratorStore
}

func New(db *gorm.DB) *Store {
	return &Store{
		User:             NewUserStore(db),
		File:             NewFileStore(db),
		FileCollaborator: NewFileCollaboratorStore(db),
	}
}
