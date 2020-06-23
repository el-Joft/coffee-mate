package migration

import (
	"coffee-mate/src/database/entity"

	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

// CreateUser is create user tabel for migration
func CreateUser(conn *gorm.DB) {
	conn.AutoMigrate(&entity.User{})

	logrus.Info("Success running migration")
}
