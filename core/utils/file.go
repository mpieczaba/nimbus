package utils

import (
	"io"
	"io/ioutil"
	"os"
)

func CreateDataDirectory() error {
	if _, err := os.Stat(os.Getenv("DATA_DIRECTORY_PATH")); os.IsNotExist(err) {
		if err := os.Mkdir(os.Getenv("DATA_DIRECTORY_PATH"), 0777); err != nil {
			return err
		}
	}

	return nil
}

func WriteFile(id string, file io.Reader) error {
	fileContent, err := ioutil.ReadAll(file)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(os.Getenv("DATA_DIRECTORY_PATH")+"/"+id, fileContent, 0777)
}

func ReadFile(id string) ([]byte, error) {
	return ioutil.ReadFile(os.Getenv("DATA_DIRECTORY_PATH") + "/" + id)
}

func RemoveFile(id string) error {
	return os.Remove(os.Getenv("DATA_DIRECTORY_PATH") + "/" + id)
}
