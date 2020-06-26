package services

import (
	"coffee-mate/src/database/entity"
	"coffee-mate/src/repositories"
	"coffee-mate/src/validations"
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
 