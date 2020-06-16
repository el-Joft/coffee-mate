package repositories

import (
	db "coffee-mate/src/database"
	"coffee-mate/src/database/entity"
	"time"

	"github.com/jinzhu/gorm"
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
	FullName   string     `json:"full_name,omitempty"`
	Email      string     `json:"email,omitempty"`
	Age        int64      `json:"age,omitempty"`
	Address    string     `json:"address,omitempty"`
	AvatarPath string     `json:"avatar_path,omitempty"`
	DeletedAt  *time.Time `json:"deleted_at,omitempty"`
}

// UserExistParams -> Optional params for user exist
type UserExistParams struct {
	Email string
	ID    uint
}

// UserExist -> method to check if user already exist in database by email or id
func (r *UserRepository) UserExist(param UserExistParams) entity.User {
	user := entity.User{}
	if param.ID == 0 {
		r.Conn.Select("email").Where(&entity.User{Email: param.Email}).First(&user)
	} else {
		r.Conn.Select("id").Where(&entity.User{ID: param.ID}).First(&user)
	}
	return user
}

// CreateUser -> method to add user in database
func (r *UserRepository) CreateUser(user entity.User) GetUser {
	r.Conn.Create(&user)
	userCreated := GetUser{}
	r.Conn.Select("full_name, email, address, age").Where("id = ?", user.ID).First(&userCreated)
	return userCreated
}
