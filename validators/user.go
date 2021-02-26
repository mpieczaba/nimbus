package validators

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

func CheckUserUsername(fl validator.FieldLevel) bool {
	err, _ := regexp.MatchString("^(?:[A-Za-z0-9_]+)$", fl.Field().String())

	return err
}

func CheckUserPassword(fl validator.FieldLevel) bool {
	hasLower := strings.ContainsAny(fl.Field().String(), "abcdefghijklmnopqrstuvwxyz")
	hasCapital := strings.ContainsAny(fl.Field().String(), "ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	hasDigit := strings.ContainsAny(fl.Field().String(), "0123456789")
	hasSymbol := strings.ContainsAny(fl.Field().String(), " !\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~")

	if !hasLower || !hasCapital || !hasDigit || !hasSymbol {
		return false
	}

	return true
}
