package validations

import (
	"fmt"
	"reflect"

	"coffee-mate/src/middleware/exception"
	"strings"

	eng "github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

// use a single instance , it caches struct info
// var uni      *ut.UniversalTranslator

var (
	en       = eng.New()
	uni      = ut.New(en, en)
	validate = validator.New()
	trans, _ = uni.GetTranslator("en")
)

// Validate -> function to validate request
func Validate(schema interface{}, errors []map[string]interface{}) {
	/**
	 * create validator instance
	 */
	// NOTE: ommitting allot of error checking for brevity

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)

	// validate = validator.New()
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
	en_translations.RegisterDefaultTranslations(validate, trans)

	if err := validate.Struct(schema); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			errors = append(errors, map[string]interface{}{
				"message": fmt.Sprint(err), "flag": "INVALID_BODY"},
			)
		}

		for _, err := range err.(validator.ValidationErrors) {
			errors = append(errors, map[string]interface{}{
				"field":   fmt.Sprint(err.Field()),
				"message": fmt.Sprint(err.Translate(trans)), "flag": "INVALID_BODY"},
			)
		}
		exception.BadRequest("Validation error", errors)
	}
	if errors != nil {
		exception.BadRequest("Validation error", errors)
	}
}
