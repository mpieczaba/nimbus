package utils

import (
	"github.com/mpieczaba/nimbus/models"

	"github.com/rs/xid"
)

func FileTagsInputToTags(input models.FileTagsInput) []*models.Tag {
	var tags []*models.Tag

	for _, tagName := range input.TagNames {
		tags = append(tags, &models.Tag{
			ID:   xid.New().String(),
			Name: tagName,
			FileTags: []models.FileTag{
				{
					FileID:  input.FileID,
					TagName: tagName,
				},
			},
		})
	}

	return tags
}
