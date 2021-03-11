package file_share

import (
	"fmt"
	"io"
	"strconv"
)

type FileShareKind string

const (
	FileShareKindCoOwner FileShareKind = "CoOwner"
	FileShareKindEditor  FileShareKind = "Editor"
	FileShareKindViewer  FileShareKind = "Viewer"
)

var AllFileShareKind = []FileShareKind{
	FileShareKindCoOwner,
	FileShareKindEditor,
	FileShareKindViewer,
}

func (e FileShareKind) IsValid() bool {
	switch e {
	case FileShareKindCoOwner, FileShareKindEditor, FileShareKindViewer:
		return true
	}
	return false
}

func (e FileShareKind) String() string {
	return string(e)
}

func (e *FileShareKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = FileShareKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid FileShareKind", str)
	}
	return nil
}

func (e FileShareKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
