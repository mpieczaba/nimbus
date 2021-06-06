package filesystem

import (
	"io"
	"log"
	"os"

	"github.com/vektah/gqlparser/v2/gqlerror"
)

type Filesystem struct{}

func New() *Filesystem {
	return &Filesystem{}
}

func (fs *Filesystem) WriteFile(id string, file io.Reader) error {
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

func (fs *Filesystem) RemoveFile(id string) error {
	if err := os.Remove(os.Getenv("DATA_DIRECTORY_PATH") + "/" + id); err != nil {
		log.Println(err)

		return gqlerror.Errorf("Failed to remove file!")
	}

	return nil
}
