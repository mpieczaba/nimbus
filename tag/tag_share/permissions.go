package tag_share

import (
	"fmt"
	"io"
	"strconv"
)

type TagShareKind string

const (
	TagShareKindCoOwner TagShareKind = "CoOwner"
	TagShareKindViewer  TagShareKind = "Viewer"
)

var AllTagShareKind = []TagShareKind{
	TagShareKindCoOwner,
	TagShareKindViewer,
}

func (e TagShareKind) IsValid() bool {
	switch e {
	case TagShareKindCoOwner, TagShareKindViewer:
		return true
	}
	return false
}

func (e TagShareKind) String() string {
	return string(e)
}

func (e *TagShareKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TagShareKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TagShareKind", str)
	}
	return nil
}

func (e TagShareKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
