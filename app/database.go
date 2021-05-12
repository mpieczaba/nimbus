package app

import (
	"log"
	"os"

	"github.com/mpieczaba/nimbus/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func (app *App) ConnectToDatabase() {
	dsn := os.ExpandEnv("${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	app.db = db

	app.db.AutoMigrate(user.User{})

	log.Println("Successfully connected to the database!")
}
