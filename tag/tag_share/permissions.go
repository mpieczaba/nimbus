package tag_share

import (
	"fmt"
	"io"
	"strconv"
)

type TagSharePermissions string

const (
	TagSharePermissionsCoOwner TagSharePermissions = "CoOwner"
	TagSharePermissionsViewer  TagSharePermissions = "Viewer"
)

var AllTagSharePermissions = []TagSharePermissions{
	TagSharePermissionsCoOwner,
	TagSharePermissionsViewer,
}

func (e TagSharePermissions) IsValid() bool {
	switch e {
	case TagSharePermissionsCoOwner, TagSharePermissionsViewer:
		return true
	}
	return false
}

func (e TagSharePermissions) String() string {
	return string(e)
}

func (e *TagSharePermissions) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TagSharePermissions(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TagSharePermissions", str)
	}
	return nil
}

func (e TagSharePermissions) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
