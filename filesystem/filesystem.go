package filesystem

import (
	"io"
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

func WriteFile(id string, file io.Reader) error {
	f, err := os.Create(os.Getenv("DATA_DIRECTORY_PATH") + "/" + id)

	if err != nil {
		log.Println(err)

		return gqlerror.Errorf("Failed to create file!")
	}

	defer f.Close()

	if _, err = io.Copy(f, file); err != nil {
		log.Println(err)

		return gqlerror.Errorf("Failed to clone file!")
	}

	return nil
}

func RemoveFile(id string) error {
	if err := os.Remove(os.Getenv("DATA_DIRECTORY_PATH") + "/" + id); err != nil {
		log.Println(err)

		return gqlerror.Errorf("Failed to remove file!")
	}

	return nil
}
