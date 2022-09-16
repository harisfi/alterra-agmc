package configs

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"github.com/harisfi/alterra-agmc/day4/submission/models"
)

var DB *gorm.DB

func InitDB() {
	dbHost := os.Getenv("MYSQL_HOST")
	dbPort := os.Getenv("MYSQL_PORT")
	dbDatabase := os.Getenv("MYSQL_DBNAME")
	dbUsername := os.Getenv("MYSQL_USER")
	dbPassword := os.Getenv("MYSQL_PASSWORD")

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUsername, dbPassword, dbHost, dbPort, dbDatabase)

	var e error
	DB, e = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if e != nil {
		panic("failed to connect database")
	} else {
		log.Println("connected to database")
	}

	InitMigrate()
}

func InitMigrate() {
	DB.AutoMigrate(&models.User{})
}
