package database

import (
	"log"

	"github.com/patrike-miranda/gin-go-rest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConnectDB() {
	connectionString := "host=localhost user=root password=root dbname=gingo port=5432 sslmode=disable"

	DB, err = gorm.Open(postgres.Open(connectionString))

	if err != nil {
		log.Panic("Error connecting database " + err.Error())
	}

	DB.AutoMigrate(&models.Student{})
}
