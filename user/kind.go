package user

import (
	"fmt"
	"io"
	"strconv"
)

type UserKind string

const (
	UserKindAdmin  UserKind = "Admin"
	UserKindUser   UserKind = "User"
	UserKindBanned UserKind = "Banned"
)

var AllUserKind = []UserKind{
	UserKindAdmin,
	UserKindUser,
	UserKindBanned,
}

func (e UserKind) IsValid() bool {
	switch e {
	case UserKindAdmin, UserKindUser, UserKindBanned:
		return true
	}
	return false
}

func (e UserKind) String() string {
	return string(e)
}

func (e *UserKind) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = UserKind(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid UserKind", str)
	}
	return nil
}

func (e UserKind) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
