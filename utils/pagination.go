package utils

import (
	"encoding/base64"

	"github.com/mpieczaba/nimbus/models"
)

func EncodeCursor(cursor string) string {
	return base64.StdEncoding.EncodeToString([]byte(cursor))
}

func DecodeCursor(cursor string) (string, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)

	if err != nil {
		return "", err
	}

	return string(decodedCursor), nil
}

func GetEmptyPageInfo() models.PageInfo {
	return models.PageInfo{
		HasNextPage:     false,
		HasPreviousPage: false,
		StartCursor:     nil,
		EndCursor:       nil,
	}
}
