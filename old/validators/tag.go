package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckTagName(fl validator.FieldLevel) bool {
	err, _ := regexp.MatchString("^(?:[A-Za-z0-9\\-]+)$", fl.Field().String())

	return err
}
