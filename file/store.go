package file

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

func (s *Store) GetFile(query interface{}, args ...interface{}) (*File, error) {
	var file File

	if err := s.db.Where(query, args).First(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *Store) GetAllFiles() ([]*File, error) {
	var files []*File

	if err := s.db.Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}

func (s *Store) GetAllFilesWithCondition(query interface{}, args ...interface{}) ([]*File, error) {
	var files []*File

	if err := s.db.Where(query, args).Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting all files!")
	}

	return files, nil
}

func (s *Store) SaveFile(file *File) (*File, error) {
	if err := s.db.Save(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	return file, nil
}

func (s *Store) DeleteFile(query interface{}, args ...interface{}) (*File, error) {
	var file File

	if err := s.db.Where(query, args).First(&file).Delete(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *Store) GetAllFileShares(query interface{}, args ...interface{}) ([]*FileShare, error) {
	var fileShares []*FileShare

	if err := s.db.Where(query, args).Find(&fileShares).Error; err != nil {
		return nil, gqlerror.Errorf("Internal database error occurred while getting file shares!")
	}

	return fileShares, nil
}

func (s *Store) SaveFileTags(fileTags []*FileTag) ([]*FileTag, error) {
	if err := s.db.Save(&fileTags).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot save file tags!")
	}

	return fileTags, nil
}

func (s *Store) DeleteFileTags(query interface{}, args ...interface{}) ([]*FileTag, error) {
	var fileTags []*FileTag

	if err := s.db.Where(query, args).Find(&fileTags).Delete(&fileTags).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete file tags!")
	}

	return fileTags, nil
}

func (s *Store) SaveFileShares(fileShares []*FileShare) ([]*FileShare, error) {
	if err := s.db.Save(&fileShares).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot save file shares!")
	}

	return fileShares, nil
}

func (s *Store) DeleteFileShares(query interface{}, args ...interface{}) ([]*FileShare, error) {
	var fileShares []*FileShare

	if err := s.db.Where(query, args).Find(&fileShares).Delete(&fileShares).Error; err != nil {
		return nil, gqlerror.Errorf("Cannot delete file shares!")
	}

	return fileShares, nil
}
