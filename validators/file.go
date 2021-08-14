package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckFileName(fl validator.FieldLevel) bool {
	ok, _ := regexp.MatchString("^[^/>|:&]+$", fl.Field().String())

	return ok
}
