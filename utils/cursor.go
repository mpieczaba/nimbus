package utils

import (
	"encoding/base64"
	"encoding/json"
)

func EncodeCursor(fieldValue string) (string, error) {
	encodedCursor, err := json.Marshal(map[string]string{"id": fieldValue})

	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(encodedCursor), nil
}

func DecodeCursor(cursor string) (string, error) {
	decodedCursor, err := base64.StdEncoding.DecodeString(cursor)

	if err != nil {
		return "", err
	}

	var fieldSet map[string]string

	err = json.Unmarshal(decodedCursor, &fieldSet)

	if err != nil {
		return "", err
	}

	return fieldSet["id"], nil
}
