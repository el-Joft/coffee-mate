package db

import (
	"coffee-mate/src/database/entity"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/jpoles1/gopherbadger/logging"
)

var conn *gorm.DB
var err error

// Connection -> create connection with credentials
func Connection() {
	authDB := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PASS"),
	)

	conn, err = gorm.Open("postgres", authDB)
	if err != nil {
		logging.Error("DB", err)
	}
}

// SeedConnection is method to create connection for seeder
func SeedConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// AppConnection -> method to create connection for application
func AppConnection() {
	if err := godotenv.Load(); err != nil {
		logging.Error("ENV", err)
	}
	Connection()
}

// TestConnection -> method to create connection for application testing
func TestConnection() {
	if _, err := os.Stat("../../.env.test"); !os.IsNotExist(err) {
		var err error
		err = godotenv.Load(os.ExpandEnv("../../.env.test"))
		if err != nil {
			log.Fatalf("Error getting env %v\n", err)
		}
		Connection()
	} else {
		CIBuild()
	}
}

// GetDB -> method to get connection instance
func GetDB() *gorm.DB {
	return conn
}

// DropAllTable -> method to drop all database table (using this only for testing)
func DropAllTable() {
	conn.DropTable(&entity.User{})
}

// RefreshUserTable ->
func RefreshUserTable() error {
	err := conn.DropTableIfExists(&entity.User{}).Error
	if err != nil {
		return err
	}
	err = conn.AutoMigrate(&entity.User{}).Error
	if err != nil {
		return err
	}
	log.Printf("Successfully refreshed table")
	return nil
}

// CIBuild -> continous integration
func CIBuild() {
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", "127.0.0.1", "5433", "developer", "coffee_mate_test", "developer")
	conn, err = gorm.Open("postgres", DBURL)
	if err != nil {
		logging.Error("DB", err)
	}

}
