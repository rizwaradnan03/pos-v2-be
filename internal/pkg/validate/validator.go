package validate

import (
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
)

var (
	Validator  *validator.Validate
	Translator ut.Translator
)

func InitValidator() {
	Validator = validator.New()

	ind := id.New()
	uni := ut.New(ind, ind)

	trans, _ := uni.GetTranslator("id")
	Translator = trans

	idTranslations.RegisterDefaultTranslations(Validator, Translator)
}
