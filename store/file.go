package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
	"github.com/mpieczaba/nimbus/store/scopes/filters"
	"github.com/mpieczaba/nimbus/utils"

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

func (s *FileStore) GetFile(claims *auth.Claims, permission models.FilePermissions, query interface{}, args ...interface{}) (*models.File, error) {
	var file models.File

	if err := s.db.Scopes(filters.FilterFilesByFilePermissions(claims, permission)).Where(query, args...).First(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *FileStore) GetAllFiles(claims *auth.Claims, after, before *string, first, last *int, name *string, permission models.FilePermissions, tags []string) (*models.FileConnection, error) {
	var fileConnection models.FileConnection
	var files []*models.File

	if err := s.db.Scopes(scopes.Paginate(
		s.db.Scopes(
			filters.FilterFilesByFilePermissions(claims, permission),
			filters.FilterByName(name),
			filters.FilterFilesByTags(tags),
		).Model(models.File{}),
		after, before, first, last,
	)).Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all files!")
	}

	pageInfo := utils.GetEmptyPageInfo()

	if len(files) > 0 {
		for _, file := range files {
			fileConnection.Edges = append(fileConnection.Edges, &models.FileEdge{
				Cursor: utils.EncodeCursor(file.ID),
				Node:   file,
			})
		}

		pageInfo.StartCursor = &fileConnection.Edges[0].Cursor
		pageInfo.EndCursor = &fileConnection.Edges[len(fileConnection.Edges)-1].Cursor

		if err := s.db.Scopes(scopes.HasPreviousPage(
			s.db.Scopes(
				filters.FilterFilesByFilePermissions(claims, permission),
				filters.FilterByName(name),
				filters.FilterFilesByTags(tags),
			).Model(models.File{}),
			files[0].ID,
		)).First(&models.User{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(scopes.HasNextPage(
			s.db.Scopes(
				filters.FilterFilesByFilePermissions(claims, permission),
				filters.FilterByName(name),
				filters.FilterFilesByTags(tags),
			).Model(models.File{}),
			files[len(files)-1].ID,
		)).First(&models.User{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}
	}

	fileConnection.PageInfo = &pageInfo

	return &fileConnection, nil
}

func (s *FileStore) CreateFile(file *models.File, callback func() error) (*models.File, error) {
	tx := s.db.Begin()

	if err := tx.Create(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	if err := callback(); err != nil {
		return nil, err
	}

	return file, nil
}

func (s *FileStore) UpdateFile(file *models.File, callback func() error) (*models.File, error) {
	tx := s.db.Begin()

	if err := tx.Save(file).Error; err != nil {
		return nil, gqlerror.Errorf("Incorrect form data or file already exists!")
	}

	if err := callback(); err != nil {
		tx.Rollback()

		return nil, err
	}

	return file, nil
}

func (s *FileStore) DeleteFile(file *models.File, callback func() error) (*models.File, error) {
	tx := s.db.Begin()

	if err := tx.Select("Collaborators", "FileTags").Delete(file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	if err := callback(); err != nil {
		return nil, err
	}

	return file, nil
}
