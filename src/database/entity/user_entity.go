package entity

import "time"

// User -> user entity schema
type User struct {
	ID         uint       `gorm:"primary_key" json:"id"`
	FullName   string     `json:"fullName" faker:"fullName"`
	Age        int64      `json:"age"`
	Email      string     `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Address    string     `gorm:"index:addr" json:"address" faker:"word"`
	AvatarPath string     `gorm:"size:255;null;" json:"avatar_path"`
	Password   string     `gorm:"type:varchar(255)" json:"password" faker:"password"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"update_at"`
	DeletedAt  *time.Time `sql:"index" json:"deleted_at"`
}
