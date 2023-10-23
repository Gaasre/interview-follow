package db

import (
	"interview-follow/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Database *gorm.DB

func Init() *gorm.DB {
	dbUrl := "postgres://pg:pass@localhost:5430/interview"
	var err error
	Database, err = gorm.Open(postgres.Open(dbUrl), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to database open")

	Database.AutoMigrate((&models.User{}), (&models.Application{}), (&models.Interview{}))
	log.Println("Database migrated")

	return nil
}
