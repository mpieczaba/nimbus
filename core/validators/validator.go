package validators

import (
	"fmt"

	"github.com/mpieczaba/nimbus/core/utils"

	"github.com/go-playground/validator/v10"
)

type Validator struct {
	validate *validator.Validate
}

func New() *Validator {
	val := &Validator{
		validate: validator.New(),
	}

	// User
	val.validate.RegisterValidation("username", CheckUserUsername)
	val.validate.RegisterValidation("password", CheckUserPassword)

	// File
	val.validate.RegisterValidation("filename", CheckFileName)

	// Tag
	val.validate.RegisterValidation("tagname", CheckTagName)

	return val
}

func (val *Validator) Validate(model interface{}) error {
	if err := val.validate.Struct(model); err != nil {
		return fmt.Errorf("Incorrect " + utils.ValidationError(err.(validator.ValidationErrors)) + "!")
	}

	return nil
}
