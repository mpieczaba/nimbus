package utils

import (
	"regexp"
	"strings"

	"github.com/go-playground/validator/v10"
)

// Function that returns PascalCase error fields in normal lowercase format
func ValidationError(errs validator.ValidationErrors) string {
	reg := regexp.MustCompile("([A-Z]{2,})")
	err := reg.ReplaceAllStringFunc(errs[0].Field(), unCapital)

	reg = regexp.MustCompile("([A-Z])")
	err = reg.ReplaceAllString(err, " $1")

	return strings.TrimSpace(strings.ToLower(err))
}

func unCapital(str string) string {
	return str[0:1] + strings.ToLower(str[1:])
}
