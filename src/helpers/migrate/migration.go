package main

import (
	db "coffee-mate/src/database"
	"coffee-mate/src/database/migration"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db.AppConnection()
	conn := db.GetDB()
	defer conn.Close()

	migration.CreateUser(conn)
}
