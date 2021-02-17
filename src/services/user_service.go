package services

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/middleware/exception"
	"coffee-mate/src/repositories"
	tokenutil "coffee-mate/src/utils/security/token"
	"coffee-mate/src/utils/security"
	"coffee-mate/src/validations"

	"golang.org/x/crypto/bcrypt"
)

// UserService -> the propose of user service is handling business logic application
type UserService struct {
	UserRepository repositories.UserRepository
}

// UService -> user service instance
func UService() UserService {
	return UserService{
		UserRepository: repositories.URepository(),
	}
}

// LoginDTO -> get user struct format
type LoginDTO struct {
	user  interface{}
	token string
	time  int64
}

// CreateUser -> create user service logic
func (s *UserService) CreateUser(user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(
		repositories.UserExistParams{Email: user.Email, Username: user.Username},
	)

	// log.Printf("Data %s\n", len(userExist))
	// log.Printf("Data %s\n", userExist[0].Username)

	if len(userExist) != 0 {
		validations.UserExistValidation(userExist, user)
	}
	data := s.UserRepository.CreateUser(user)
	return data
}

// LoginUser -> create user login logic
func (s *UserService) LoginUser(email, password string) interface{} {
	var errors []map[string]interface{}
	var err error
	var token string

	// check user email and password if they exist
	user := s.UserRepository.GetUserByEmailForLogin(email)

	// check if the password match
	err = security.VerifyPassword(user.Password, password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		if err != nil {
			errors = append(errors, map[string]interface{}{
				"message": "Please provide valid Login details"},
			)
			exception.BadRequest("error", errors)
		}
	}
	//compare the user from the request, with the one we defined:
	if user.Email != email || user.Password != password {
		if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
			if err != nil {
				errors = append(errors, map[string]interface{}{
					"message": "Please provide valid login details"},
				)
				exception.BadRequest("error", errors)
			}
		}
	}

	// Generate Auth token
	token, err = tokenutil.CreateToken(user.ID, user.Email, user.FirstName)
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.BadRequest("error", errors)
	}
	var resp = map[string]interface{}{}

	resp["token"] = token //Store the token in the response
	// this is used to copy struct to another based on match fields
	// resource := &repositories.GetUser{}
	// deepcopier.Copy(user).To(resource)
	// Password should not be returned
	user.Password = ""
	resp["user"] = user
	return resp

}
