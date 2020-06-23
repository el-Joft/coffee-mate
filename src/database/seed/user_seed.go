package seed

import (
	"coffee-mate/src/database/entity"

	"github.com/bxcodec/faker"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CreateUser to seed user data
func CreateUser(db *gorm.DB) error {
	user := entity.User{}
	if err := faker.FakeData(&user); err != nil {
		logrus.Errorln("Error user seed", err)
	}

	return db.Create(&entity.User{
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Password:  user.Password,
		Username:  user.Username,
		Age:       user.Age,
		Email:     user.Email,
	}).Error
}
