package entity

// User -> user entity schema
type User struct {
	Base
	ID         uint   `gorm:"primary_key" json:"id"`
	FullName   string `json:"fullname" faker:"fullname"`
	Age        int64  `json:"age"`
	Email      string `gorm:"type:varchar(100);unique_index" json:"email" faker:"email"`
	Address    string `gorm:"index:addr";null json:"address" faker:"word"`
	AvatarPath string `gorm:"size:255;null;" json:"avatar_path"`
	Password   string `gorm:"type:varchar(255)" json:"password" faker:"password"`
}
