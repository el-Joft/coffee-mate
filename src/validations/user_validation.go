package validations

import (
	"coffee-mate/src/validations/schemas"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

// CreateUser -> validations to create user
func CreateUser(c *gin.Context) {
	var errors []map[string]interface{}
	var user schemas.CreateUser
	if err := c.ShouldBindBodyWith(&user, binding.JSON); err != nil {
		errors = append(errors, map[string]interface{}{
			"message": fmt.Sprint(err.Error()), "flag": "INVALID_BODY"},
		)
	}

	userValidate := &schemas.CreateUser{
		Fullname: user.Fullname,
		Password: user.Password,
		Address:  user.Address,
		Age:      user.Age,
		Email:    user.Email,
	}

	validate.RegisterTranslation("passwd", trans, func(ut ut.Translator) error {
		return ut.Add("passwd", "{0} is not strong enough", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("passwd", fe.Field())
		return t
	})
	validate.RegisterValidation("passwd", func(fl validator.FieldLevel) bool {
		return len(fl.Field().String()) > 6
	})

	Validate(userValidate, errors)

}
