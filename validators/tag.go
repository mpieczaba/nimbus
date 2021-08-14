package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckTagName(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^[A-Za-z0-9_]+$", fl.Field().String())

	return ok
}
