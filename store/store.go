package store

import "gorm.io/gorm"

type Store struct {
	User             *UserStore
	File             *FileStore
	FileTag          *FileTagStore
	FileCollaborator *FileCollaboratorStore
	Tag              *TagStore
}

func New(db *gorm.DB) *Store {
	return &Store{
		User:             NewUserStore(db),
		File:             NewFileStore(db),
		FileTag:          NewFileTagStore(db),
		FileCollaborator: NewFileCollaboratorStore(db),
		Tag:              NewTagStore(db),
	}
}
