package validations

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/middleware/exception"
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
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
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

// UserExistValidation -> validate if user exist
func UserExistValidation(userExist []entity.User, user entity.User) {
	var errors []map[string]interface{}
	switch len(userExist) {
	case 1:
		if userExist[0].Email == user.Email {
			errors = append(errors, map[string]interface{}{
				"message": "User with this email already exist",
				"field":   "email",
				"flag":    "USER_ALREADY_EXIST"},
			)
		}
		if userExist[0].Username == user.Username {
			errors = append(errors, map[string]interface{}{
				"message": "User with this username already exist",
				"field":   "username",
				"flag":    "USER_ALREADY_EXIST"},
			)
		}
		break
	case 2:
		errors = append(errors, map[string]interface{}{
			"message": "Email and Username already exist",
			"flag":    "USER_ALREADY_EXIST"},
		)
		break
	}
	exception.Conflict("User Conflict", errors)
}
