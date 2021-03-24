package filesystem

import (
	"io/ioutil"
	"os"
)

type Filesystem struct {
	dataDirectoryPath string
}

func NewFilesystem() *Filesystem {
	return &Filesystem{dataDirectoryPath: os.Getenv("DATA_DIRECTORY_PATH")}
}

func (fs *Filesystem) CreateDataDirectory() error {
	if _, err := os.Stat(fs.dataDirectoryPath); os.IsNotExist(err) {
		if err = os.Mkdir(fs.dataDirectoryPath, 0777); err != nil {
			return err
		}
	}

	return nil
}

func (fs *Filesystem) WriteFile(id string, file []byte) error {
	return ioutil.WriteFile(fs.dataDirectoryPath+"/"+id, file, 0777)
}

func (fs *Filesystem) ReadFile(id string) ([]byte, error) {
	return ioutil.ReadFile(fs.dataDirectoryPath + "/" + id)
}

func (fs *Filesystem) RemoveFile(id string) error {
	return os.Remove(fs.dataDirectoryPath + "/" + id)
}
