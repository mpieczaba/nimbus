package store

import (
	"github.com/mpieczaba/nimbus/auth"
	"github.com/mpieczaba/nimbus/models"
	"github.com/mpieczaba/nimbus/store/scopes"
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

func (s *FileStore) GetFile(claims *auth.Claims, permission models.FilePermission, query interface{}, args ...interface{}) (*models.File, error) {
	var file models.File

	if err := s.db.Scopes(
		scopes.FilePermission(models.File{}, "file_id", permission, "collaborator_id = ? OR ? = ?", claims.ID, claims.Kind, models.UserKindAdmin),
	).Where(query, args...).First(&file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return &file, nil
}

func (s *FileStore) GetAllFiles(claims *auth.Claims, after, before *string, first, last *int, name *string, permission models.FilePermission, tags []string) (*models.FileConnection, error) {
	var fileConnection models.FileConnection
	var files []*models.File

	if err := s.db.Scopes(
		scopes.FilePermission(models.File{}, "file_id", permission, "collaborator_id = ? OR ? = ?", claims.ID, claims.Kind, models.UserKindAdmin),
		scopes.FileTag(models.File{}, "file_id", "id", "tag_name IN ?", tags),
		scopes.NameLike(models.File{}, "name", name),
		scopes.Paginate(after, before, first, last),
	).Find(&files).Error; err != nil {
		return nil, gqlerror.Errorf("Invalid pagination input or internal database error occurred while getting all files!")
	}

	pageInfo := models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
	}

	if len(files) > 0 {
		fileConnection.Nodes = files

		for _, file := range files {
			cursor, err := utils.EncodeCursor(file.ID)

			if err != nil {
				return nil, gqlerror.Errorf("An error occurred while getting all files!")
			}

			fileConnection.Edges = append(fileConnection.Edges, &models.FileEdge{
				Cursor: cursor,
				Node:   file,
			})
		}

		if err := s.db.Scopes(
			scopes.FilePermission(models.File{}, "file_id", permission, "collaborator_id = ? OR ? = ?", claims.ID, claims.Kind, models.UserKindAdmin),
			scopes.FileTag(models.File{}, "file_id", "id", "tag_name IN (?)", tags),
			scopes.NameLike(models.File{}, "name", name),
			scopes.GetBefore(files[0].ID),
		).First(&models.File{}).Error; err == nil {
			pageInfo.HasPreviousPage = true
		}

		if err := s.db.Scopes(
			scopes.FilePermission(models.File{}, "file_id", permission, "collaborator_id = ? OR ? = ?", claims.ID, claims.Kind, models.UserKindAdmin),
			scopes.FileTag(models.File{}, "file_id", "id", "tag_name IN (?)", tags),
			scopes.NameLike(models.File{}, "name", name),
			scopes.GetAfter(files[len(files)-1].ID),
		).First(&models.File{}).Error; err == nil {
			pageInfo.HasNextPage = true
		}

		pageInfo.StartCursor = &fileConnection.Edges[0].Cursor
		pageInfo.EndCursor = &fileConnection.Edges[len(fileConnection.Edges)-1].Cursor
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

func (s *FileStore) DeleteFile(file *models.File) (*models.File, error) {
	if err := s.db.Select("Collaborators", "FileTags").Delete(file).Error; err != nil {
		return nil, gqlerror.Errorf("File not found!")
	}

	return file, nil
}
