package validators

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func CheckFileName(fl validator.FieldLevel) bool {
	err, _ := regexp.MatchString("^(?:[^/>|:&]+)$", fl.Field().String())

	return err
}
