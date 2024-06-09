package database

import (
	"gin_api/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func DbConnect() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading env variables")
	}
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbname := os.Getenv("DBNAME")
	dbport := os.Getenv("DB_PORT")
	sslmode := os.Getenv("SSLMODE")

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + dbport + " sslmode=" + sslmode
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic("Error connection with database !")
	}
	DB.AutoMigrate(&models.Student{})

}
