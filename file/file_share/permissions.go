package file_share

import (
	"fmt"
	"io"
	"strconv"
)

type FileSharePermissions string

const (
	FileSharePermissionsCoOwner FileSharePermissions = "CoOwner"
	FileSharePermissionsEditor  FileSharePermissions = "Editor"
	FileSharePermissionsViewer  FileSharePermissions = "Viewer"
)

var AllFileSharePermissions = []FileSharePermissions{
	FileSharePermissionsCoOwner,
	FileSharePermissionsEditor,
	FileSharePermissionsViewer,
}

func (e FileSharePermissions) IsValid() bool {
	switch e {
	case FileSharePermissionsCoOwner, FileSharePermissionsEditor, FileSharePermissionsViewer:
		return true
	}
	return false
}

func (e FileSharePermissions) String() string {
	return string(e)
}

func (e *FileSharePermissions) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileSharePermissions(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileSharePermissions", str)
	}
	return nil
}

func (e FileSharePermissions) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
