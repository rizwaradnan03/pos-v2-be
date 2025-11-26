package translate

import (
	"errors"
	"fmt"
	"pos-v2-be/internal/pkg/validate"

	"github.com/go-playground/validator/v10"
)

func TranslateError(err error) error {
	if errs, ok := err.(validator.ValidationErrors); ok {
		var message string

		for _, e := range errs {
			message += fmt.Sprintf("%s\n", e.Translate(validate.Translator))
		}

		return errors.New(message)
	}

	return err
}
