package repositories

import (
	db "coffee-mate/src/database"
	"coffee-mate/src/database/entity"
	"time"

	"coffee-mate/src/middleware/exception"

	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
)

// UserRepository -> the propose of user repository is handling query for user entity
type UserRepository struct {
	Conn *gorm.DB
}

// URepository -> user repository instance to get user table connection
func URepository() UserRepository {
	return UserRepository{Conn: db.GetDB().Table("users")}
}

// GetUser -> get user struct format
type GetUser struct {
	ID        uuid.UUID  `json:"id,omitempty"`
	FirstName string     `json:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty"`
	Email     string     `json:"email,omitempty"`
	Age       int64      `json:"age,omitempty"`
	Username  string     `json:"username,omitempty"`
	Gender    string     `json:"gender,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// UserExistParams -> Optional params for user exist
type UserExistParams struct {
	Email    string
	ID       uuid.UUID
	Username string
}

// UserExist -> method to check if user already exist in database by email or id
func (r *UserRepository) UserExist(param UserExistParams) []entity.User {
	users := []entity.User{}
	if param.ID == uuid.Nil {
		r.Conn.Select("username, email").Where(&entity.User{Email: param.Email}).Or(&entity.User{Username: param.Username}).Find(&users)
	} else {

		r.Conn.Select("id").Where(&entity.Base{ID: param.ID}).First(&users)
	}
	return users
}

// CreateUser -> method to add user in database
func (r *UserRepository) CreateUser(user entity.User) GetUser {
	var errors []map[string]interface{}
	var err error
	err = r.Conn.Debug().Create(&user).Error
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.BadRequest("error", errors)
	}
	userCreated := GetUser{}
	r.Conn.Select("id, username, email, first_name, age").Where("id = ?", user.ID).First(&userCreated)
	return userCreated
}

// GetUserByEmail -> Get User by Email
func (r *UserRepository) GetUserByEmail(email string) GetUser {
	var errors []map[string]interface{}
	var err error
	returnUser := GetUser{}
	err = r.Conn.Select("id, username, email, first_name, age").Where("email = ?", email).First(returnUser).Error
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": err.Error()},
		)
		exception.NotFound("error", errors)
	}
	return returnUser
}

// GetUserByEmailForLogin -> Get User by Email
func (r *UserRepository) GetUserByEmailForLogin(email string) entity.User {
	var errors []map[string]interface{}
	var err error
	returnUser := entity.User{}
	err = r.Conn.Select("id, username, email, first_name, age, password").Where("email = ?", email).First(&returnUser).Error
	if err != nil {
		errors = append(errors, map[string]interface{}{
			"message": "Please Provide valid Login details"},
		)
		exception.NotFound("error", errors)
	}

	return returnUser
}
