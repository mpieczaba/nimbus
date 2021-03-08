package utils

import (
	"github.com/mpieczaba/nimbus/tag/tag_share"
)

func TagShareInputsToTagShares(tagShareInputs []tag_share.TagShareInput) []tag_share.TagShare {
	var tagShares []tag_share.TagShare

	for _, tagShareInput := range tagShareInputs {
		tagShares = append(tagShares, tag_share.TagShare{
			UserID:      tagShareInput.UserID,
			Permissions: tagShareInput.Permissions,
		})
	}

	return tagShares
}
