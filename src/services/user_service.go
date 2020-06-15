package services

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/middleware/exception"
	"coffee-mate/src/repositories"
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

// CreateUser -> create user service logic
func (s *UserService) CreateUser(user entity.User) repositories.GetUser {
	userExist := s.UserRepository.UserExist(
		repositories.UserExistParams{Email: user.Email},
	)

	if (userExist != entity.User{}) {
		exception.Conflict("User conflict", []map[string]interface{}{
			{"message": "User with this email already exist", "flag": "USER_ALREADY_EXIST"},
		})
	}
	data := s.UserRepository.CreateUser(user)
	return data
}
