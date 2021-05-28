package store

import (
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/paginator"

	"github.com/vektah/gqlparser/v2/gqlerror"
	"gorm.io/gorm"
)

type FileStore struct {
	db *gorm.DB
}

func NewFileStore(db *gorm.DB) *FileStore {
	return &FileStore{
		db: db,
	}
}

func (s *FileStore) GetFile(query interface{}, args ...interface{}) (*models.File, error) {
	var file models.File

	if err := s.db.Where(query, args...).First(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *FileStore) GetAllFiles(after, before *string, first, last *int) (*models.FileConnection, error) {
	var fileConnection models.FileConnection
	var files []*models.File

	if err := s.db.Model(models.File{}).Scopes(paginator.Paginate(after, before, first, last)).Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all files!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(files) > 0 {
		fileConnection.Nodes = files

		for _, file := range files {
			cursor, err := paginator.EncodeCursor(file.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all files!")
			}

			fileConnection.Edges = append(fileConnection.Edges, &models.FileEdge{
				Cursor: cursor,
				Node:   file,
			})
		}

		if err := s.db.Model(models.File{}).Scopes(paginator.GetBefore(files[0].ID)).First(&models.File{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Model(models.File{}).Scopes(paginator.GetAfter(files[len(files)-1].ID)).First(&models.File{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileConnection.PageInfo = &pageInfo

	return &fileConnection, nil
}

func (s *FileStore) CreateFile(file *models.File) (*models.File, error) {
	if err := s.db.Create(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	return file, nil
}

func (s *FileStore) UpdateFile(file *models.File) (*models.File, error) {
	if err := s.db.Save(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	return file, nil
}

func (s *FileStore) DeleteFile(query interface{}, args ...interface{}) (*models.File, error) {
	var file models.File

	if err := s.db.Where(query, args...).First(&file).Delete(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}
