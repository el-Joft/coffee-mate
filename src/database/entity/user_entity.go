package entity

import (
	"coffee-mate/src/utils/security"
	"time"
)

// User -> user entity schema
type User struct {
	Base
	FirstName string    `gorm:"first_name" json:"first_name,omitempty" faker:"firstName"`
	LastName  string    `gorm:"last_name" json:"last_name" faker:"lastName"`
	Username  string    `gorm:"size:255;not null;unique" json:"username" faker:"username"`
	Age       int64     `json:"age"`
	Gender    string    `gorm:"type:varchar(100);null" json:"gender" faker:"gender"`
	Email     string    `gorm:"type:varchar(100);unique" json:"email" faker:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password" faker:"password"`
	isActive  bool      `gorm:"type:bool json:isActive" faker:"boolean"`
	LastLogin time.Time `gorm:"column:last_login" json:"last_login"`
}

// BeforeSave ..
func (u *User) BeforeSave() error {
	hashedPassword, err := security.Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}
