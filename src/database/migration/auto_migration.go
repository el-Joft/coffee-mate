package migration

import (
	"github.com/jinzhu/gorm"
	"coffee-mate/src/database/entity"
)

// AutoMigration -> to auth migrate the database
func AutoMigration(conn *gorm.DB) {
	conn.AutoMigrate(entity.User{})
}