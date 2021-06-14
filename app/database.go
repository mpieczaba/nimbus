package app

import (
	"log"
	"os"

	"github.com/mpieczaba/nimbus/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func (app *App) ConnectToDatabase() {
	dsn := os.ExpandEnv("${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}?charset=utf8mb4&parseTime=True&loc=Local")

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal(err)
	}

	app.db = db

	app.db.AutoMigrate(models.User{}, models.File{}, models.FileCollaborator{}, models.Tag{}, models.FileTag{})

	log.Println("Successfully connected to the database!")
}
